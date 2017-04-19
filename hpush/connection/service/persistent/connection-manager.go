package persistent

import "sync"

type ConnectionManager struct {
	checkList1 map[string]map[string]ConnectionInfo
	checkList2 map[IConnection]ConnectionInfo
	mutex      sync.Mutex
}

func (m *ConnectionManager) Init() {
	m.mutex.Lock()
	m.checkList1 = make(map[string]map[string]ConnectionInfo)
	m.checkList2 = make(map[IConnection]ConnectionInfo)
	m.mutex.Unlock()
}

func (m *ConnectionManager) AppendConnection(ci ConnectionInfo) {
	m.mutex.Lock()
	userkey := ci.UserInfoKeyString()
	if _, exist := m.checkList2[userkey]; !exist {
		m.checkList1[userkey] = make(map[string]ConnectionInfo)
	}
	m.checkList1[userkey][ci.guid] = ci
	m.checkList2[ci.connection] = ci
	m.mutex.Unlock()
}

func (m *ConnectionManager) DeleteConnection(conn IConnection) {
	m.mutex.Lock()
	if ci, exist := m.checkList2[conn]; exist {
		userkey := ci.UserInfoKeyString()
		if _, exist := m.checkList1[userkey]; exist {
			delete(m.checkList1[userkey], ci.guid)
		}
		delete(m.checkList2, conn)
	}
	m.mutex.Unlock()
}
