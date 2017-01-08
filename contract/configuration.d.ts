declare module Submarine.Configuration {
  interface Client {
    version: string;
    apiServerBaseUri: string;
  }

  interface Server {
    apiServerBaseUri: string;
    battleServerBaseUri: string;
    database: {
      host: string;
      port: integer;
      user: string;
      password: string;
    }
  }
}
