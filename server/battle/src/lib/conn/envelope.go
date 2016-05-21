package conn

type envelope struct {
	messageType int
	data        []byte
}
