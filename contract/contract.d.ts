/// <reference path='./battle.d.ts' />

interface integer {}

declare module Submarine {
  type degrees = number;       // 0-360
  type timeStamp = integer;    // see http://currentmillis.com/
  type milliSeconds = integer;

  interface Error {
    code: integer;
    name: string;
    message: string;
  }

  interface Config {
    version: string;
    apiServerBaseUri: string;
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
  }

  interface JoinedRoom extends Room {
    battleServerBaseUri: string;
    roomKey: string;
  }

  /** @noAuthRequired */
  function ping(message: string): { message: string; };
  /** @noAuthRequired */
  function signUp(): { user: LoggedInUser; authToken: string; accessToken: string; };
  /** @noAuthRequired */
  function login(authToken: string): { user: LoggedInUser; accessToken: string; };

  function findUser(name: string): { user?: User; };

  function createRoom(): { room: JoinedRoom; };
  function getRooms(): { rooms: Room[]; }
  function joinIntoRoom(roomId: integer): { room: JoinedRoom; };
}
