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

  type ping = TyphenApi.RTMMessage<{ message: string }>;
  type room = TyphenApi.RTMMessage<Submarine.Room>;

  type now = TyphenApi.RTMMessage<{ time: timeStamp }>;
  type start = TyphenApi.RTMMessage<Start>;
  type finish = TyphenApi.RTMMessage<Finish>;

  type actor = TyphenApi.RTMMessage<Actor>;
  type visibility = TyphenApi.RTMMessage<Visibility>;
  type movement = TyphenApi.RTMMessage<Movement>;
  type destruction = TyphenApi.RTMMessage<Destruction>;
  type pinger = TyphenApi.RTMMessage<Pinger>;
  type equipment = TyphenApi.RTMMessage<Equipment>;

  type startRequest = TyphenApi.RTMMessage<{}>;
  type accelerationRequest = TyphenApi.RTMMessage<{ direction: degrees; }>;
  type brakeRequest = TyphenApi.RTMMessage<{ direction: degrees; }>;
  type turnRequest = TyphenApi.RTMMessage<{ direction: degrees; }>;
  type pingerRequest = TyphenApi.RTMMessage<{}>;
  type torpedoRequest = TyphenApi.RTMMessage<{}>;
  type watcherRequest = TyphenApi.RTMMessage<{}>;

  type addBotRequest = TyphenApi.RTMMessage<{}>;
  type removeBotRequest = TyphenApi.RTMMessage<{botId: integer}>;
}
