package main

import (
	"HPush/hpush/command"
	"HPush/hpush/common"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/facebookgo/inject"
	"github.com/golang/glog"
)

func main() {
	glog.MaxSize = 1024 * 1024 * 32
	rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())

	var cmd command.Commander
	config, err := common.NewConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := inject.Populate(&cmd, &config); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cmd.Init()
	cmd.Run()
}
