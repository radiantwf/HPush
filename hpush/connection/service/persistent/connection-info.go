package persistent

import (
	"HPush/hpush/common/guid"
	"HPush/hpush/connection/user"
)

type ConnectionInfo struct {
	Guid       string
	User       *user.UserInfo
	Connection IConnection
}

func NewConnectionInfo(conn IConnection) (ci ConnectionInfo) {
	ci = ConnectionInfo{Guid: (&guid.GUID{}).NewGUID(),
		Connection: conn,
	}
	return
}
