package connection

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const (
	writeDeadLine     = 10 * time.Second
	pongDeadLine      = 60 * time.Second
	pingPeriod        = (60 * time.Second * 9) / 10
	maxMessageSize    = 512
	messageBufferSize = 512
	readBufferSize    = 1024
	writeBufferSize   = 1024
)

// Connection wraps a web socket connection.
type Connection struct {
	base               *websocket.Conn
	Upgrader           *websocket.Upgrader
	OnMessageReceive   func([]byte)
	OnDisconnect       func()
	OnError            func(error)
	WriteBinaryMessage chan []byte
	WriteTextMessage   chan string
	WriteCloseMessage  chan struct{}
}

// NewConnection creates a Connection.
func NewConnection() *Connection {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  readBufferSize,
		WriteBufferSize: writeBufferSize,
	}
	connection := &Connection{
		Upgrader:           upgrader,
		WriteBinaryMessage: make(chan []byte, messageBufferSize),
		WriteTextMessage:   make(chan string, messageBufferSize),
		WriteCloseMessage:  make(chan struct{}),
	}
	connection.Upgrader = upgrader
	return connection
}

// Connect connects to a client.
func (conn *Connection) Connect(responseWriter http.ResponseWriter, request *http.Request) error {
	websocketConn, err := conn.Upgrader.Upgrade(responseWriter, request, nil)
	if err != nil {
		return err
	}

	conn.base = websocketConn
	go conn.writePump()
	go conn.readPump()
	return nil
}

func (conn *Connection) writeMessage(messageType int, data []byte) error {
	conn.base.SetWriteDeadline(time.Now().Add(writeDeadLine))
	return conn.base.WriteMessage(messageType, data)
}

func (conn *Connection) writeCloseMessage() error {
	return conn.writeMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}

func (conn *Connection) writePingMessage() error {
	return conn.writeMessage(websocket.PingMessage, []byte{})
}

func (conn *Connection) writeBinaryMessage(data []byte) error {
	return conn.writeMessage(websocket.BinaryMessage, data)
}

func (conn *Connection) writeTextMessage(text string) error {
	data := []byte(text)
	return conn.writeMessage(websocket.TextMessage, data)
}

func (conn *Connection) writePump() {
	defer conn.base.Close()

	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

loop:
	for {
		select {
		case data := <-conn.WriteBinaryMessage:
			if err := conn.writeBinaryMessage(data); err != nil {
				if conn.OnError != nil {
					conn.OnError(err)
				}
				break loop
			}
		case text := <-conn.WriteTextMessage:
			if err := conn.writeTextMessage(text); err != nil {
				if conn.OnError != nil {
					conn.OnError(err)
				}
				break loop
			}
		case <-conn.WriteCloseMessage:
			if err := conn.writeCloseMessage(); err != nil && conn.OnError != nil {
				conn.OnError(err)
			}
			break loop
		case <-ticker.C:
			if err := conn.writePingMessage(); err != nil {
				if conn.OnError != nil {
					conn.OnError(err)
				}
				break loop
			}
		}
	}
}

func (conn *Connection) readPump() {
	defer conn.base.Close()

	conn.base.SetReadLimit(maxMessageSize)
	conn.base.SetReadDeadline(time.Now().Add(pongDeadLine))
	conn.base.SetPongHandler(func(string) error {
		conn.base.SetReadDeadline(time.Now().Add(pongDeadLine))
		return nil
	})

	for {
		messageType, data, err := conn.base.ReadMessage()
		if err != nil {
			if conn.OnError != nil {
				conn.OnError(err)
			}
			break
		}

		if messageType == websocket.BinaryMessage && conn.OnMessageReceive != nil {
			conn.OnMessageReceive(data)
		}
	}

	if conn.OnDisconnect != nil {
		conn.OnDisconnect()
	}
}
