package main

import (
	"github.com/tokentransfer/node/conf"
	"github.com/tokentransfer/node/net"
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
}
