/// <reference path='./contract.d.ts' />

declare module Submarine {
  module Battle {
    interface RoomMember {
      roomKey: RoomKey;
      user: User;
    }

    interface Room {
      id: integer;
      members: RoomMember[];
    }

    function findRoom(room_id: integer): { room?: Room; }
    function closeRoom(room_id: integer): void;
  }
}
