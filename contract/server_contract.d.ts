/// <reference path='./contract.d.ts' />

declare module Submarine {
  module Battle {
    interface RoomMember extends User {
      roomKey: string;
    }

    interface Room {
      id: integer;
      members: RoomMember[];
    }

    function findRoom(room_id: integer): { room?: Room; }
    function closeRoom(room_id: integer): void;
  }
}
