package connection

type envelope struct {
	messageType int
	data        []byte
}
