package connection

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

// Connection wraps a web socket connection.
type Connection struct {
	base                   *websocket.Conn
	settings               *Settings
	Upgrader               *websocket.Upgrader
	OnBinaryMessageReceive func([]byte)
	OnTextMessageReceive   func(string)
	OnDisconnect           func()
	OnError                func(error)
	WriteBinaryMessage     chan []byte
	WriteTextMessage       chan string
	WriteCloseMessage      chan struct{}
}

// NewConnection creates a Connection.
func NewConnection(settings *Settings) *Connection {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  settings.ReadBufferSize,
		WriteBufferSize: settings.WriteBufferSize,
	}
	connection := &Connection{
		settings:           settings,
		Upgrader:           upgrader,
		WriteBinaryMessage: make(chan []byte, settings.MessageBufferSize),
		WriteTextMessage:   make(chan string, settings.MessageBufferSize),
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
	conn.base.SetWriteDeadline(time.Now().Add(conn.settings.WriteWait))
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

	ticker := time.NewTicker(conn.settings.PingPeriod)
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

	conn.base.SetReadLimit(conn.settings.MaxMessageSize)
	conn.base.SetReadDeadline(time.Now().Add(conn.settings.PongWait))
	conn.base.SetPongHandler(func(string) error {
		conn.base.SetReadDeadline(time.Now().Add(conn.settings.PongWait))
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

		switch messageType {
		case websocket.BinaryMessage:
			if conn.OnBinaryMessageReceive != nil {
				conn.OnBinaryMessageReceive(data)
			}
		case websocket.TextMessage:
			if conn.OnTextMessageReceive != nil {
				text := string(data)
				conn.OnTextMessageReceive(text)
			}
		}
	}

	if conn.OnDisconnect != nil {
		conn.OnDisconnect()
	}
}
