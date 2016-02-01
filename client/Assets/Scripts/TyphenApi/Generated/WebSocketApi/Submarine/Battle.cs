// This file was generated by typhen-api

using System;
using System.Collections.Generic;

namespace TyphenApi.WebSocketApi.Parts.Submarine
{
    public partial class Battle : TyphenApi.IWebSocketApi
    {
        public enum MessageType
        {
            Ping = -1287020902,
            Room = -1286955548,
            Now = -1565539188,
            Start = -1240750730,
            Finish = -162791524,
            Actor = -1257891252,
            Movement = 1298310360,
            Destruction = -1118469016,
            AccelerationRequest = -710337400,
            BrakeRequest = 1492486768,
            PingerRequest = 110864488,
            ActorRequest = 275126852,
        }

        readonly IWebSocketSession session;

        public event Action<TyphenApi.Type.Submarine.Battle.PingObject> OnPingReceive;
        public event Action<TyphenApi.Type.Submarine.Room> OnRoomReceive;
        public event Action<TyphenApi.Type.Submarine.Battle.NowObject> OnNowReceive;
        public event Action<TyphenApi.Type.Submarine.Battle.Start> OnStartReceive;
        public event Action<TyphenApi.Type.Submarine.Battle.Finish> OnFinishReceive;
        public event Action<TyphenApi.Type.Submarine.Battle.Actor> OnActorReceive;
        public event Action<TyphenApi.Type.Submarine.Battle.Movement> OnMovementReceive;
        public event Action<TyphenApi.Type.Submarine.Battle.Destruction> OnDestructionReceive;
        public event Action<TyphenApi.Type.Submarine.Battle.AccelerationRequestObject> OnAccelerationRequestReceive;
        public event Action<TyphenApi.Type.Submarine.Battle.BrakeRequestObject> OnBrakeRequestReceive;
        public event Action<TyphenApi.Type.Submarine.Battle.PingerRequestObject> OnPingerRequestReceive;
        public event Action<TyphenApi.Type.Submarine.Battle.ActorRequestObject> OnActorRequestReceive;


        public Battle(IWebSocketSession session)
        {
            this.session = session;

        }

        public void SendPing(TyphenApi.Type.Submarine.Battle.PingObject ping)
        {
            session.Send((int)MessageType.Ping, ping);
        }

        public void SendPing(string message)
        {
            session.Send((int)MessageType.Ping, new TyphenApi.Type.Submarine.Battle.PingObject()
            {
                Message = message,
            });
        }
        public void SendRoom(TyphenApi.Type.Submarine.Room room)
        {
            session.Send((int)MessageType.Room, room);
        }

        public void SendRoom(long id, List<TyphenApi.Type.Submarine.User> members)
        {
            session.Send((int)MessageType.Room, new TyphenApi.Type.Submarine.Room()
            {
                Id = id,
                Members = members,
            });
        }
        public void SendNow(TyphenApi.Type.Submarine.Battle.NowObject now)
        {
            session.Send((int)MessageType.Now, now);
        }

        public void SendNow(long time)
        {
            session.Send((int)MessageType.Now, new TyphenApi.Type.Submarine.Battle.NowObject()
            {
                Time = time,
            });
        }
        public void SendStart(TyphenApi.Type.Submarine.Battle.Start start)
        {
            session.Send((int)MessageType.Start, start);
        }

        public void SendStart(long startedAt)
        {
            session.Send((int)MessageType.Start, new TyphenApi.Type.Submarine.Battle.Start()
            {
                StartedAt = startedAt,
            });
        }
        public void SendFinish(TyphenApi.Type.Submarine.Battle.Finish finish)
        {
            session.Send((int)MessageType.Finish, finish);
        }

        public void SendFinish(bool hasWon, long finishedAt)
        {
            session.Send((int)MessageType.Finish, new TyphenApi.Type.Submarine.Battle.Finish()
            {
                HasWon = hasWon,
                FinishedAt = finishedAt,
            });
        }
        public void SendActor(TyphenApi.Type.Submarine.Battle.Actor actor)
        {
            session.Send((int)MessageType.Actor, actor);
        }

        public void SendActor(long id, long userId, TyphenApi.Type.Submarine.Battle.ActorType type, TyphenApi.Type.Submarine.Battle.Vector position)
        {
            session.Send((int)MessageType.Actor, new TyphenApi.Type.Submarine.Battle.Actor()
            {
                Id = id,
                UserId = userId,
                Type = type,
                Position = position,
            });
        }
        public void SendMovement(TyphenApi.Type.Submarine.Battle.Movement movement)
        {
            session.Send((int)MessageType.Movement, movement);
        }

        public void SendMovement(long actorId, TyphenApi.Type.Submarine.Battle.Vector position, TyphenApi.Type.Submarine.Battle.Vector velocity, long movedAt)
        {
            session.Send((int)MessageType.Movement, new TyphenApi.Type.Submarine.Battle.Movement()
            {
                ActorId = actorId,
                Position = position,
                Velocity = velocity,
                MovedAt = movedAt,
            });
        }
        public void SendDestruction(TyphenApi.Type.Submarine.Battle.Destruction destruction)
        {
            session.Send((int)MessageType.Destruction, destruction);
        }

        public void SendDestruction(long actorId)
        {
            session.Send((int)MessageType.Destruction, new TyphenApi.Type.Submarine.Battle.Destruction()
            {
                ActorId = actorId,
            });
        }
        public void SendAccelerationRequest(TyphenApi.Type.Submarine.Battle.AccelerationRequestObject accelerationRequest)
        {
            session.Send((int)MessageType.AccelerationRequest, accelerationRequest);
        }

        public void SendAccelerationRequest()
        {
            session.Send((int)MessageType.AccelerationRequest, new TyphenApi.Type.Submarine.Battle.AccelerationRequestObject()
            {
            });
        }
        public void SendBrakeRequest(TyphenApi.Type.Submarine.Battle.BrakeRequestObject brakeRequest)
        {
            session.Send((int)MessageType.BrakeRequest, brakeRequest);
        }

        public void SendBrakeRequest()
        {
            session.Send((int)MessageType.BrakeRequest, new TyphenApi.Type.Submarine.Battle.BrakeRequestObject()
            {
            });
        }
        public void SendPingerRequest(TyphenApi.Type.Submarine.Battle.PingerRequestObject pingerRequest)
        {
            session.Send((int)MessageType.PingerRequest, pingerRequest);
        }

        public void SendPingerRequest()
        {
            session.Send((int)MessageType.PingerRequest, new TyphenApi.Type.Submarine.Battle.PingerRequestObject()
            {
            });
        }
        public void SendActorRequest(TyphenApi.Type.Submarine.Battle.ActorRequestObject actorRequest)
        {
            session.Send((int)MessageType.ActorRequest, actorRequest);
        }

        public void SendActorRequest(TyphenApi.Type.Submarine.Battle.ActorType type)
        {
            session.Send((int)MessageType.ActorRequest, new TyphenApi.Type.Submarine.Battle.ActorRequestObject()
            {
                Type = type,
            });
        }

        public TyphenApi.TypeBase DispatchMessageEvent(int messageType, byte[] messageData)
        {
            switch ((MessageType)messageType)
            {
                case MessageType.Ping:
                {
                    var message = session.MessageDeserializer.Deserialize<TyphenApi.Type.Submarine.Battle.PingObject>(messageData);

                    if (OnPingReceive != null)
                    {
                        OnPingReceive(message);
                    }

                    return message;
                }
                case MessageType.Room:
                {
                    var message = session.MessageDeserializer.Deserialize<TyphenApi.Type.Submarine.Room>(messageData);

                    if (OnRoomReceive != null)
                    {
                        OnRoomReceive(message);
                    }

                    return message;
                }
                case MessageType.Now:
                {
                    var message = session.MessageDeserializer.Deserialize<TyphenApi.Type.Submarine.Battle.NowObject>(messageData);

                    if (OnNowReceive != null)
                    {
                        OnNowReceive(message);
                    }

                    return message;
                }
                case MessageType.Start:
                {
                    var message = session.MessageDeserializer.Deserialize<TyphenApi.Type.Submarine.Battle.Start>(messageData);

                    if (OnStartReceive != null)
                    {
                        OnStartReceive(message);
                    }

                    return message;
                }
                case MessageType.Finish:
                {
                    var message = session.MessageDeserializer.Deserialize<TyphenApi.Type.Submarine.Battle.Finish>(messageData);

                    if (OnFinishReceive != null)
                    {
                        OnFinishReceive(message);
                    }

                    return message;
                }
                case MessageType.Actor:
                {
                    var message = session.MessageDeserializer.Deserialize<TyphenApi.Type.Submarine.Battle.Actor>(messageData);

                    if (OnActorReceive != null)
                    {
                        OnActorReceive(message);
                    }

                    return message;
                }
                case MessageType.Movement:
                {
                    var message = session.MessageDeserializer.Deserialize<TyphenApi.Type.Submarine.Battle.Movement>(messageData);

                    if (OnMovementReceive != null)
                    {
                        OnMovementReceive(message);
                    }

                    return message;
                }
                case MessageType.Destruction:
                {
                    var message = session.MessageDeserializer.Deserialize<TyphenApi.Type.Submarine.Battle.Destruction>(messageData);

                    if (OnDestructionReceive != null)
                    {
                        OnDestructionReceive(message);
                    }

                    return message;
                }
                case MessageType.AccelerationRequest:
                {
                    var message = session.MessageDeserializer.Deserialize<TyphenApi.Type.Submarine.Battle.AccelerationRequestObject>(messageData);

                    if (OnAccelerationRequestReceive != null)
                    {
                        OnAccelerationRequestReceive(message);
                    }

                    return message;
                }
                case MessageType.BrakeRequest:
                {
                    var message = session.MessageDeserializer.Deserialize<TyphenApi.Type.Submarine.Battle.BrakeRequestObject>(messageData);

                    if (OnBrakeRequestReceive != null)
                    {
                        OnBrakeRequestReceive(message);
                    }

                    return message;
                }
                case MessageType.PingerRequest:
                {
                    var message = session.MessageDeserializer.Deserialize<TyphenApi.Type.Submarine.Battle.PingerRequestObject>(messageData);

                    if (OnPingerRequestReceive != null)
                    {
                        OnPingerRequestReceive(message);
                    }

                    return message;
                }
                case MessageType.ActorRequest:
                {
                    var message = session.MessageDeserializer.Deserialize<TyphenApi.Type.Submarine.Battle.ActorRequestObject>(messageData);

                    if (OnActorRequestReceive != null)
                    {
                        OnActorRequestReceive(message);
                    }

                    return message;
                }
            }


            return null;
        }
    }
}
