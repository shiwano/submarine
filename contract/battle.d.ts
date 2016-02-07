/// <reference path='./contract.d.ts' />

declare module Submarine.Battle {
  enum ActorType {
    Submarine,
    Torpedo,
    Decoy,
    Lookout,
  }

  interface Vector {
    x: number;
    y: number;
  }

  interface Start {
    startedAt: timeStamp;
    controllableActorId: integer;
  }

  interface Finish {
    hasWon: boolean;
    finishedAt: timeStamp;
  }

  interface Actor {
    id: integer;
    userId: integer;
    type: ActorType;
    position: Vector;
  }

  interface Movement {
    actorId: integer;
    position: Vector;
    velocity: Vector;
    movedAt: timeStamp;
  }

  interface Destruction {
    actorId: integer;
  }

  var ping: { message: string; }
  var room: Submarine.Room;

  var now: { time: timeStamp; };
  var start: Start;
  var finish: Finish;

  var actor: Actor;
  var movement: Movement;
  var destruction: Destruction;

  var accelerationRequest: {};
  var brakeRequest: {};
  var turnRequest: { direction: Vector; };
  var pingerRequest: {};
  var torpedoRequest: {};
}
