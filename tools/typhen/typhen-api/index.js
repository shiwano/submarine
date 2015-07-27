'use strict';

var path = require('path');
var assert = require('assert');
var _ = require('lodash');

function isWebApiModule(module) {
  if (module.functions.length > 0) {
    return true;
  } else if (module.modules) {
    return _.any(module.modules, function(m) { return isWebApiModule(m); });
  } else {
    return false;
  }
}

module.exports = function(typhen, options) {
  assert(options, 'options is empty');
  assert(typeof options.templateName === 'string', 'options.templateName is required');

  var helpers = {
    method: function(symbol) {
      assert(symbol.isSignature, 'should be a function call signature');
      var method = symbol.tagTable.method ? symbol.tagTable.method.value : 'post';
      assert(_.includes(['post', 'get', 'delete', 'put'], method), 'unsupported HTTP method: ' + method);
      return method;
    },
    uriPath: function(symbol) {
      assert(symbol.isFunction || symbol.isSignature, 'should be a function or signature');
      var inflection = symbol.ancestorModules[0].tagTable.uriInflection;
      var helperName = inflection ? inflection.value : 'underscore';
      assert(_.includes(['underscore', 'upperCamelCase', 'lowerCamelCase'], helperName), 'unsupported inflection type: ' + helperName);
      return typhen.helpers[helperName](symbol.fullName).split(template.namespaceSeparator).slice(1).join('/');
    },
    uriSuffix: function(symbol) {
      assert(symbol.isFunction || symbol.isSignature, 'should be a function or signature');
      return symbol.ancestorModules[0].tagTable.uriSuffix;
    },
    serializablePropertyName: function(symbol) {
      assert(symbol.isProperty || symbol.isParameter, 'should be a property or function parameter');
      var inflection = symbol.ancestorModules[0].tagTable.serializablePropertyInflection;
      var helperName = inflection ? inflection.value : 'underscore';
      assert(_.includes(['underscore', 'upperCamelCase', 'lowerCamelCase'], helperName), 'unsupported inflection type: ' + helperName);
      return typhen.helpers[helperName](symbol.name);
    }
  };

  var template = require(path.join(__dirname, 'lib', options.templateName, 'index.js'))(typhen, options, helpers);

  if (template.requiredTargetModule) {
    assert(typeof options.targetModule === 'string', 'options.targetModule is required');
  }

  return typhen.createPlugin({
    pluginDirectory: __dirname,
    customPrimitiveTypes: ['integer'],
    namespaceSeparator: template.namespaceSeparator,
    disallow: {
      any: true,
      overload: true,
      unionType: true,
      tuple: true,
      anonymousFunction: true
    },
    handlebarsOptions: {
      data: options,
      helpers: template.helpers
    },

    rename: template.rename,

    generate: function(generator, types, modules) {
      var targetModule = null;

      if (template.requiredTargetModule) {
        targetModule = modules.filter(function(m) { return m.fullName === options.targetModule; })[0];
        assert(targetModule, options.targetModule + ' module is not found');
      }

      var filteredTypes = types.filter(function(t) { return !t.tagTable.internal; });
      var filteredModules = modules.filter(function(m) { return !m.isGlobalModule; });

      filteredModules.forEach(function(module) {
        if (module.parentModule === null && isWebApiModule(module)) {
          var errorType = module.types.filter(function(t) { return t.name === 'Error'; })[0];
          assert(errorType, 'Undefined the Error type in ' + module.name + ' module');
        }
      });

      return template.generate(generator, filteredTypes, filteredModules, targetModule);
    }
  });
};
