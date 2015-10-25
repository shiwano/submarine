package typhenapi

// Coercer coerces the interface of the TyphenAPI type.
type Coercer interface {
	Coerce() error
}

// RealTimeMessage is kind of a TyphenAPI RealTimeMessage.
type RealTimeMessage interface {
	RealTimeMessageType() int32
	Coerce() error
}
