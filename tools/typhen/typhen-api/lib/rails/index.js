'use strict';

var path = require('path');
var fs = require('fs-extra');

module.exports = function(typhen) {
  var helpers = {
    controllerName: function(func) {
      return typhen.helpers.upperCamelCase(func.fullName).split('::').slice(1).join('::');
    },
    uriPath: function(func) {
      var inflection = func.ancestorModules[0].tagTable.uriInflection;
      var helperName = inflection ? inflection.value : 'underscore';
      return typhen.helpers[helperName](func.fullName).split('::').slice(1).join('/');
    },
    uriSuffix: function(symbol) {
      return symbol.ancestorModules[0].tagTable.uriSuffix;
    },
    controllerPath: function(symbol) {
      return typhen.helpers.underscore(symbol.fullName).split('::').slice(1).join('/');
    },
    typeName: function(type) {
      var name = type.isPrimitiveType || type.isArray ? type.name : 'TyphenApi::Model::' + type.fullName;
      return name === 'nil' ? name : typhen.helpers.upperCamelCase(name);
    },
    responsePropertyName: function(symbol) {
      var inflection = symbol.ancestorModules[0].tagTable.responsePropertyInflection;
      var helperName = inflection ? inflection.value : 'underscore';
      return typhen.helpers[helperName](symbol.name);
    }
  };

  return {
    requiredTargetModule: true,
    namespaceSeparator: '::',
    helpers: helpers,

    rename: function(symbol, name) {
      if (symbol.kind === typhen.SymbolKind.Array) {
        return 'Array[' + typhen.helpers.upperCamelCase(symbol.type) + ']';
      } else if (name === 'void') {
        return 'nil';
      }
      return name;
    },

    generate: function(generator, types, modules, targetModule) {
      fs.removeSync(path.join(generator.outputDirectory, 'lib/typhen_api'));

      generator.generateUnlessExist('lib/rails/templates/controller/validation.hbs', 'app/controllers/concerns/typhen_api_validation.rb');
      generator.generate('lib/rails/templates/typhen_api.hbs', 'lib/typhen_api/typhen_api.rb');
      generator.generate('lib/rails/templates/controller.hbs', 'lib/typhen_api/typhen_api/controller.rb');
      generator.generate('lib/rails/templates/model.hbs', 'lib/typhen_api/typhen_api/model.rb');

      var functions = modules.filter(function(m) { return m === targetModule || m.ancestorModules.indexOf(targetModule) > -1; })
        .map(function(module) { return module.functions; })
        .reduce(function(a, b) { return a.concat(b); });

      generator.generate('lib/rails/templates/routes.hbs', 'lib/typhen_api/typhen_api/routes.rb', functions);

      functions.forEach(function(func) {
        var controllerPath = 'app/controllers/' + helpers.controllerPath(func) + '_controller.rb';
        generator.generateUnlessExist('lib/rails/templates/controller/app_controller.hbs', controllerPath, func);
        generator.generate('lib/rails/templates/controller/controller.hbs', 'underscore:lib/typhen_api/typhen_api/controller/**/*.rb', func);
        generator.generateUnlessExist('lib/rails/templates/controller/module.hbs', 'underscore:lib/typhen_api/typhen_api/controller/**/*.rb', func.parentModule);

        if (func.parentModule !== targetModule) {
          var modulePath = 'app/controllers/' + helpers.controllerPath(func.parentModule) + '.rb';
          generator.generateUnlessExist('lib/rails/templates/controller/app_module.hbs', modulePath, func.parentModule);
        }
      });

      types.forEach(function(type) {
        switch (type.kind) {
          case typhen.SymbolKind.Enum:
            generator.generate('lib/rails/templates/model/enum.hbs', 'underscore:lib/typhen_api/typhen_api/model/**/*.rb', type);
            break;
          case typhen.SymbolKind.Interface:
            if (!type.isGenericType || type.typeArguments.length > 0) {
              generator.generate('lib/rails/templates/model/object.hbs', 'underscore:lib/typhen_api/typhen_api/model/**/*.rb', type);
            }
            break;
          case typhen.SymbolKind.ObjectType:
            generator.generate('lib/rails/templates/model/object.hbs', 'underscore:lib/typhen_api/typhen_api/model/**/*.rb', type);
            break;
          default:
            return;
        }
        generator.generateUnlessExist('lib/rails/templates/model/module.hbs', 'underscore:lib/typhen_api/typhen_api/model/**/*.rb', type.parentModule);
      });
    }
  };
};
