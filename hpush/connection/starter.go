package connection

import (
	nonpersistent "HPush/hpush/connection/non-persistent"
	"HPush/hpush/connection/persistent"
	"time"
)

type IConnectionService interface {
	Init() (err error)
	Port() (port int)
	IsValid() (ret bool)
	StartServe() (err error)
}

type ConnectionStarter struct {
	services []IConnectionService
	Ws       *persistent.WSService      `inject:""`
	Rest     *nonpersistent.RestService `inject:""`
}

func (start *ConnectionStarter) Init() (err error) {
	start.services = []IConnectionService{
		start.Ws,
		start.Rest,
	}
	return
}
func (start *ConnectionStarter) Run() (err error) {
	for _, service := range start.services {
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
