// This file was generated by typhen-api

package battle

import (
	"app/typhenapi/core"
	submarine "app/typhenapi/type/submarine"
	submarine_battle "app/typhenapi/type/submarine/battle"
	"fmt"
)

const (
	MessageType_Ping                int32 = -1287020902
	MessageType_Room                int32 = -1286955548
	MessageType_Now                 int32 = -1565539188
	MessageType_Start               int32 = -1240750730
	MessageType_Finish              int32 = -162791524
	MessageType_Actor               int32 = -1257891252
	MessageType_Movement            int32 = 1298310360
	MessageType_Destruction         int32 = -1118469016
	MessageType_StartRequest        int32 = 504335322
	MessageType_AccelerationRequest int32 = -710337400
	MessageType_BrakeRequest        int32 = 1492486768
	MessageType_TurnRequest         int32 = 698416554
	MessageType_PingerRequest       int32 = 110864488
	MessageType_TorpedoRequest      int32 = 1327463172
)

// WebSocketAPI sends messages, and dispatches message events.
type WebSocketAPI struct {
	session      typhenapi.Session
	serializer   typhenapi.Serializer
	errorHandler func(error)

	PingHandler                func(message *submarine_battle.PingObject)
	RoomHandler                func(message *submarine.Room)
	NowHandler                 func(message *submarine_battle.NowObject)
	StartHandler               func(message *submarine_battle.Start)
	FinishHandler              func(message *submarine_battle.Finish)
	ActorHandler               func(message *submarine_battle.Actor)
	MovementHandler            func(message *submarine_battle.Movement)
	DestructionHandler         func(message *submarine_battle.Destruction)
	StartRequestHandler        func(message *submarine_battle.StartRequestObject)
	AccelerationRequestHandler func(message *submarine_battle.AccelerationRequestObject)
	BrakeRequestHandler        func(message *submarine_battle.BrakeRequestObject)
	TurnRequestHandler         func(message *submarine_battle.TurnRequestObject)
	PingerRequestHandler       func(message *submarine_battle.PingerRequestObject)
	TorpedoRequestHandler      func(message *submarine_battle.TorpedoRequestObject)
}

// New creates a WebSocketAPI.
func New(session typhenapi.Session, serializer typhenapi.Serializer, errorHandler func(error)) *WebSocketAPI {
	api := &WebSocketAPI{}
	api.session = session
	api.serializer = serializer
	api.errorHandler = errorHandler
	return api
}

// Send sends a message.
func (api *WebSocketAPI) Send(body typhenapi.Type) (message *typhenapi.Message, err error) {
	switch messageBody := body.(type) {
	case *submarine_battle.PingObject:
		message, err = typhenapi.NewMessage(api.serializer, MessageType_Ping, messageBody)
	case *submarine.Room:
		message, err = typhenapi.NewMessage(api.serializer, MessageType_Room, messageBody)
	case *submarine_battle.NowObject:
		message, err = typhenapi.NewMessage(api.serializer, MessageType_Now, messageBody)
	case *submarine_battle.Start:
		message, err = typhenapi.NewMessage(api.serializer, MessageType_Start, messageBody)
	case *submarine_battle.Finish:
		message, err = typhenapi.NewMessage(api.serializer, MessageType_Finish, messageBody)
	case *submarine_battle.Actor:
		message, err = typhenapi.NewMessage(api.serializer, MessageType_Actor, messageBody)
	case *submarine_battle.Movement:
		message, err = typhenapi.NewMessage(api.serializer, MessageType_Movement, messageBody)
	case *submarine_battle.Destruction:
		message, err = typhenapi.NewMessage(api.serializer, MessageType_Destruction, messageBody)
	case *submarine_battle.StartRequestObject:
		message, err = typhenapi.NewMessage(api.serializer, MessageType_StartRequest, messageBody)
	case *submarine_battle.AccelerationRequestObject:
		message, err = typhenapi.NewMessage(api.serializer, MessageType_AccelerationRequest, messageBody)
	case *submarine_battle.BrakeRequestObject:
		message, err = typhenapi.NewMessage(api.serializer, MessageType_BrakeRequest, messageBody)
	case *submarine_battle.TurnRequestObject:
		message, err = typhenapi.NewMessage(api.serializer, MessageType_TurnRequest, messageBody)
	case *submarine_battle.PingerRequestObject:
		message, err = typhenapi.NewMessage(api.serializer, MessageType_PingerRequest, messageBody)
	case *submarine_battle.TorpedoRequestObject:
		message, err = typhenapi.NewMessage(api.serializer, MessageType_TorpedoRequest, messageBody)
	default:
		err = fmt.Errorf("Unsupported TyphenAPI type is given: %v", messageBody)
	}

	if err == nil {
		err = api.session.Send(message.Bytes())
	}

	if err != nil && api.errorHandler != nil {
		api.errorHandler(err)
	}
	return
}

// SendPing sends a ping message.
func (api *WebSocketAPI) SendPing(ping *submarine_battle.PingObject) error {
	message, err := typhenapi.NewMessage(api.serializer, MessageType_Ping, ping)

	if err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}

	if err := api.session.Send(message.Bytes()); err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}
	return nil
}

// SendRoom sends a room message.
func (api *WebSocketAPI) SendRoom(room *submarine.Room) error {
	message, err := typhenapi.NewMessage(api.serializer, MessageType_Room, room)

	if err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}

	if err := api.session.Send(message.Bytes()); err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}
	return nil
}

// SendNow sends a now message.
func (api *WebSocketAPI) SendNow(now *submarine_battle.NowObject) error {
	message, err := typhenapi.NewMessage(api.serializer, MessageType_Now, now)

	if err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}

	if err := api.session.Send(message.Bytes()); err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}
	return nil
}

// SendStart sends a start message.
func (api *WebSocketAPI) SendStart(start *submarine_battle.Start) error {
	message, err := typhenapi.NewMessage(api.serializer, MessageType_Start, start)

	if err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}

	if err := api.session.Send(message.Bytes()); err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}
	return nil
}

// SendFinish sends a finish message.
func (api *WebSocketAPI) SendFinish(finish *submarine_battle.Finish) error {
	message, err := typhenapi.NewMessage(api.serializer, MessageType_Finish, finish)

	if err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}

	if err := api.session.Send(message.Bytes()); err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}
	return nil
}

// SendActor sends a actor message.
func (api *WebSocketAPI) SendActor(actor *submarine_battle.Actor) error {
	message, err := typhenapi.NewMessage(api.serializer, MessageType_Actor, actor)

	if err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}

	if err := api.session.Send(message.Bytes()); err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}
	return nil
}

// SendMovement sends a movement message.
func (api *WebSocketAPI) SendMovement(movement *submarine_battle.Movement) error {
	message, err := typhenapi.NewMessage(api.serializer, MessageType_Movement, movement)

	if err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}

	if err := api.session.Send(message.Bytes()); err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}
	return nil
}

// SendDestruction sends a destruction message.
func (api *WebSocketAPI) SendDestruction(destruction *submarine_battle.Destruction) error {
	message, err := typhenapi.NewMessage(api.serializer, MessageType_Destruction, destruction)

	if err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}

	if err := api.session.Send(message.Bytes()); err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}
	return nil
}

// SendStartRequest sends a startRequest message.
func (api *WebSocketAPI) SendStartRequest(startRequest *submarine_battle.StartRequestObject) error {
	message, err := typhenapi.NewMessage(api.serializer, MessageType_StartRequest, startRequest)

	if err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}

	if err := api.session.Send(message.Bytes()); err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}
	return nil
}

// SendAccelerationRequest sends a accelerationRequest message.
func (api *WebSocketAPI) SendAccelerationRequest(accelerationRequest *submarine_battle.AccelerationRequestObject) error {
	message, err := typhenapi.NewMessage(api.serializer, MessageType_AccelerationRequest, accelerationRequest)

	if err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}

	if err := api.session.Send(message.Bytes()); err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}
	return nil
}

// SendBrakeRequest sends a brakeRequest message.
func (api *WebSocketAPI) SendBrakeRequest(brakeRequest *submarine_battle.BrakeRequestObject) error {
	message, err := typhenapi.NewMessage(api.serializer, MessageType_BrakeRequest, brakeRequest)

	if err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}

	if err := api.session.Send(message.Bytes()); err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}
	return nil
}

// SendTurnRequest sends a turnRequest message.
func (api *WebSocketAPI) SendTurnRequest(turnRequest *submarine_battle.TurnRequestObject) error {
	message, err := typhenapi.NewMessage(api.serializer, MessageType_TurnRequest, turnRequest)

	if err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}

	if err := api.session.Send(message.Bytes()); err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}
	return nil
}

// SendPingerRequest sends a pingerRequest message.
func (api *WebSocketAPI) SendPingerRequest(pingerRequest *submarine_battle.PingerRequestObject) error {
	message, err := typhenapi.NewMessage(api.serializer, MessageType_PingerRequest, pingerRequest)

	if err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}

	if err := api.session.Send(message.Bytes()); err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}
	return nil
}

// SendTorpedoRequest sends a torpedoRequest message.
func (api *WebSocketAPI) SendTorpedoRequest(torpedoRequest *submarine_battle.TorpedoRequestObject) error {
	message, err := typhenapi.NewMessage(api.serializer, MessageType_TorpedoRequest, torpedoRequest)

	if err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}

	if err := api.session.Send(message.Bytes()); err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}
	return nil
}

// DispatchMessageEvent dispatches a binary message.
func (api *WebSocketAPI) DispatchMessageEvent(data []byte) error {

	message, err := typhenapi.NewMessageFromBytes(data)

	if err != nil {
		if api.errorHandler != nil {
			api.errorHandler(err)
		}
		return err
	}

	switch message.Type {
	case MessageType_Ping:
		typhenType := new(submarine_battle.PingObject)
		if err := api.serializer.Deserialize(message.Body, typhenType); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if err := typhenType.Coerce(); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if api.PingHandler != nil {
			api.PingHandler(typhenType)
		}
	case MessageType_Room:
		typhenType := new(submarine.Room)
		if err := api.serializer.Deserialize(message.Body, typhenType); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if err := typhenType.Coerce(); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if api.RoomHandler != nil {
			api.RoomHandler(typhenType)
		}
	case MessageType_Now:
		typhenType := new(submarine_battle.NowObject)
		if err := api.serializer.Deserialize(message.Body, typhenType); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if err := typhenType.Coerce(); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if api.NowHandler != nil {
			api.NowHandler(typhenType)
		}
	case MessageType_Start:
		typhenType := new(submarine_battle.Start)
		if err := api.serializer.Deserialize(message.Body, typhenType); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if err := typhenType.Coerce(); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if api.StartHandler != nil {
			api.StartHandler(typhenType)
		}
	case MessageType_Finish:
		typhenType := new(submarine_battle.Finish)
		if err := api.serializer.Deserialize(message.Body, typhenType); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if err := typhenType.Coerce(); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if api.FinishHandler != nil {
			api.FinishHandler(typhenType)
		}
	case MessageType_Actor:
		typhenType := new(submarine_battle.Actor)
		if err := api.serializer.Deserialize(message.Body, typhenType); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if err := typhenType.Coerce(); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if api.ActorHandler != nil {
			api.ActorHandler(typhenType)
		}
	case MessageType_Movement:
		typhenType := new(submarine_battle.Movement)
		if err := api.serializer.Deserialize(message.Body, typhenType); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if err := typhenType.Coerce(); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if api.MovementHandler != nil {
			api.MovementHandler(typhenType)
		}
	case MessageType_Destruction:
		typhenType := new(submarine_battle.Destruction)
		if err := api.serializer.Deserialize(message.Body, typhenType); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if err := typhenType.Coerce(); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if api.DestructionHandler != nil {
			api.DestructionHandler(typhenType)
		}
	case MessageType_StartRequest:
		typhenType := new(submarine_battle.StartRequestObject)
		if err := api.serializer.Deserialize(message.Body, typhenType); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if err := typhenType.Coerce(); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if api.StartRequestHandler != nil {
			api.StartRequestHandler(typhenType)
		}
	case MessageType_AccelerationRequest:
		typhenType := new(submarine_battle.AccelerationRequestObject)
		if err := api.serializer.Deserialize(message.Body, typhenType); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if err := typhenType.Coerce(); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if api.AccelerationRequestHandler != nil {
			api.AccelerationRequestHandler(typhenType)
		}
	case MessageType_BrakeRequest:
		typhenType := new(submarine_battle.BrakeRequestObject)
		if err := api.serializer.Deserialize(message.Body, typhenType); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if err := typhenType.Coerce(); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if api.BrakeRequestHandler != nil {
			api.BrakeRequestHandler(typhenType)
		}
	case MessageType_TurnRequest:
		typhenType := new(submarine_battle.TurnRequestObject)
		if err := api.serializer.Deserialize(message.Body, typhenType); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if err := typhenType.Coerce(); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if api.TurnRequestHandler != nil {
			api.TurnRequestHandler(typhenType)
		}
	case MessageType_PingerRequest:
		typhenType := new(submarine_battle.PingerRequestObject)
		if err := api.serializer.Deserialize(message.Body, typhenType); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if err := typhenType.Coerce(); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if api.PingerRequestHandler != nil {
			api.PingerRequestHandler(typhenType)
		}
	case MessageType_TorpedoRequest:
		typhenType := new(submarine_battle.TorpedoRequestObject)
		if err := api.serializer.Deserialize(message.Body, typhenType); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if err := typhenType.Coerce(); err != nil {
			if api.errorHandler != nil {
				api.errorHandler(err)
			}
			return err
		}

		if api.TorpedoRequestHandler != nil {
			api.TorpedoRequestHandler(typhenType)
		}
	}

	return nil
}
