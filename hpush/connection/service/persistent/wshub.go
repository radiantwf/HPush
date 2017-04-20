package persistent

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// WSHub maintains the set of active connections and broadcasts messages to the
// connections.
type WSHub struct {
	// Registered connections.
	Manager *ConnectionManager `inject:""`

	// Inbound messages.
	in chan *inbound

	// Outbound messages.
	out chan *outbound

	// Register requests from the connections.
	register chan *WSConnection

	// Unregister requests from connections.
	unregister chan *WSConnection
}

// Inbound messages from a WSConnection.
type inbound struct {
	conn    *WSConnection
	message []byte
	time    *time.Time
}

// Outbound messages
type outbound struct {
	conn    *WSConnection
	message []byte
}

// run 定义
func (h *WSHub) run(s *WSService) {
	for {
		select {
		case c := <-h.register:
			var conn IConnection
			conn = c
			ci := NewConnectionInfo(conn)
			h.Manager.AppendNewConnection(ci)
		case c := <-h.unregister:
			var conn IConnection
			conn = c
			h.Manager.DeleteConnection(conn)
		case in := <-h.in:
			if s.submitCallback != nil {
				var conn IConnection
				conn = in.conn
				if ci, err := h.Manager.GetCIByConn(conn); err != nil {
					go s.submitCallback(in.message, ci)
				}
			}
		case out := <-h.out:
			go h.write(out.message, out.conn)
		}
	}
}

// write 定义
func (h *WSHub) write(message []byte, c *WSConnection) {
	if err := c.write(websocket.TextMessage, message); err != nil {
		log.Printf("error: %v", err)
		return
	}
	now := time.Now()
	c.lastCommunicationTime = &now
	return
}

// // writePump pumps messages from the hub to the websocket WSConnection.
// func (c *WSConnection) writePump() {
// 	ticker := time.NewTicker(pingPeriod)
// 	defer func() {
// 		ticker.Stop()
// 		c.ws.Close()
// 	}()
// 	for {
// 		select {
// 		case message, ok := <-c.send:
// 			if !ok {
// 				c.write(websocket.CloseMessage, []byte{})
// 				return
// 			}
// 			if err := c.write(websocket.TextMessage, message); err != nil {
// 				return
// 			}
// 		case <-ticker.C:
// 			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
// 				return
// 			}
// 		}
// 	}
// }
