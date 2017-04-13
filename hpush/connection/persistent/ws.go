package persistent

type WSConnection struct {
}

func (c *WSConnection) Port() (port int) {
	return 0
}

func (c *WSConnection) IsValid() (ret bool) {
	return false
}

func (c *WSConnection) StartServe() (err error) {
	return
}
