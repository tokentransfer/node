package main

import (
	"fmt"
	"io"
	"os"

	"github.com/tokentransfer/node/conf"
	"github.com/tokentransfer/node/net"
	"github.com/tokentransfer/node/vm"
)

func main() {
	readyC := make(chan struct{})
	c, err := conf.NewConfig("./config.json")
	if err != nil {
		panic(err)
	}
	n, err := net.InitNet(c, readyC)
	if err != nil {
		panic(err)
	}
	err = n.Start()
	if err != nil {
		panic(err)
	}
	err = n.InitPubSub(c.GetChainId(), 0)
	if err != nil {
		panic(err)
	}
	close(readyC)

	err = n.Stop()
	if err != nil {
		panic(err)
	}

	f, err := os.Open("./lib.wasm")
	if err != nil {
		panic(err)
	}
	wasmBytes, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	results, err := vm.RunWasm(wasmBytes, "sum", 4.2, 3.1)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)
}
