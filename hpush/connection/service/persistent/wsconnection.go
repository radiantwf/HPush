package persistent

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// WSConnection is an middleman between the websocket WSConnection and the hub.
type WSConnection struct {
	// The websocket WSConnection.
	ws *websocket.Conn
	// Connected time.
	connectedTime *time.Time
	// Last communication time.
	lastCommunicationTime *time.Time
	// Last ping time.
	lastPingTime *time.Time
}

func (c *WSConnection) GetConnectedTime() (t *time.Time) {
	t = c.connectedTime
	return
}
func (c *WSConnection) GetLastCommunicationTime() (t *time.Time) {
	t = c.lastCommunicationTime
	return
}

// readPump pumps messages from the websocket WSConnection to the hub.
func (c *WSConnection) readPump(h WSHub) {
	defer func() {
		h.unregister <- c
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		messagetype, message, err := c.ws.ReadMessage()
		if messagetype == websocket.CloseMessage {
			break
		}
		now := time.Now()
		c.lastCommunicationTime = &now
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		// if messagetype == websocket.PingMessage {
		// 	h.ping <- &ping{conn: c, pingdata: message, pingtime: &now}
		// } else
		if messagetype == websocket.TextMessage || messagetype == websocket.BinaryMessage {
			h.in <- &inbound{conn: c, message: message, time: &now}
		}
	}
}

// write writes a message with the given message type and payload.
func (c *WSConnection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}
