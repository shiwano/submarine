package typhenapi

// Coercer coerces the interface of the TyphenAPI type.
type Coercer interface {
	Coerce() error
}

// MessageTypeHolder hold the type of the TyphenAPI RealTimeMessage.
type MessageTypeHolder interface {
	MessageType() int32
	Coerce() error
}
