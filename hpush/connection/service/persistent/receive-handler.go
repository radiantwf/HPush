package persistent

import (
	"HPush/hpush/connection/packet"
	"encoding/json"
)

type ReceiveMessageStruct struct {
	appid    *string `json:"appid"`
	username *string `json:"username"`
	group    *string `json:"group"`
}

type ReceiveHandler struct {
	Manager *ConnectionManager `inject:""`
}

func (h *ReceiveHandler) OnReceived(ci ConnectionInfo, payload []byte) (err error) {
	p, err1 := packet.NewPacket(payload)
	if err1 != nil {
		err = err1
		return
	}
	var m ReceiveMessageStruct
	err = json.Unmarshal(p.Data, &m)
	if err != nil {
		return
	}
	ci.AppendUserInfo(*m.appid, *m.username, *m.group)
	return
}
