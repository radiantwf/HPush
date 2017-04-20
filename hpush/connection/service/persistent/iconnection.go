package persistent

import (
	"time"
)

type IConnection interface {
	GetConnectedTime() (t *time.Time)
	GetLastCommunicationTime() (t *time.Time)
}

// SubmitCallback 定义
type SubmitCallback func(message []byte, ci ConnectionInfo)
