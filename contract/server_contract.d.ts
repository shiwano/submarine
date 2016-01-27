/// <reference path='./contract.d.ts' />

declare module Submarine {
  module Battle {
    interface RoomMember extends User {
      id: integer;
      roomId: integer;
    }

    interface Room {
      id: integer;
    }

    /** @noAuthRequired */
    function findRoom(room_id: integer): { room?: Room; }
    /** @noAuthRequired */
    function findRoomMember(room_key: string): { room_member?: RoomMember; }
    /** @noAuthRequired */
    function closeRoom(room_id: integer): void;
  }
}
