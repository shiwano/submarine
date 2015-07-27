'use strict';

module.exports = function(typhen) {
  var pluginForRails = typhen.loadPlugin(__dirname + '/typhen-api/index', {
    templateName: 'rails',
    targetModule: 'Submarine'
  });
  var pluginForUnity = typhen.loadPlugin(__dirname + '/typhen-api/index', {
    templateName: 'unity',
    includeUniRxFiles: true
  });

  return typhen.run({
    plugin: pluginForRails,
    src: __dirname + '/../../contract/main.d.ts',
    dest: __dirname + '/../../server',
    typingDirectory: __dirname + '/../../contract',
    defaultLibFileName: __dirname + '/../../contract/lib.typhenApi.d.ts'
  }).then(function() {
    return typhen.run({
      plugin: pluginForUnity,
      src: [
        __dirname + '/../../contract/main.d.ts',
        __dirname + '/../../contract/client.d.ts'
      ],
      dest: __dirname + '/../../client/Assets/Scripts',
      typingDirectory: __dirname + '/../../contract',
      defaultLibFileName: __dirname + '/../../contract/lib.typhenApi.d.ts'
    });
  });
};
