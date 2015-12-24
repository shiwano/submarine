interface integer { }

declare module Submarine {
  interface Error {
    code: integer;
    name: string;
    message: string;
  }

  interface Config {
    version: string;
    webApiServerBaseUri: string;
  }

  interface User {
    id: integer;
    name: string;
  }

  /** @noAuthRequired */
  function ping(message: string): { message: string; };
  /** @noAuthRequired */
  function signUp(name: string, password: string): { user: User; };
  /** @noAuthRequired */
  function login(name: string, password: string): { user: User; };

  function findUser(name: string): { user?: User; };

  module Battle {
    interface Ping {
      message: string;
    }

    var ping: Ping;
  }
}
