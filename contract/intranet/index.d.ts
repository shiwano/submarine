/// <reference path='../index.d.ts' />

declare module Submarine.Battle {
  interface RoomMember extends User {
    id: integer;
    roomId: integer;
  }

  interface PlayableRoom {
    id: integer;
  }

  /** @noAuthRequired */
  function findRoom(roomId: integer): { room?: PlayableRoom; }
  /** @noAuthRequired */
  function findRoomMember(roomKey: string): { roomMember?: RoomMember; }
  /** @noAuthRequired */
  function closeRoom(roomId: integer): void;
}
