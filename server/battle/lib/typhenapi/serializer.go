package typhenapi

// Serializer is Serializer/Deserializer for TyphenAPI types.
type Serializer interface {
	Serialize(v interface{}) (data []byte, err error)
	Deserialize(b []byte, v interface{}) error
}
