/// <reference path='../tools/typhen/typhen-api/index.d.ts' />

declare module Submarine {
  type integer = TyphenApi.integer;
  type degrees = number;            // 0-360
  type timeStamp = integer;         // see http://currentmillis.com/
  type milliSeconds = integer;

  interface Error {
    code: integer;
    name: string;
    message: string;
  }

  interface User {
    id: integer;
    name: string;
  }

  interface LoggedInUser extends User {
    joinedRoom?: JoinedRoom;
  }

  interface Room {
    id: integer;
    members: User[];
    bots?: Bot[];
  }

  interface Bot {
    id: integer;
    name: string;
  }

  interface JoinedRoom extends Room {
    battleServerBaseUri: string;
    roomKey: string;
  }

  /** @noAuthRequired */
  function ping(message: string): { message: string; };
  /** @noAuthRequired */
  function signUp(name: string): { user: LoggedInUser; authToken: string; };
  /** @noAuthRequired */
  function login(authToken: string): { user: LoggedInUser; };

  function findUser(name: string): { user?: User; };

  function createRoom(): { room: JoinedRoom; };
  function getRooms(): { rooms: Room[]; }
  function joinIntoRoom(roomId: integer): { room: JoinedRoom; };
}
