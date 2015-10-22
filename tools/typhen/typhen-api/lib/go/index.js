'use strict';

var path = require('path');
var assert = require('assert');
var fs = require('fs-extra');
var execSync = require('child_process').execSync;
var _ = require('lodash');

module.exports = function(typhen, options, helpers) {
  assert(options.importBasePath, 'options.importBasePath is empty');

  helpers = _.assign(helpers, {
    requiredModules: function(type) {
      return _.chain(type.properties)
        .map(function(property) {
          var type = property.type;

          if (type.parentModule === null ||
              type.parentModule === type.parentModule ||
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
    realTimeMessageDispacherModules: function(module) {
      return _.chain(module.modules)
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
      fs.removeSync(path.join(g.outputDirectory, 'typhen_api'));

      g.generate('lib/go/templates/core/message.go', 'typhen_api/core/message.go');
      g.generate('lib/go/templates/core/type.go', 'typhen_api/core/type.go');
      g.generate('lib/go/templates/core/serializer.go', 'typhen_api/core/serializer.go');

      types.forEach(function(type) {
        switch (type.kind) {
          case typhen.SymbolKind.Enum:
            g.generate('lib/go/templates/type/enum.hbs', 'underscore:typhen_api/type/**/*.go', type);
            break;
          case typhen.SymbolKind.Interface:
            if (!type.isGenericType || type.typeArguments.length > 0) {
              g.generate('lib/go/templates/type/object.hbs', 'underscore:typhen_api/type/**/*.go', type);
            }
            break;
          case typhen.SymbolKind.ObjectType:
            g.generate('lib/go/templates/type/object.hbs', 'underscore:typhen_api/type/**/*.go', type);
            break;
        }
      });

      modules.forEach(function(module) {
        g.generate('lib/go/templates/messagedispatcher/messagedispatcher.hbs', 'underscore:typhen_api/messagedispatcher/**/*/messagedispatcher.go', module);
      });

      g.files.forEach(function(file) {
        file.contents = execSync('gofmt', { input: file.contents });
      });
    }
  };
};
