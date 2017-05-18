package persistent

import (
	"HPush/hpush/common/guid"
	"HPush/hpush/connection/user"
)

type ConnectionInfo struct {
	GUID       string
	User       *user.UserInfo
	Connection IConnection
}

func NewConnectionInfo(conn IConnection) (ci ConnectionInfo) {
	ci = ConnectionInfo{GUID: (&guid.GUID{}).NewGUID(),
		Connection: conn,
	}
	return
}

func (c *ConnectionInfo) AppendUserInfo(appid, username, group string) {
	c.User = user.New(appid, username, group)
}
