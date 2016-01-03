/// <reference path='./contract.d.ts' />

declare module Submarine {
  module Battle {
    interface RoomMember extends User {
      id: integer;
      roomKey: string;
    }

    interface Room {
      id: integer;
      members: RoomMember[];
    }

    /** @noAuthRequired */
    function findRoom(room_id: integer): { room?: Room; }
    /** @noAuthRequired */
    function closeRoom(room_id: integer): void;
  }
}
