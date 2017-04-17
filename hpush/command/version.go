package command

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type VersionCommand struct {
	// Flag is a set of flags specific to this command.
	flag flag.FlagSet
}

var cmdVersion = &VersionCommand{}

func (c *VersionCommand) Init() {
}

// Name returns the command's name: the first word in the usage line.
func (c *VersionCommand) Name() string {
	return "version"
}

// ShortUsage is the short description shown in the 'go help' output.
func (c *VersionCommand) ShortUsage() string {
	return "print hpush version"
}

func (c *VersionCommand) Description() string {
	return `Version prints the HPush version`
}

func (c *VersionCommand) Example() string {
	return "hpush version"
}

func (c *VersionCommand) Usage() {
	fmt.Fprintf(os.Stderr, "Example: %s\n", c.Example())
	fmt.Fprintf(os.Stderr, "Description:\n")
	fmt.Fprintf(os.Stderr, "  %s\n", strings.TrimSpace(c.Description()))
}

func (c *VersionCommand) Run(args []string) (err error) {
	c.flag.Parse(args)
	newArgs := c.flag.Args()
	if len(newArgs) != 0 {
	}

	fmt.Printf("version %s %s %s\n", "0.01", runtime.GOOS, runtime.GOARCH)
	return
}
