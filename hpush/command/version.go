package command

import (
	"fmt"
	"runtime"
)

var cmdVersion = &Command{
	Run:       runVersion,
	UsageLine: "version",
	Short:     "print HPush version",
	Long:      `Version prints the HPush version`,
}

func runVersion(cmd *Command, args []string) bool {
	if len(args) != 0 {
		cmd.Usage()
	}

	fmt.Printf("version %s %s %s\n", "0.01", runtime.GOOS, runtime.GOARCH)
	return true
}
