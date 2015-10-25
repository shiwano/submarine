'use strict';

var path = require('path');
var fs = require('fs-extra');
var assert = require('assert');
var glob = require('glob');
var _ = require('lodash');

module.exports = function(typhen, options, helpers) {
  helpers = _.assign(helpers, {
    errorTypeName: function(symbol) {
      if (symbol.parentModule) {
        return 'TyphenApi.Type.' + typhen.helpers.upperCamelCase(symbol.ancestorModules[0].fullName) + '.Error';
      } else {
        assert(symbol.isModule, 'should be a module');
        return 'TyphenApi.Type.' + typhen.helpers.upperCamelCase(symbol.name) + '.Error';
      }
    },
    nullableToken: function(type, isOptional) {
      if (isOptional && type.isPrimitiveType &&
        (type.name === 'bool' || type.name === 'double' || type.name === 'long')) {
        return '?';
      }
    },
    typeName: function(type) {
      return (type.isPrimitiveType && type.name !== 'void') || type.isArray || type.isTypeParameter ?
        type.name :
        'TyphenApi.Type.' + typhen.helpers.upperCamelCase(type.fullName);
    }
  });

  return {
    requiredTargetModule: false,
    namespaceSeparator: '.',
    helpers: helpers,

    rename: function(symbol, name) {
      if (symbol.isArray) {
        return 'List<' + typhen.helpers.upperCamelCase(symbol.type) + '>';
      } else if (name === 'number') {
        return 'double';
      } else if (name === 'integer') {
        return 'long';
      } else if (name === 'boolean') {
        return 'bool';
      } else if (symbol.isGenericType && symbol.typeArguments.length > 0) {
        var argNames = symbol.typeArguments.map(function(t) { return helpers.typeName(t); });
        return typhen.helpers.upperCamelCase(symbol.rawName) + '<' + argNames.join(', ') + '>';
      }
      return name;
    },

    generate: function(g, types, modules) {
      glob.sync(path.join(g.outputDirectory, 'TyphenApi/Generated/**/*.cs')).forEach(function(path) {
        fs.removeSync(path);
      });

      g.generateFiles('lib/unity/templates/Core', '**/*.cs', 'TyphenApi/Generated/Core');
      g.generate('lib/unity/templates/Type/Void.cs', 'TyphenApi/Generated/Type/Void.cs');
      g.generate('lib/unity/templates/Type/RealTimeMessage.cs', 'TyphenApi/Generated/Type/RealTimeMessage.cs');

      modules.forEach(function(module) {
        if (module.functions.length > 0) {
          g.generate('lib/unity/templates/WebApi/WebApi.hbs', 'upperCamelCase:TyphenApi/Generated/WebApi/**/*.cs', module);

          if (!module.parentModule) {
            g.generateUnlessExist('lib/unity/templates/Controller/WebApiController.hbs', 'upperCamelCase:TyphenApi/Controller/WebApi/*.cs', module);
          }
        }
      });

      types.forEach(function(type) {
        switch (type.kind) {
          case typhen.SymbolKind.Enum:
            g.generate('lib/unity/templates/Type/Enum.hbs', 'upperCamelCase:TyphenApi/Generated/Type/**/*.cs', type);
            break;
          case typhen.SymbolKind.Interface:
            if (!type.isGenericType || type.typeArguments.length === 0) {
              g.generate('lib/unity/templates/Type/Class.hbs', 'upperCamelCase:TyphenApi/Generated/Type/**/*.cs', type);
            }
            break;
          case typhen.SymbolKind.ObjectType:
            g.generate('lib/unity/templates/Type/Class.hbs', 'upperCamelCase:TyphenApi/Generated/Type/**/*.cs', type);
            break;
        }
      });

      if (!options.excludeUnityFiles) {
        g.generate('lib/unity/templates/Unity/WebApiRequestSenderWWW.cs', 'TyphenApi/Generated/Core/Unity/WebApiRequestSenderWWW.cs');
        g.generate('lib/unity/templates/Unity/WebApiRequest.Unity.cs', 'TyphenApi/Generated/Core/Unity/WebApiRequest.Unity.cs');
        g.generate('lib/unity/templates/Unity/FormDataSerializer.cs', 'TyphenApi/Generated/Core/Unity/FormDataSerializer.cs');

        if (options.includeUniRxFiles) {
          g.generate('lib/unity/templates/Unity/WebApiRequest.UniRx.cs', 'TyphenApi/Generated/Core/Unity/WebApiRequest.UniRx.cs');
        }
      }

      if (!options.excludeMiniJSON) {
        g.generateUnlessExist('lib/unity/templates/Import/MiniJSON.cs', 'TyphenApi/Import/MiniJSON.cs');
      }
    }
  };
};
