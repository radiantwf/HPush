package nonpersistent

type RestService struct {
}

func (s *RestService) Init() (err error) {
	return
}

func (s *RestService) Port() (port int) {
	return 0
}

func (s *RestService) IsValid() (ret bool) {
	return false
}

func (s *RestService) StartServe() (err error) {
	return
}
