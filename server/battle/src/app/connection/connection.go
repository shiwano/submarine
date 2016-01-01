package connection

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

// Connection wraps a web socket connection.
type Connection struct {
	conn                   *websocket.Conn
	Settings               *Settings
	Upgrader               *websocket.Upgrader
	OnBinaryMessageReceive func([]byte)
	OnTextMessageReceive   func(string)
	OnDisconnect           func()
	OnError                func(error)
	WriteBinaryMessage     chan []byte
	WriteTextMessage       chan string
	WriteCloseMessage      chan struct{}
}

// New creates a Connection.
func New() *Connection {
	connection := &Connection{
		Settings: NewSettings(),
		Upgrader: new(websocket.Upgrader),
	}
	return connection
}

// Connect connects to the client.
func (c *Connection) Connect(responseWriter http.ResponseWriter, request *http.Request) error {
	c.Upgrader.ReadBufferSize = c.Settings.ReadBufferSize
	c.Upgrader.WriteBufferSize = c.Settings.WriteBufferSize
	c.WriteBinaryMessage = make(chan []byte, c.Settings.MessageBufferSize)
	c.WriteTextMessage = make(chan string, c.Settings.MessageBufferSize)
	c.WriteCloseMessage = make(chan struct{})

	websocketConn, err := c.Upgrader.Upgrade(responseWriter, request, nil)
	if err != nil {
		return err
	}
	c.conn = websocketConn

	go c.writePump()
	go c.readPump()
	return nil
}

func (c *Connection) writeMessage(messageType int, data []byte) error {
	c.conn.SetWriteDeadline(time.Now().Add(c.Settings.WriteWait))
	return c.conn.WriteMessage(messageType, data)
}

func (c *Connection) writeCloseMessage() error {
	return c.writeMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}

func (c *Connection) writePingMessage() error {
	return c.writeMessage(websocket.PingMessage, []byte{})
}

func (c *Connection) writeBinaryMessage(data []byte) error {
	return c.writeMessage(websocket.BinaryMessage, data)
}

func (c *Connection) writeTextMessage(text string) error {
	data := []byte(text)
	return c.writeMessage(websocket.TextMessage, data)
}

func (c *Connection) writePump() {
	defer c.conn.Close()

	ticker := time.NewTicker(c.Settings.PingPeriod)
	defer ticker.Stop()

loop:
	for {
		select {
		case data := <-c.WriteBinaryMessage:
			if err := c.writeBinaryMessage(data); err != nil {
				if c.OnError != nil {
					c.OnError(err)
				}
				break loop
			}
		case text := <-c.WriteTextMessage:
			if err := c.writeTextMessage(text); err != nil {
				if c.OnError != nil {
					c.OnError(err)
				}
				break loop
			}
		case <-c.WriteCloseMessage:
			if err := c.writeCloseMessage(); err != nil && c.OnError != nil {
				c.OnError(err)
			}
			break loop
		case <-ticker.C:
			if err := c.writePingMessage(); err != nil {
				if c.OnError != nil {
					c.OnError(err)
				}
				break loop
			}
		}
	}
}

func (c *Connection) readPump() {
	defer c.conn.Close()

	c.conn.SetReadLimit(c.Settings.MaxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(c.Settings.PongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(c.Settings.PongWait))
		return nil
	})

	for {
		messageType, data, err := c.conn.ReadMessage()
		if err != nil {
			if c.OnError != nil {
				c.OnError(err)
			}
			break
		}

		switch messageType {
		case websocket.BinaryMessage:
			if c.OnBinaryMessageReceive != nil {
				c.OnBinaryMessageReceive(data)
			}
		case websocket.TextMessage:
			if c.OnTextMessageReceive != nil {
				text := string(data)
				c.OnTextMessageReceive(text)
			}
		}
	}

	if c.OnDisconnect != nil {
		c.OnDisconnect()
	}
}
