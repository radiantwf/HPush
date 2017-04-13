package connection

import (
	nonpersistent "HPush/hpush/connection/non-persistent"
	"HPush/hpush/connection/persistent"
)

type IConnection interface {
	Port() (port int)
	IsValid() (ret bool)
	StartServe() (err error)
}

type ConnectionStarter struct {
	ws   *persistent.WSConnection      `inject:""`
	rest *nonpersistent.RestConnection `inject:""`
}

func (start *ConnectionStarter) Run() {
}
