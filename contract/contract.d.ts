/// <reference path='./battle.d.ts' />

interface integer { }

declare module Submarine {
  type timeStamp = integer; // see http://currentmillis.com/

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
    name: string;
  }

  interface LoggedInUser extends User {
    joinedRoom?: JoinedRoom;
  }

  interface Room {
    id: integer;
    members: User[];
  }

  interface JoinedRoom extends Room {
    battleServerBaseUri: string;
    roomKey: string;
  }

  /** @noAuthRequired */
  function ping(message: string): { message: string; };
  /** @noAuthRequired */
  function signUp(name: string, password: string): { user: LoggedInUser; };
  /** @noAuthRequired */
  function login(name: string, password: string): { user: LoggedInUser; };

  function findUser(name: string): { user?: User; };

  function createRoom(): { room: JoinedRoom; };
  function getRooms(): { rooms: Room[]; }
  function joinIntoRoom(roomId: integer): { room: JoinedRoom; };

  module Battle {
    var ping: { message: string; }
    var room: Submarine.Room;
  }
}
