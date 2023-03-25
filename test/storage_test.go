package test

import (
	"fmt"
	"io"
	"os"
	"testing"

	. "github.com/tokentransfer/check"
	"github.com/tokentransfer/node/chunk"
	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/store"

	libstore "github.com/tokentransfer/interfaces/store"
)

type StorageSuite struct{}

func Test_Config(t *testing.T) {
	s := Suite(&StorageSuite{})
	TestingRun(t, s)
}

// func (suite *StorageSuite) testCode(c *C) {
// 	s, err := kv.NewKvStorage(2048)
// 	c.Assert(err, IsNil)
// 	test(c, s)
// }

func (suite *StorageSuite) TestStack(c *C) {
	db0 := &store.MemoryService{
		Name: "memory",
	}
	err := db0.Init(nil)
	c.Assert(err, IsNil)
	err = db0.Start()
	c.Assert(err, IsNil)

	db1 := &store.MemoryService{
		Name: "memory",
	}
	err = db1.Init(nil)
	c.Assert(err, IsNil)
	err = db1.Start()
	c.Assert(err, IsNil)

	stackdb, err := store.CreateStackService(db0)
	c.Assert(err, IsNil)
	err = stackdb.Push(db1)
	c.Assert(err, IsNil)

	s, err := chunk.NewStorageWith(stackdb, 2048*2048*8)
	c.Assert(err, IsNil)
	test(c, s)

	s0, err := chunk.LoadStorageWith(stackdb, 2048*2048*8)
	c.Assert(err, IsNil)
	dump(s0, "s0")

	err = stackdb.Commit()
	c.Assert(err, IsNil)

	remove(c, s)
	err = stackdb.Commit()
	c.Assert(err, IsNil)

	s1, err := chunk.LoadStorageWith(db0, 2048*2048*8)
	c.Assert(err, IsNil)
	dump(s1, "s1")

	s2, err := chunk.LoadStorageWith(db1, 2048*2048*8)
	c.Assert(err, IsNil)
	dump(s2, "s2")

	// dumpKv(stackdb, "stackdb")
	// dumpKv(db0, "db0")
	// dumpKv(db1, "db1")
}

func dumpKv(s libstore.KvService, name string) {
	fmt.Println("~~~~~~~~~~~~~~~~~~~" + name + "~~~~~~~~~~~~~~~~~~~")
	s.ListData(func(key []byte, value []byte) error {
		fmt.Println(key, value)
		return nil
	})
}

func remove(c *C, s core.Storage) {
	rootGroup, err := s.Group("/")
	c.Assert(err, IsNil)

	codeGroup, err := rootGroup.Group("code")
	c.Assert(err, IsNil)

	err = codeGroup.Remove("0xc287B1266732495Fe8c93CE3Ba631597153fdd91")
	c.Assert(err, IsNil)
	err = codeGroup.Commit()
	c.Assert(err, IsNil)
	dump(s, "remove code from dir")

	err = codeGroup.Remove("demo.wasm")
	c.Assert(err, IsNil)
	err = codeGroup.Commit()
	c.Assert(err, IsNil)
	dump(s, "remove file from dir")
}

func test(c *C, s core.Storage) {
	dump(s, "init")

	rootGroup, err := s.Group("/")
	c.Assert(err, IsNil)
	err = rootGroup.Commit()
	c.Assert(err, IsNil)
	dump(s, "group")

	codeGroup, err := rootGroup.AddGroup("code")
	c.Assert(err, IsNil)
	err = codeGroup.Commit()
	c.Assert(err, IsNil)
	dump(s, "code")

	t := s.Create("0xc287B1266732495Fe8c93CE3Ba631597153fdd91")
	_, err = t.Write([]byte(`
		#include <stdio.h>

		int main() {
			printf("hello world!\n");
			return 0;
		}
	`))
	c.Assert(err, IsNil)
	err = t.Close()
	c.Assert(err, IsNil)

	d := t.Data()
	ok, err := codeGroup.AddData("0xc287B1266732495Fe8c93CE3Ba631597153fdd91", d.Key())
	c.Assert(err, IsNil)
	c.Assert(ok, Equals, false)

	err = codeGroup.Commit()
	c.Assert(err, IsNil)
	d.Dispose()
	t.Dispose()

	dump(s, "data")

	size, err := s.Size(d.Key())
	c.Assert(err, IsNil)
	fmt.Println("0xc287B1266732495Fe8c93CE3Ba631597153fdd91", size)

	t = s.Create("demo.wasm")
	f, err := os.Open("../vm/testdata/demo.wasm")
	c.Assert(err, IsNil)
	_, err = io.Copy(t, f)
	c.Assert(err, IsNil)
	err = f.Close()
	c.Assert(err, IsNil)
	err = t.Close()
	c.Assert(err, IsNil)

	d = t.Data()
	ok, err = codeGroup.AddData("demo.wasm", d.Key())
	c.Assert(err, IsNil)
	c.Assert(ok, Equals, false)

	err = codeGroup.Commit()
	c.Assert(err, IsNil)
	d.Dispose()
	t.Dispose()

	dump(s, "file")

	size, err = s.Size(d.Key())
	c.Assert(err, IsNil)
	fmt.Println("demo.wasm", size)
}

func dump(s core.Storage, name string) {
	fmt.Println("=================" + name + "=================")
	s.DumpLog(logPrinter{})
}

type logPrinter struct {
}

func (p logPrinter) Printf(format string, v ...interface{}) {
	fmt.Printf(format+"\n", v...)
}