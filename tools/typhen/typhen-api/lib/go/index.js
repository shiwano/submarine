'use strict';

var path = require('path');
var assert = require('assert');
var fs = require('fs-extra');
var execSync = require('child_process').execSync;
var _ = require('lodash');

module.exports = function(typhen, options, helpers) {
  assert(options.importBasePath, 'options.importBasePath is empty');

  helpers = _.assign(helpers, {
    requiredModules: function(symbol, withErrorType) {
      assert(symbol.isType || symbol.isModule, 'should be a type or module');
      var types;
      var currentModule;

      if (symbol.isFunction) {
        types = symbol.callSignatures[0].parameters.map(function(p) { return p.type; });
        currentModule = null;
      } else if (symbol.isType) {
        types = symbol.properties.map(function(p) {
          return p.type.isArray ? p.type.type : p.type;
        });
        currentModule = symbol.parentModule;
      } else { // symbol.isModule
        types = helpers.realTimeMessages(symbol.types);
        currentModule = null;
      }

      if (withErrorType) {
        var errorType = helpers.errorType(symbol);
        if (errorType !== undefined) {
          types.push(errorType);
        }
      }

      return _.chain(types)
        .map(function(type) {
          if (type.parentModule === null ||
              type.parentModule === currentModule ||
              type.parentModule.isGlobalModule ||
              type.isPrimitiveType) {
            return null;
          }
          return { module: type.parentModule, alias: helpers.namespace(type, '_'), path: helpers.namespace(type, '/') };
        })
        .filter(function(x) { return x !== null; })
        .uniq(function(x) { return x.path; })
        .value();
    },
    webApiModules: function(module) {
      return _.chain(module.modules)
        .filter(function(module) { return helpers.isWebApiModule(module); })
        .map(function(module) { return { module: module, alias: helpers.moduleName(module, '_'), path: helpers.moduleName(module, '/') }; })
        .uniq(function(x) { return x.path; })
        .value();
    },
    webSocketApiModules: function(module) {
      return _.chain(module.modules)
        .filter(function(module) { return helpers.isWebSocketApiModule(module); })
        .map(function(module) { return { module: module, alias: helpers.moduleName(module, '_'), path: helpers.moduleName(module, '/') }; })
        .uniq(function(x) { return x.path; })
        .value();
    },
    namespace: function(type, sep) {
      return type.ancestorModules.map(function(m) { return typhen.helpers.lowerCamelCase(m.name); }).join(sep);
    },
    moduleName: function(module, sep) {
      return [helpers.namespace(module, sep), typhen.helpers.lowerCamelCase(module.name)].join(sep);
    },
    typeName: function(type, currentModule, isOptional, hasPointerMark) {
      var pointerMark = hasPointerMark ? '*' : '';
      if (type.isPrimitiveType && type.name === 'nil') {
        return pointerMark + 'typhenapi.Void';
      } else if (type.isPrimitiveType || type.isEnum) {
        return isOptional ? pointerMark + type.name : type.name;
      } else if (type.isArray) {
        return '[]' + helpers.typeName(type.type, currentModule, false, hasPointerMark)
      } else if (type.parentModule !== null && type.parentModule !== currentModule) {
        return pointerMark + [helpers.namespace(type, '_'), typhen.helpers.upperCamelCase(type.name)].join('.');
      } else {
        return pointerMark + typhen.helpers.upperCamelCase(type.name);
      }
    },
    isErrorType: function(type) {
      return type.isType && type.name === 'Error' && type.ancestorModules.length === 1;
    },
    isRequiredRequestBody: function(func) {
      var method = helpers.upperCaseMethod(func);
      return method === 'POST' || method === 'PATCH' || method === 'PUT';
    },
    upperCaseMethod: function(func) {
      return helpers.method(func).toUpperCase();
    },
    optionalParameters: function(parameters) {
      return parameters.filter(function(p) { return p.isOptional; });
    }
  });

  return {
    namespaceSeparator: '.',
    helpers: helpers,

    rename: function(symbol, name) {
      if (name === 'integer') {
        return 'int64';
      } else if (name === 'number') {
        return 'float64';
      } else if (name === 'boolean') {
        return 'bool';
      } else if (name === 'void') {
        return 'nil';
      }
      return name;
    },

    generate: function(g, types, modules, targetModule) {
      fs.removeSync(path.join(g.outputDirectory, 'typhenapi'));
      g.generateFiles('lib/go/templates/core', '**/*.go', 'typhenapi');

      types.forEach(function(type) {
        switch (type.kind) {
          case typhen.SymbolKind.Enum:
            g.generate('lib/go/templates/type/enum.hbs', 'underscore:typhenapi/type/**/*.go', type);
            break;
          case typhen.SymbolKind.Interface:
            if (!type.isGenericType || type.typeArguments.length > 0) {
              g.generate('lib/go/templates/type/struct.hbs', 'underscore:typhenapi/type/**/*.go', type);
            }
            break;
          case typhen.SymbolKind.ObjectType:
            g.generate('lib/go/templates/type/struct.hbs', 'underscore:typhenapi/type/**/*.go', type);
            break;
        }
      });

      modules.forEach(function(module) {
        if (helpers.isWebApiModule(module)) {
          g.generate('lib/go/templates/web/api.hbs', 'underscore:typhenapi/web/**/*/api.go', module);

          module.functions.forEach(function(func) {
            g.generate('lib/go/templates/web/request_body.hbs', 'underscore:typhenapi/web/**/*_request_body.go', func);
          });
        }

        if (helpers.isWebSocketApiModule(module)) {
          g.generate('lib/go/templates/websocket/api.hbs', 'underscore:typhenapi/websocket/**/*/api.go', module);
        }
      });

      g.files.forEach(function(file) {
        file.contents = execSync('gofmt', { input: file.contents });
      });
    }
  };
};
