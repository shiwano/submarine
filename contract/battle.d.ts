/// <reference path='./contract.d.ts' />

declare module Submarine.Battle {
  enum ActorType {
    Submarine,
    Torpedo,
    Decoy,
    Lookout,
  }

  interface Point {
    x: number;
    y: number;
  }

  interface Accelerator {
    maxSpeed: number;
    duration: milliSeconds;
    startRate: number;
    isAccelerating: boolean;
  }

  interface Start {
    startedAt: timeStamp;
  }

  interface Finish {
    winnerUserId: integer;
    finishedAt: timeStamp;
  }

  interface Actor {
    id: integer;
    userId: integer;
    type: ActorType;
    movement: Movement;
  }

  interface Movement {
    actorId: integer;
    position: Point;
    direction: degrees;
    movedAt: timeStamp;
    accelerator?: Accelerator;
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
  var turnRequest: { direction: degrees; };
  var pingerRequest: {};
  var torpedoRequest: {};
}
