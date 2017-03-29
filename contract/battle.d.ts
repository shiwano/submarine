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

  interface EquipmentItem {
    cooldownStartedAt: timeStamp;
    cooldownDuration: milliSeconds;
  }

  interface Ping extends TyphenApi.RealTimeMessage {
    message: string;
  }

  interface Room extends TyphenApi.RealTimeMessage, Submarine.Room { }

  interface Now extends TyphenApi.RealTimeMessage {
    time: timeStamp
  }

  interface Start extends TyphenApi.RealTimeMessage {
    startedAt: timeStamp;
  }

  interface Finish extends TyphenApi.RealTimeMessage {
    winnerUserId?: integer;
    finishedAt: timeStamp;
  }

  interface Actor extends TyphenApi.RealTimeMessage {
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

  interface Visibility extends TyphenApi.RealTimeMessage {
    actorId: integer;
    isVisible: boolean;
    movement: Movement;
  }

  interface Movement extends TyphenApi.RealTimeMessage {
    actorId: integer;
    position: Point;
    direction: degrees;
    movedAt: timeStamp;
    accelerator?: Accelerator;
  }

  interface Destruction extends TyphenApi.RealTimeMessage {
    actorId: integer;
  }

  interface Pinger extends TyphenApi.RealTimeMessage {
    actorId: integer;
    isFinished: boolean;
  }

  interface Equipment extends TyphenApi.RealTimeMessage {
    actorId: integer;
    torpedos: EquipmentItem[];
    pinger: EquipmentItem;
    watcher: EquipmentItem;
  }

  interface StartRequest extends TyphenApi.RealTimeMessage { }
  interface AccelerationRequest extends TyphenApi.RealTimeMessage { direction: degrees; }
  interface BrakeRequest extends TyphenApi.RealTimeMessage { direction: degrees; }
  interface TurnRequest extends TyphenApi.RealTimeMessage { direction: degrees; }
  interface PingerRequest extends TyphenApi.RealTimeMessage { }
  interface TorpedoRequest extends TyphenApi.RealTimeMessage { }
  interface WatcherRequest extends TyphenApi.RealTimeMessage { }
  interface AddBotRequest extends TyphenApi.RealTimeMessage { }
  interface RemoveBotRequest extends TyphenApi.RealTimeMessage { botId: integer }
}
