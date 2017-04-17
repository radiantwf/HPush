package command

import (
	"flag"
	"fmt"
	"os"
)

type Commander struct {
	Commands      []ICommand
	CmdHelp       *HelpCommand       `inject:""`
	CmdVersion    *VersionCommand    `inject:""`
	CmdConnection *ConnectionCommand `inject:""`
}

func (c *Commander) Init() {
	c.Commands = []ICommand{
		c.CmdHelp,
		c.CmdVersion,
		c.CmdConnection,
	}
	for _, cmd := range c.Commands {
		cmd.Init()
	}
}

func (c *Commander) Run() {
	flag.Parse()
	args := flag.Args()
	flag.Usage = c.Usage
	if len(args) < 1 {
		c.CmdHelp.Run([]string{})
		flag.Usage()
		return
	}
	argsLen := len(args)
	for index, arg := range args {
		for _, cmd := range c.Commands {
			if arg == cmd.Name() {
				cmd.Run(args[index+1:])
				if index == 0 && argsLen == 1 && arg == c.CmdHelp.Name() {
					flag.Usage()
				}
				return
			}
		}
	}
	fmt.Fprintf(os.Stderr, "hpush: unknown subcommand %q\nRun 'hpush help' for usage.\n", args[0])
}

func (c *Commander) Usage() {
	fmt.Fprintf(os.Stderr, "For Logging, use \"hpush [logging_options] [command]\". The logging options are:\n")
	flag.PrintDefaults()
}
