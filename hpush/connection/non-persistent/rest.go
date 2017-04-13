package nonpersistent

type RestConnection struct {
}

func (c *RestConnection) Port() (port int) {
	return 0
}

func (c *RestConnection) IsValid() (ret bool) {
	return false
}

func (c *RestConnection) StartServe() (err error) {
	return
}
