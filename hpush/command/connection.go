package command

import (
	"HPush/hpush/connection"
)

var cmdConnection = &Command{
	Run:       runConnection,
	UsageLine: "connection",
}

func runConnection(cmd *Command, args []string) bool {
	// if len(args) != 0 {
	// 	cmd.Usage()
	// }
	(&connection.ConnectionStarter{}).Run()
	return true
}
