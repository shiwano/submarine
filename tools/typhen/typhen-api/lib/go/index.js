'use strict';

var path = require('path');
var assert = require('assert');
var fs = require('fs-extra');
var execSync = require('child_process').execSync;
var _ = require('lodash');

module.exports = function(typhen, options, helpers) {
  assert(options.importBasePath, 'options.importBasePath is empty');

  helpers = _.assign(helpers, {
    requiredModules: function(typeOrModule) {
      assert(typeOrModule.isType || typeOrModule.isModule, 'should be a type or a module');
      var types;
      var currentModule;

      if (typeOrModule.isType) {
        types = typeOrModule.properties.map(function(p) { return p.type; });
        currentModule = typeOrModule.parentModule;
      } else {
        types = typeOrModule.variables.map(function(v) { return v.type; });
        currentModule = null;
      }
      return _.chain(types)
        .map(function(type) {
          if (type.parentModule === null ||
              type.parentModule === currentModule ||
              type.parentModule.isGlobalModule ||
              type.isPrimitiveType) {
            return null;
          } else {
            return {
              alias: helpers.namespace(type, '_'),
              path: helpers.namespace(type, '/'),
            };
          }
        })
        .filter(function(x) { return x !== null; })
        .uniq(function(x) { return x.path; })
        .value();
    },
    webSocketApiModules: function(module) {
      return _.chain(module.modules)
        .filter(function(module) { return helpers.isWebSocketApiModule(module); })
        .map(function(module) {
          return {
            alias: helpers.moduleName(module, '_'),
            path: helpers.moduleName(module, '/'),
          };
        })
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
        return '';
      } else if (type.isPrimitiveType || type.isArray) {
        return isOptional ? pointerMark + type.name : type.name;
      } else if (type.parentModule !== null && type.parentModule !== currentModule) {
        return pointerMark + [helpers.namespace(type, '_'), typhen.helpers.upperCamelCase(type.name)].join('.');
      } else {
        return pointerMark + typhen.helpers.upperCamelCase(type.name);
      }
    },
    isErrorType: function(type) {
      return type.isType && type.name === 'Error' && type.ancestorModules.length === 1;
    }
  });

  return {
    namespaceSeparator: '.',
    helpers: helpers,

    rename: function(symbol, name) {
      if (symbol.kind === typhen.SymbolKind.Array) {
        return '[]' + typhen.helpers.upperCamelCase(symbol.type);
      } else if (name === 'integer') {
        return 'int';
      } else if (name === 'void') {
        return 'nil';
      }
      return name;
    },

    generate: function(g, types, modules, targetModule) {
      fs.removeSync(path.join(g.outputDirectory, 'typhenapi'));

      g.generateFiles('lib/go/templates/core', '**/*.go', 'typhenapi/core');
      g.generate('lib/go/templates/core/message_test.hbs', 'typhenapi/core/message_test.go');

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
        g.generate('lib/go/templates/websocket/api.hbs', 'underscore:typhenapi/websocket/**/*/api.go', module);
      });

      g.files.forEach(function(file) {
        file.contents = execSync('gofmt', { input: file.contents });
      });
    }
  };
};
