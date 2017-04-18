package main

import (
	"HPush/hpush/command"
	"HPush/hpush/common"
	"HPush/hpush/common/data/crypto"
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
	key, err := crypto.GenRsaPrivateKey()
	fmt.Println(key)
	pubkey, err := crypto.GenRsaPublicKey(key)
	fmt.Println(pubkey)
	buf, _ := crypto.RsaCipher(pubkey, []byte("123123123123"))
	fmt.Println(string(buf))
	buf2, _ := crypto.RsaDecipher(key, buf)
	fmt.Println(string(buf2))

	cmd.Init()
	cmd.Run()
}
