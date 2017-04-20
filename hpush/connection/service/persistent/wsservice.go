package persistent

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 5120
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WSService struct {
	submitCallback SubmitCallback
	Hub            *WSHub `inject:""`
	port           int
}

func (s *WSService) Port() (port int) {
	port = s.port
	return
}

func (s *WSService) Init() (err error) {
	s.port = 8000
	return
}

func (s *WSService) IsValid() (ret bool) {
	ret = true
	return
}

// serveWs handles websocket requests from the peer.
func (s *WSService) serveWs(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	c := &WSConnection{ws: ws, connectedTime: &now}
	s.Hub.register <- c
	c.readPump(*s.Hub)
}

func (s *WSService) StartServe() (err error) {
	go s.Hub.run(s)
	http.HandleFunc("/ws", s.serveWs)
	addr := ":8000"
	err = http.ListenAndServe(addr, nil)
	return
}
