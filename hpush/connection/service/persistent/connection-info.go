package persistent

import (
	"HPush/hpush/connection/user"
	"time"
)

type ConnectionInfo struct {
	guid string
	user.UserInfo
	FirstConnectedTime time.Time
	LastConnectedTime  time.Time
	connection         IConnection
}
