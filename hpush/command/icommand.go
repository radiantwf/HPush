package command

type ICommand interface {
	Init()
	Usage()
	Name() string
	ShortUsage() string
	Description() string
	Example() string
	Run(args []string) (err error)
}
