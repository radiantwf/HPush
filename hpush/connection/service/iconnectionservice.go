package service

type IConnectionService interface {
	Init() (err error)
	Port() (port int)
	IsValid() (ret bool)
	StartServe() (err error)
}
