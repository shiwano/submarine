/// <reference path='./index.d.ts' />

declare module Submarine.Battle {
  enum ActorType {
    Submarine,
    Torpedo,
    Decoy,
    Watcher,
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
    winnerUserId?: integer;
    finishedAt: timeStamp;
  }

  interface Actor {
    id: integer;
    userId: integer;
    type: ActorType;
    movement: Movement;
    isVisible: boolean;
    submarine?: {
      isUsingPinger: boolean;
      equipment?: Equipment;
    }
  }

  interface Visibility {
    actorId: integer;
    isVisible: boolean;
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

  interface Pinger {
    actorId: integer;
    isFinished: boolean;
  }

  interface Equipment {
    actorId: integer;
    torpedos: EquipmentItem[];
    pinger: EquipmentItem;
    watcher: EquipmentItem;
  }

  interface EquipmentItem {
    cooldownStartedAt: timeStamp;
    cooldownDuration: milliSeconds;
  }

  var ping: { message: string; }
  var room: Submarine.Room;

  var now: { time: timeStamp; };
  var start: Start;
  var finish: Finish;

  var actor: Actor;
  var visibility: Visibility;
  var movement: Movement;
  var destruction: Destruction;
  var pinger: Pinger;
  var equipment: Equipment;

  var startRequest: {};
  var accelerationRequest: { direction: degrees; };
  var brakeRequest: { direction: degrees; };
  var turnRequest: { direction: degrees; };
  var pingerRequest: {};
  var torpedoRequest: {};
  var watcherRequest: {};

  var addBotRequest: {};
  var removeBotRequest: { botId: integer; };
}
