/*
Copyright (C) BABEC. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package net

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	. "github.com/tokentransfer/check"
	"github.com/tokentransfer/node/conf"
)

const (
	chainId1 = "chain1"
	chainId2 = "chain2"
	msgFlag  = "TEST_PUSH"
)

type ConfigSuite struct{}

func Test_Config(t *testing.T) {
	s := Suite(&ConfigSuite{})
	TestingRun(t, s)
}

// TestNet 网络侧基本功能测试
func (suite *ConfigSuite) TestNet(c *C) {
	certPath := filepath.Join("./testdata/cert")
	pidPath := filepath.Join("./testdata/pid")
	defer func() {
		_ = filepath.Walk(filepath.Join("./"), func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() && strings.Contains(path, "default.log") {
				_ = os.Remove(path)
			}
			return nil
		})
	}()
	caFile6666 := filepath.Join(certPath, "ca1.crt")
	caFile7777 := filepath.Join(certPath, "ca2.crt")
	key1Path := filepath.Join(certPath, "key1.key")
	cert1Path := filepath.Join(certPath, "cert1.crt")
	pid1Bytes, err := ioutil.ReadFile(filepath.Join(pidPath, "pid1.nodeid"))
	c.Assert(err, IsNil)
	pid1 := string(pid1Bytes)
	key2Path := filepath.Join(certPath, "key2.key")
	cert2Path := filepath.Join(certPath, "cert2.crt")
	pid2Bytes, err := ioutil.ReadFile(filepath.Join(pidPath, "pid2.nodeid"))
	c.Assert(err, IsNil)
	pid2 := string(pid2Bytes)

	readyC := make(chan struct{})

	// start node A
	a, err := InitNet(&conf.Config{
		NetConfig: conf.NetConfig{
			Provider:   "libp2p", //libp2p,liquid
			ListenAddr: "/ip4/127.0.0.1/tcp/6666",
			TLSConfig: conf.TLSConfig{
				PrivKeyFile: key1Path,
				CertFile:    cert1Path,
			},
			CustomChainTrustRoots: []conf.ChainTrustRoots{
				{
					ChainId: chainId1,
					TrustRoots: []conf.TrustRoots{
						{
							Root: caFile6666,
						},
						{
							Root: caFile7777,
						},
					},
				},
				{
					ChainId: chainId2,
					TrustRoots: []conf.TrustRoots{
						{
							Root: caFile6666,
						},
						{
							Root: caFile7777,
						},
					},
				},
			},
		},
	}, readyC)
	c.Assert(err, IsNil)
	//a.AddSeed("/ip4/127.0.0.1/tcp/7777/p2p/QmeyNRs2DwWjcHTpcVHoUSaDAAif4VQZ2wQDQAUNDP33gH")
	err = a.Start()
	c.Assert(err, IsNil)
	err = a.InitPubSub(chainId1, 0)
	c.Assert(err, IsNil)
	err = a.InitPubSub(chainId2, 0)
	c.Assert(err, IsNil)
	fmt.Println("node A is running...")

	// start node B
	b, err := InitNet(&conf.Config{
		NetConfig: conf.NetConfig{
			Provider:   "libp2p", //libp2p,liquid
			ListenAddr: "/ip4/127.0.0.1/tcp/7777",
			TLSConfig: conf.TLSConfig{
				PrivKeyFile: key2Path,
				CertFile:    cert2Path,
			},
			CustomChainTrustRoots: []conf.ChainTrustRoots{
				{
					ChainId: chainId1,
					TrustRoots: []conf.TrustRoots{
						{
							Root: caFile6666,
						},
						{
							Root: caFile7777,
						},
					},
				},
				{
					ChainId: chainId2,
					TrustRoots: []conf.TrustRoots{
						{
							Root: caFile6666,
						},
						{
							Root: caFile7777,
						},
					},
				},
			},
		},
	}, readyC)
	c.Assert(err, IsNil)
	err = b.AddSeed("/ip4/127.0.0.1/tcp/6666/p2p/" + pid1)
	c.Assert(err, IsNil)
	err = b.Start()
	c.Assert(err, IsNil)
	err = b.InitPubSub(chainId1, 0)
	c.Assert(err, IsNil)
	err = b.InitPubSub(chainId2, 0)
	c.Assert(err, IsNil)
	fmt.Println("node B is running...")

	close(readyC)

	// test A send msg to B
	data := []byte("hello")
	toNodeB := pid2

	passChan := make(chan bool)

	recHandlerB := func(id string, msg []byte) error {
		fmt.Println("[B][chain1] recv a msg from peer[", id, "], msg：", string(msg))
		passChan <- true
		return nil
	}
	err = b.DirectMsgHandle(chainId1, msgFlag, recHandlerB)
	c.Assert(err, IsNil)
	fmt.Println("[B]B register receive msg handler for chain1")
	go func() {
		fmt.Println("[A]A send msg to B in chain1")
		for {
			if err = a.SendMsg(chainId1, toNodeB, msgFlag, data); err != nil {
				fmt.Println(err)
				time.Sleep(time.Second)
				continue
			}
			break
		}
	}()
	timer := time.NewTimer(time.Minute)
	select {
	case <-timer.C:
		fmt.Println("==== test A send msg to B timeout ====")
		c.Fatal("test A send msg to B timeout")
	case <-passChan:
		fmt.Println("==== test A send msg to B pass ====")
	}

	// test broadcast
	//testTopic := "testTopic"
	//subHandlerB := func(_ string, msg []byte) error {
	//	fmt.Println("[B]recv a sub msg of chain1：", string(msg))
	//	passChan <- true
	//	return nil
	//}
	//err = b.SubscribeWithChainId(chainId1, testTopic, subHandlerB)
	//require.Nil(t, err)
	//fmt.Println("[B]B subscribe topic of chain1")
	//err = a.BroadcastWithChainId(chainId1, testTopic, data)
	//require.Nil(t, err)
	//fmt.Println("[A]A broadcast a msg to chain1:", string(data))
	//
	//timer = time.NewTimer(time.Minute)
	//select {
	//case <-timer.C:
	//	fmt.Println("==== test broadcast timeout ====")
	//	t.Fatal("test broadcast failed")
	//case <-passChan:
	//	fmt.Println("==== test broadcast pass ====")
	//}
	//
	//// test cancel broadcast
	//err = b.CancelSubscribeWithChainId(chainId1, testTopic)
	//require.Nil(t, err)
	//fmt.Println("[B]B cancel subscribe topic of chain1")
	//
	//err = a.BroadcastWithChainId(chainId1, testTopic, data)
	//require.Nil(t, err)
	//fmt.Println("[A]A broadcast a msg to chain1:", string(data))
	//
	//timer = time.NewTimer(10 * time.Second)
	//select {
	//case <-timer.C:
	//	fmt.Println("==== test cancel broadcast pass ====")
	//case <-passChan:
	//	fmt.Println("==== test cancel broadcast failed ====")
	//	t.Fatal("test cancel broadcast failed")
	//}

	err = a.Stop()
	c.Assert(err, IsNil)
	err = b.Stop()
	c.Assert(err, IsNil)
}
