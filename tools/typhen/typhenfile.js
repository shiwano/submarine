'use strict';

module.exports = function(typhen) {
  return typhen.run({
    plugin: typhen.loadPlugin(__dirname + '/typhen-api/index', {
      templateName: 'rails',
      targetModule: 'Submarine'
    }),
    src: __dirname + '/../../contract/main.d.ts',
    dest: __dirname + '/../../server/api',
    typingDirectory: __dirname + '/../../contract',
    defaultLibFileName: __dirname + '/../../contract/lib.typhenApi.d.ts'
  }).then(function() {
    return typhen.run({
      plugin: typhen.loadPlugin(__dirname + '/typhen-api/index', {
        templateName: 'unity',
        includeUniRxFiles: true
      }),
      src: __dirname + '/../../contract/main.d.ts',
      dest: __dirname + '/../../client/Assets/Scripts',
      typingDirectory: __dirname + '/../../contract',
      defaultLibFileName: __dirname + '/../../contract/lib.typhenApi.d.ts'
    });
  }).then(function() {
    return typhen.run({
      plugin: typhen.loadPlugin(__dirname + '/typhen-api/index', {
        templateName: 'go',
        importBasePath: 'app/typhenapi',
      }),
      src: __dirname + '/../../contract/main.d.ts',
      dest: __dirname + '/../../server/battle/src/app',
      typingDirectory: __dirname + '/../../contract',
      defaultLibFileName: __dirname + '/../../contract/lib.typhenApi.d.ts'
    });
  });
};
