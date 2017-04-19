package connection

import (
	"HPush/hpush/connection/service"
	nonpersistent "HPush/hpush/connection/service/non-persistent"
	"HPush/hpush/connection/service/persistent"
	"time"
)

type ConnectionServer struct {
	services []service.IConnectionService
	Ws       *persistent.WSService      `inject:""`
	Rest     *nonpersistent.RestService `inject:""`
}

func (server *ConnectionServer) Init() (err error) {
	server.services = []service.IConnectionService{
		server.Ws,
		server.Rest,
	}
	return
}

func (server *ConnectionServer) Run() (err error) {
	for _, service := range server.services {
		service.Init()
		if service.IsValid() {
			go service.StartServe()
		}
	}
	for {
		select {
		case <-time.After(10 * time.Minute):
		}
	}
}
