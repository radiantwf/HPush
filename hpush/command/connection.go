package command

import (
	"HPush/hpush/connection"
	"flag"
	"fmt"
	"os"
	"strings"
)

type ConnectionOptions struct {
	port *int
}

type ConnectionCommand struct {
	options ConnectionOptions
	// Flag is a set of flags specific to this command.
	flag   flag.FlagSet
	Server *connection.ConnectionServer `inject:""`
}

func (c *ConnectionCommand) Init() {
	c.options.port = c.flag.Int("port", 8080, "http listen port")
}

// Name returns the command's name: the first word in the usage line.
func (c *ConnectionCommand) Name() string {
	return "connection"
}

// ShortUsage is the short description shown in the 'go help' output.
func (c *ConnectionCommand) ShortUsage() string {
	return "print hpush connection"
}

func (c *ConnectionCommand) Description() string {
	return ""
}

func (c *ConnectionCommand) Example() string {
	return ""
}

func (c *ConnectionCommand) Usage() {
	fmt.Fprintf(os.Stderr, "Example: %s\n", c.Example())
	fmt.Fprintf(os.Stderr, "Default Parameters:\n")
	c.flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "Description:\n")
	fmt.Fprintf(os.Stderr, "  %s\n", strings.TrimSpace(c.Description()))
}

func (c *ConnectionCommand) Run(args []string) (err error) {
	c.flag.Parse(args)
	// newArgs := c.flag.Args()

	c.Server.Init()
	c.Server.Run()
	return
}
