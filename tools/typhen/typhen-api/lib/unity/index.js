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
      assert(type.isType, 'should be a type');
      if (isOptional && type.isPrimitiveType &&
        (type.name === 'bool' || type.name === 'double' || type.name === 'long')) {
        return '?';
      }
    },
    baseType: function(type) {
      assert(type.isType, 'should be a type');
      if (!type.baseTypes) { return null; }
      return type.baseTypes.filter(function(t) { return t.rawFullName !== 'TyphenApi.RealTimeMessage'; })[0] || null;
    },
    typeName: function(type) {
      assert(type.isType, 'should be a type');
      return (type.isPrimitiveType && type.name !== 'void') || type.isArray || type.isTypeParameter ?
        type.name :
        'TyphenApi.Type.' + typhen.helpers.upperCamelCase(type.fullName);
    },
    typeConstraint: function(typeParameters) {
      assert(_.every(typeParameters, function(t) { return t.isTypeParameter; }), 'should be type parameters');
      var code = typeParameters
        .filter(function(t) { return _.isObject(t.constraint); })
        .map(function(t) { return 'where ' + t.name + ' : ' + helpers.typeName(t.constraint); })
        .join(', ');
      return _.isEmpty(code) ? '' : ' ' + code;
    }
  });

  return {
    requiredTargetModule: false,
    namespaceSeparator: '.',
    helpers: helpers,

    rename: function(symbol, name) {
      if (symbol.isArray) {
        return 'List<' + helpers.typeName(symbol.type) + '>';
      } else if (symbol.isPrimitiveType && name === 'number') {
        return 'double';
      } else if (symbol.isPrimitiveType && name === 'integer') {
        return 'long';
      } else if (symbol.isPrimitiveType && name === 'boolean') {
        return 'bool';
      } else if (symbol.isGenericType && symbol.typeArguments.length > 0) {
        var argNames = symbol.typeArguments.map(function(t) { return helpers.typeName(t); });
        return typhen.helpers.upperCamelCase(symbol.rawName) + '<' + argNames.join(', ') + '>';
      } else if (symbol.isGenericType && symbol.typeArguments.length === 0) {
        var paramNames = symbol.typeParameters.map(function(t) { return helpers.typeName(t); });
        return typhen.helpers.upperCamelCase(symbol.rawName) + '<' + paramNames.join(', ') + '>';
      }
      return name;
    },

    generate: function(g, types, modules) {
      glob.sync(path.join(g.outputDirectory, 'TyphenApi/Generated/**/*.cs')).forEach(function(path) {
        fs.removeSync(path);
      });

      g.generateFiles('lib/unity/templates/Core', '**/*.cs', 'TyphenApi/Generated/Core');
      g.generate('lib/unity/templates/Type/Void.cs', 'TyphenApi/Generated/Type/Void.cs');

      modules.forEach(function(module) {
        if (helpers.isWebApiModule(module)) {
          g.generate('lib/unity/templates/WebApi/WebApiBase.hbs', 'upperCamelCase:TyphenApi/Generated/WebApi/**/*.cs', module);

          if (!module.parentModule) {
            g.generateUnlessExist('lib/unity/templates/WebApi/WebApi.hbs', 'upperCamelCase:TyphenApi/WebApi/*.cs', module);
          }
        }

        if (helpers.isWebSocketApiModule(module)) {
          g.generate('lib/unity/templates/WebSocketApi/WebSocketApi.hbs', 'upperCamelCase:TyphenApi/Generated/WebSocketApi/**/*.cs', module);

          if (!options.excludeUnityFiles && options.includeUniRxFiles) {
            g.generate('lib/unity/templates/Unity/WebSocketApi.UniRx.hbs', 'upperCamelCase:TyphenApi/Generated/WebSocketApi/**/*.UniRx.cs', module);
          }

          if (!module.parentModule) {
            g.generate('lib/unity/templates/WebSocketSession/WebSocketSessionBase.hbs', 'upperCamelCase:TyphenApi/Generated/WebSocketSession/*.cs', module);
            g.generateUnlessExist('lib/unity/templates/WebSocketSession/WebSocketSession.hbs', 'upperCamelCase:TyphenApi/WebSocketSession/*.cs', module);
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
        g.generate('lib/unity/templates/Unity/WebApiRequestSenderUnity.cs', 'TyphenApi/Generated/Core/Unity/WebApiRequestSenderUnity.cs');
        g.generate('lib/unity/templates/Unity/WebApiRequest.Unity.cs', 'TyphenApi/Generated/Core/Unity/WebApiRequest.Unity.cs');

        if (options.includeUniRxFiles) {
          g.generate('lib/unity/templates/Unity/WebApiRequest.UniRx.cs', 'TyphenApi/Generated/Core/Unity/WebApiRequest.UniRx.cs');
        }
      }
    }
  };
};
