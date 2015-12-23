package typhenapi

// Session for Web Socket API.
type Session interface {
	Send(msg []byte)
}
