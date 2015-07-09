'use strict';

module.exports = function(typhen) {
  var pluginForRails = typhen.loadPlugin(__dirname + '/typhen-api/index', {
    templateName: 'server/rails',
    targetModule: 'Submarine'
  });

  return typhen.run({
    plugin: pluginForRails,
    src: __dirname + '/../../contract/main.d.ts',
    dest: __dirname + '/../../server',
    typingDirectory: __dirname + '/../../contract',
    defaultLibFileName: __dirname + '/../../contract/lib.typhenApi.d.ts'
  });
};
