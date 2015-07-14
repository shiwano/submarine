'use strict';

var path = require('path');
var assert = require('assert');

module.exports = function(typhen, options) {
  assert(options, 'options is empty');
  assert(typeof options.templateName === 'string', 'options.templateName is required');

  var template = require(path.join(__dirname, 'lib', options.templateName, 'index.js'))(typhen, options);

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

      modules.forEach(function(module) {
        if (module.parentModule === null || !module.parentModule.isGlobalModule) { return; }
        var errorType = module.types.filter(function(t) { return t.name === 'Error'; })[0];
        assert(errorType, 'Not found Error type in ' + module.name + ' module');
      });

      var filteredTypes = types.filter(function(t) { return !t.tagTable.internal; });
      return template.generate(generator, filteredTypes, modules, targetModule);
    }
  });
};
