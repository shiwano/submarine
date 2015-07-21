'use strict';

var path = require('path');
var fs = require('fs-extra');
var assert = require('assert');
var glob = require('glob');

module.exports = function(typhen, options) {
  var helpers = {
    errorTypeName: function(symbol) {
      if (symbol.parentModule) {
        return 'TyphenApi.Type.' + typhen.helpers.upperCamelCase(symbol.ancestorModules[0].fullName) + '.Error';
      } else {
        assert(symbol.isModule, 'not module');
        return 'TyphenApi.Type.' + typhen.helpers.upperCamelCase(symbol.name) + '.Error';
      }
    },
    nullableToken: function(type, isOptional) {
      if (isOptional && type.isPrimitiveType &&
        (type.name === 'bool' || type.name === 'float' || type.name === 'int')) {
        return '?';
      }
    },
    typeName: function(type) {
      return (type.isPrimitiveType && type.name !== 'void') || type.isArray || type.isTypeParameter ?
        type.name :
        'TyphenApi.Type.' + typhen.helpers.upperCamelCase(type.fullName);
    },
    method: function(func) {
      var name = func.tagTable.method ? func.tagTable.method : 'post';
      return typhen.helpers.upperCamelCase(name);
    },
    uriPath: function(func) {
      var inflection = func.ancestorModules[0].tagTable.uriInflection;
      var helperName = inflection ? inflection.value : 'underscore';
      return typhen.helpers[helperName](func.fullName).split('.').slice(1).join('/');
    },
    uriSuffix: function(symbol) {
      return symbol.ancestorModules[0].tagTable.uriSuffix;
    },
    serializablePropertyName: function(symbol) {
      var inflection = symbol.ancestorModules[0].tagTable.serializablePropertyInflection;
      var helperName = inflection ? inflection.value : 'underscore';
      return typhen.helpers[helperName](symbol.name);
    }
  };

  return {
    requiredTargetModule: false,
    namespaceSeparator: '.',
    helpers: helpers,

    rename: function(symbol, name) {
      if (symbol.isArray) {
        return 'List<' + typhen.helpers.upperCamelCase(symbol.type) + '>';
      } else if (name === 'number') {
        return 'float';
      } else if (name === 'integer') {
        return 'int';
      } else if (name === 'boolean') {
        return 'bool';
      } else if (symbol.isGenericType && symbol.typeArguments.length > 0) {
        var argNames = symbol.typeArguments.map(function(t) { return helpers.typeName(t); });
        return typhen.helpers.upperCamelCase(symbol.rawName) + '<' + argNames.join(', ') + '>';
      }
      return name;
    },

    generate: function(generator, types, modules) {
      glob.sync(path.join(generator.outputDirectory, 'TyphenApi/Generated/**/*.cs')).forEach(function(path) {
        fs.removeSync(path);
      });

      generator.generateFiles('lib/unity/templates/Core', '**/*.cs', 'TyphenApi/Generated/Core');
      generator.generate('lib/unity/templates/Type/Void.cs', 'TyphenApi/Generated/Type/Void.cs');

      modules.forEach(function(module) {
        if (!module.isGlobalModule) {
          generator.generate('lib/unity/templates/WebApi/WebApi.hbs', 'upperCamelCase:TyphenApi/Generated/WebApi/**/*.cs', module);

          if (!module.parentModule) {
            generator.generateUnlessExist('lib/unity/templates/Controller/WebApiController.hbs', 'upperCamelCase:TyphenApi/Controller/*.cs', module);
          }
        }
      });

      var apiModules = modules.filter(function(module) { return !module.isGlobalModule && !module.parentModule; });
      generator.generateUnlessExist('lib/unity/templates/Api.hbs', 'TyphenApi/Generated/Api.cs', { modules: apiModules });

      types.forEach(function(type) {
        switch (type.kind) {
          case typhen.SymbolKind.Enum:
            generator.generate('lib/unity/templates/Type/Enum.hbs', 'upperCamelCase:TyphenApi/Generated/Type/**/*.cs', type);
            break;
          case typhen.SymbolKind.Interface:
            if (!type.isGenericType || type.typeArguments.length === 0) {
              generator.generate('lib/unity/templates/Type/Class.hbs', 'upperCamelCase:TyphenApi/Generated/Type/**/*.cs', type);
            }
            break;
          case typhen.SymbolKind.ObjectType:
            generator.generate('lib/unity/templates/Type/Class.hbs', 'upperCamelCase:TyphenApi/Generated/Type/**/*.cs', type);
            break;
        }
      });

      if (!options.excludeUnityFiles) {
        generator.generate('lib/unity/templates/Unity/WebApiRequestSenderWWW.cs', 'TyphenApi/Generated/Core/Unity/WebApiRequestSenderWWW.cs');
        generator.generate('lib/unity/templates/Unity/WebApiRequest.Unity.cs', 'TyphenApi/Generated/Core/Unity/WebApiRequest.Unity.cs');

        if (options.includeUniRxFiles) {
          generator.generate('lib/unity/templates/Unity/WebApiRequest.UniRx.cs', 'TyphenApi/Generated/Core/Unity/WebApiRequest.UniRx.cs');
        }
      }

      if (!options.excludeMiniJSON) {
        generator.generateUnlessExist('lib/unity/templates/Import/MiniJSON.cs', 'TyphenApi/Import/MiniJSON.cs');
      }
    }
  };
};
