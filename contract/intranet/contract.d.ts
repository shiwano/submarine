/// <reference path='../contract.d.ts' />

declare module Submarine.Battle {
  interface RoomMember extends User {
    id: integer;
    roomId: integer;
  }

  interface Room {
    id: integer;
  }

  /** @noAuthRequired */
  function findRoom(roomId: integer): { room?: Room; }
  /** @noAuthRequired */
  function findRoomMember(roomKey: string): { roomMember?: RoomMember; }
  /** @noAuthRequired */
  function closeRoom(roomId: integer): void;
}
