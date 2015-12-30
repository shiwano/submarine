package typhenapi

// Type is the interface of the TyphenAPI type.
type Type interface {
	Coerce() error
	Bytes(serializer *Serializer) ([]byte, error)
}
