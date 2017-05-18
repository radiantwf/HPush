package persistent

import (
	"HPush/hpush/connection/user"
	"errors"
	"sync"
)

type ConnectionManager struct {
	checkList1 map[string]map[string]ConnectionInfo
	checkList2 map[IConnection]ConnectionInfo
	tmpList    map[IConnection]ConnectionInfo
	mutex      sync.Mutex
}

func (m *ConnectionManager) Init() {
	m.mutex.Lock()
	m.checkList1 = make(map[string]map[string]ConnectionInfo)
	m.checkList2 = make(map[IConnection]ConnectionInfo)
	m.tmpList = make(map[IConnection]ConnectionInfo)
	m.mutex.Unlock()
}

// func (m *ConnectionManager) AppendConnection(ci ConnectionInfo) (err error) {
// 	if ci.User == nil {
// 		err = errors.New("ConnectionInfo的用户信息不能为空")
// 		return
// 	}
// 	m.mutex.Lock()
// 	userkey := ci.User.UserInfoKeyString()
// 	if _, exist := m.checkList2[userkey]; !exist {
// 		m.checkList1[userkey] = make(map[string]ConnectionInfo)
// 	}
// 	m.checkList1[userkey][ci.GUID] = ci
// 	m.checkList2[ci.Connection] = ci
// 	m.mutex.Unlock()
// 	return
// }
func (m *ConnectionManager) AppendNewConnection(ci ConnectionInfo) (err error) {
	m.mutex.Lock()
	m.mutex.Unlock()
	return
}

func (m *ConnectionManager) RegistryUserInfo(ci ConnectionInfo) (err error) {
	m.mutex.Lock()
	m.mutex.Unlock()
	return
}

func (m *ConnectionManager) DeleteConnection(conn IConnection) (err error) {
	m.mutex.Lock()
	if ci, exist := m.checkList2[conn]; exist {
		userkey := ci.User.UserInfoKeyString()
		if l, exist := m.checkList1[userkey]; exist {
			delete(l, ci.GUID)
		}
		delete(m.checkList2, conn)
	}
	m.mutex.Unlock()
	return
}

func (m *ConnectionManager) GetConnectionsByUser(user user.UserInfo) (ciList []ConnectionInfo, err error) {
	userkey := user.UserInfoKeyString()
	if l, exist := m.checkList1[userkey]; exist {
		ciList = make([]ConnectionInfo, len(l))
		i := 0
		for _, v := range l {
			ciList[i] = v
			i++
		}
	}
	return
}

func (m *ConnectionManager) GetCIByConn(conn IConnection) (ci ConnectionInfo, err error) {
	if i, exist := m.checkList2[conn]; exist {
		ci = i
	} else {
		err = errors.New("无法找到这个链接")
	}
	return
}
