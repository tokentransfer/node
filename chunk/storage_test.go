package chunk

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"reflect"
	"strings"
	"testing"

	libcore "github.com/tokentransfer/interfaces/core"
	"github.com/tokentransfer/node/core"
)

func TestStorage(t *testing.T) {
	var key core.Key
	fmt.Println(key.String())
	if key.String() != "" {
		t.FailNow()
	}
}

func dump(s core.Storage, name string) {
	fmt.Println("=================" + name + "=================")
	s.DumpLog(logPrinter{})
}

func TestFile(t *testing.T) {
	N := 10 * 1024
	s, err := NewStorage(80 * 1024)
	if err != nil {
		t.Fatal(err)
	}
	// dump(s, "storage.create")
	temp := s.Create(fmt.Sprintf("%v random bytes", N))
	// dump(s, "temp.create")

	buf := make([]byte, N)
	for i, _ := range buf {
		buf[i] = byte(rand.Int())
	}
	if _, err := temp.Write(buf); err != nil {
		panic(err)
	}
	// dump(s, "temp.write")
	if err := temp.Close(); err != nil {
		panic(err)
	}
	// dump(s, "temp.close")
	f := temp.Data()
	// dump(s, "file.create")
	f.Dispose()
	// dump(s, "file.dispose")
}

func TestEntry(t *testing.T) {
	entry := &chunkEntry{
		name: "test",
		data: []byte("test"),
		chunks: []chunkRef{
			chunkRef{
				key:     core.MustParseKey("E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855"),
				nextPos: 0,
			},
		},
		refs: 1,
	}
	data, err := entry.MarshalBinary()
	if err != nil {
		t.FailNow()
	}

	var newEntry chunkEntry
	err = newEntry.UnmarshalBinary(data)
	if err != nil {
		t.FailNow()
	}

	if !reflect.DeepEqual(entry, &newEntry) {
		t.FailNow()
	}
}

func TestDirectory(t *testing.T) {
	_s, err := NewStorage(2000)
	if err != nil {
		t.Fatal(err)
	}
	s := _s.(*chunkStorage)
	f1 := addData(t, s, 128)
	f2 := addData(t, s, 256)

	dir, err := s.Group("/cache")
	if err != nil {
		t.Fatal(err)
	}
	_, err = dir.AddData("f1", f1.Key())
	if err != nil {
		t.Fatal(err)
	}
	_, err = dir.AddData("f2", f2.Key())
	if err != nil {
		t.Fatal(err)
	}
	list := dir.List()
	if len(list) != 2 {
		t.FailNow()
	}
	if list[0] != "f1" {
		t.FailNow()
	}
	if list[1] != "f2" {
		t.FailNow()
	}

	key := dir.Key()
	if key == nil {
		t.FailNow()
	}

	err = dir.Commit()
	if err != nil {
		t.Fatal(err)
	}
	// dir.Dump(logPrinter{})

	ndir, err := s.Group("/data")
	if err != nil {
		t.Fatal(err)
	}
	_, err = ndir.AddData("f1", f1.Key())
	if err != nil {
		t.Fatal(err)
	}
	_, err = ndir.AddData("f2", f2.Key())
	if err != nil {
		t.Fatal(err)
	}
	f3 := addData(t, s, 64)
	_, err = ndir.AddData("f3", f3.Key())
	if err != nil {
		t.Fatal(err)
	}
	err = ndir.Commit()
	if err != nil {
		t.Fatal(err)
	}
	// ndir.Dump(logPrinter{})
	// dump(s, "list")

	dir.Dispose()
	// dump(s, "after dir disposed")

	ndir.Dispose()
	// dump(s, "after ndir disposed")

	_f1 := f1.(*chunkData)
	entry := _f1.Entry()
	if entry.refs != 1 {
		t.Fatalf("Refs != 1: %v", entry.refs)
	}

	_f2 := f2.(*chunkData)
	entry = _f2.Entry()
	if entry.refs != 1 {
		t.Fatalf("Refs != 1: %v", entry.refs)
	}

	_f3 := f3.(*chunkData)
	entry = _f3.Entry()
	if entry.refs != 1 {
		t.Fatalf("Refs != 1: %v", entry.refs)
	}
}

func TestDirRecursive(t *testing.T) {
	_s, err := NewStorage(5000)
	if err != nil {
		t.Fatal(err)
	}
	s := _s.(*chunkStorage)
	f1 := addData(t, s, 128)
	f2 := addData(t, s, 256)

	group, err := s.Group("/")
	if err != nil {
		t.Fatal(err)
	}
	_, err = group.AddData("f1", f1.Key())
	if err != nil {
		t.Fatal(err)
	}
	_, err = group.AddData("f2", f2.Key())
	if err != nil {
		t.Fatal(err)
	}
	list := group.List()
	if len(list) != 2 {
		t.FailNow()
	}
	if list[0] != "f1" {
		t.FailNow()
	}
	if list[1] != "f2" {
		t.FailNow()
	}

	key := group.Key()
	if key == nil {
		t.FailNow()
	}

	cacheDir, err := group.AddGroup("cache")
	if err != nil {
		t.Fatal(err)
	}
	f4 := addData(t, s, 32)
	_, err = cacheDir.AddData("f4", f4.Key())
	if err != nil {
		t.Fatal(err)
	}
	err = cacheDir.Commit()
	if err != nil {
		t.Fatal(err)
	}
	// cacheDir.Dump(logPrinter{})

	dataGroup, err := group.AddGroup("data")
	if err != nil {
		t.Fatal(err)
	}
	f3 := addData(t, s, 64)
	_, err = dataGroup.AddData("f3", f3.Key())
	if err != nil {
		t.Fatal(err)
	}
	err = dataGroup.Commit()
	if err != nil {
		t.Fatal(err)
	}
	// dataDir.Dump(logPrinter{})

	err = group.Commit()
	if err != nil {
		t.Fatal(err)
	}
	// dir.Dump(logPrinter{})
	// dump(s, "list")

	err = group.Remove("data")
	if err != nil {
		t.Fatal(err)
	}
	err = group.Commit()
	if err != nil {
		t.Fatal(err)
	}
	// dump(s, "after remove data dir")
	// dir.Dump(logPrinter{})

	cacheDir.Dispose()
	// dump(s, "after dispose cache dir")
	// dir.Dump(logPrinter{})

	err = group.Remove("cache")
	if err != nil {
		t.Fatal(err)
	}
	err = group.Commit()
	if err != nil {
		t.Fatal(err)
	}
	// dump(s, "after remove data dir")
	// dir.Dump(logPrinter{})

	err = group.Remove("f1")
	if err != nil {
		t.Fatal(err)
	}
	err = group.Remove("f2")
	if err != nil {
		t.Fatal(err)
	}
	err = group.Commit()
	if err != nil {
		t.Fatal(err)
	}
	// dump(s, "after remove f1 and f2")
	// dir.Dump(logPrinter{})
}

func TestSimple(t *testing.T) {
	s, err := NewStorage(1000)
	if err != nil {
		t.Fatal(err)
	}
	_ = addData(t, s, 128)
}

func TestTwo(t *testing.T) {
	s, err := NewStorage(1000)
	if err != nil {
		t.Fatal(err)
	}
	f1 := addData(t, s, 128)
	f2 := addData(t, s, 256)
	if f1.Key().Equals(f2.Key()) {
		t.FailNow()
	}
}

func TestSame(t *testing.T) {
	s, err := NewStorage(1000)
	if err != nil {
		t.Fatal(err)
	}
	f1 := addData(t, s, 128)
	f2 := addData(t, s, 128)
	if !f1.Key().Equals(f2.Key()) {
		t.FailNow()
	}
}

func TestEmptyFile(t *testing.T) {
	s, err := NewStorage(1000)
	if err != nil {
		t.Fatal(err)
	}
	f := addData(t, s, 0)
	if f.Size() != 0 {
		t.FailNow()
	}
	iter := f.Chunks()
	if !iter.Next() {
		t.Fatal("Expected empty file to have at least one chunk")
	}
	if iter.Key().String() != "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855" {
		t.Fatalf("Unexpected key of empty chunk: %v", iter.Key())
	}
	if iter.Next() {
		t.Fatal("Expected empty file to not have any further chunks")
	}
}

type logPrinter struct {
}

func (p logPrinter) Printf(format string, v ...interface{}) {
	fmt.Printf(format+"\n", v...)
}

func TestLRU(t *testing.T) {
	s, err := NewStorage(1000)
	if err != nil {
		t.Fatal(err)
	}
	f1 := addData(t, s, 400)
	f1.Dispose()
	// s.DumpStatistics(logPrinter{})
	f2 := addData(t, s, 350)
	f2.Dispose()
	// s.DumpStatistics(logPrinter{})
	f3 := addData(t, s, 250)
	f3.Dispose()
	// s.DumpStatistics(logPrinter{})
	f4 := addData(t, s, 450)
	f4.Dispose()
	// s.DumpStatistics(logPrinter{})
	var key core.Key
	key = f1.Key()
	if _, err := s.Get(key); !strings.HasPrefix(err.Error(), "can't find") {
		t.Fatalf("f1 should have been removed. err:%v", err)
	}
	key = f2.Key()
	if _, err := s.Get(key); !strings.HasPrefix(err.Error(), "can't find") {
		t.Fatalf("f2 should have been removed. err:%v", err)
	}
	key = f4.Key()
	if _, err := s.Get(key); !strings.HasPrefix(err.Error(), "can't find") {
		t.Fatalf("f4 should have been removed. err:%v", err)
	}
	key = f3.Key()
	if _, err := s.Get(key); !strings.HasPrefix(err.Error(), "can't find") {
		t.Fatalf("f3 should not have been removed. err:%v", err)
	}

	// s.DumpStatistics(logPrinter{})

	// Now f3 is youngest, then f4 (f1 and f2 are gone)
	addData(t, s, 500).Dispose()

	key = f4.Key()
	if _, err := s.Get(key); !strings.HasPrefix(err.Error(), "can't find") {
		t.Fatalf("f4 should have been removed. err:%v", err)
	}
	key = f3.Key()
	if _, err := s.Get(key); !strings.HasPrefix(err.Error(), "can't find") {
		t.Fatalf("f3 should have been removed. err:%v", err)
	}

	{
		defer func() {
			if v := recover(); strings.HasPrefix(v.(error).Error(), "not enough space") {
				t.Logf("Expectedly recovered from: %v", v)
			} else {
				t.Fatalf("Expected to recover from something other than: %v", v)
			}
		}()
		addData(t, s, 1010)
	}
}
func TestCompression(t *testing.T) {
	s, err := NewStorage(8 * 1024 * 1024)
	if err != nil {
		t.Fatal(err)
	}
	f1 := addData(t, s, 512*1024)
	f1r := f1.Open()
	buf1 := new(bytes.Buffer)
	if _, err := io.Copy(buf1, f1r); err != nil {
		t.Fatal(err)
	}

	f2 := addData(t, s, 512*1024+1)
	f2r := f2.Open()
	buf2 := new(bytes.Buffer)
	if _, err := io.Copy(buf2, f2r); err != nil {
		t.Fatal(err)
	}

	h1 := cs.NewHasher()
	h1.Write(buf1.Bytes())
	if f1.Key().String() != libcore.EncodeToString(h1.Sum(nil)) {
		t.Fatal(err)
	}

	h2 := cs.NewHasher()
	h2.Write(buf2.Bytes())
	if f2.Key().String() != libcore.EncodeToString(h2.Sum(nil)) {
		t.Fatal(err)
	}

	// dump(s, "list")

	f1.Dispose()
	f1r.Close()

	f2.Dispose()
	f2r.Close()

	// dump(s, "after dispose")
}

func TestCompression2(t *testing.T) {
	s, err := NewStorage(1000000)
	if err != nil {
		t.Fatal(err)
	}
	temp := s.Create("Adding cyclic random data")
	defer temp.Dispose()
	cycle := 65536
	times := 24
	r := rand.New(rand.NewSource(0))
	data := make([]byte, cycle)
	for i := 0; i < cycle; i++ {
		data[i] = byte(r.Int())
	}
	t.Logf("data=%016x...", data[:8])
	for i := 0; i < times; i++ {
		if _, err := temp.Write(data); err != nil {
			t.Errorf("Error on Write: %v", err)
		}
	}
	if err := temp.Close(); err != nil {
		t.Errorf("Error on Close: %v", err)
	}

	f := temp.Data()
	defer f.Dispose()
	w := f.Open()
	data2 := make([]byte, 1)
	for i := 0; i < times*cycle; i++ {
		if n, err := io.ReadFull(w, data2); err != nil || n != 1 {
			t.Fatalf("Error on Read: %v (n=%d)", err, n)
		}
		if data2[0] != data[i%cycle] {
			t.Fatalf("Data read != data written on byte %d: %02x != %02x", i, data2[0], data[i%cycle])
		}
	}
}

func TestCommit(t *testing.T) {
	a, err := NewStorage(80 * 1024)
	if err != nil {
		t.Fatal(err)
	}
	f1 := addRandomData(t, a, 15*1024)
	f2 := addRandomData(t, a, 25*1024)

	b, err := NewStorage(80 * 1024)
	if err != nil {
		t.Fatal(err)
	}
	addRandomData(t, b, 10*1024)
	err = a.Commit(b, f1.Key())
	if err != nil {
		t.Fatal(err)
	}

	err = a.Commit(b, f2.Key())
	if err != nil {
		t.Fatal(err)
	}

	f1.Dispose()
	f2.Dispose()
}

func TestRemove(t *testing.T) {
	a, err := NewStorage(80 * 1024)
	if err != nil {
		t.Fatal(err)
	}
	f1 := addRandomData(t, a, 15*1024)
	f2 := addRandomData(t, a, 25*1024)
	// a.Dump(logPrinter{})

	b, err := NewStorage(80 * 1024)
	if err != nil {
		t.Fatal(err)
	}
	f3 := addRandomData(t, b, 10*1024)
	// b.Dump(logPrinter{})

	err = a.Remove(f1.Key())
	if err != nil {
		t.Fatal(err)
	}
	// a.Dump(logPrinter{})

	err = a.Remove(f2.Key())
	if err != nil {
		t.Fatal(err)
	}
	// a.Dump(logPrinter{})

	err = b.Remove(f3.Key())
	if err != nil {
		t.Fatal(err)
	}
	// b.Dump(logPrinter{})
}

func TestRefCounting(t *testing.T) {
	_s, err := NewStorage(80 * 1024)
	if err != nil {
		t.Fatal(err)
	}
	// s := _s.(*kvStorage)
	_f := addRandomData(t, _s, 60*1024)
	f := _f.(*chunkData)
	// dump(s, "after.add")
	entry := f.Entry()
	if entry.refs != 1 {
		t.Fatalf("Refs != 1 before dispose: %v", entry.refs)
	}
	_f.Dispose()
	// dump(s, "after.dispose")
	entry = f.Entry()
	if entry != nil {
		t.Fatalf("Removed after dispose: %v", entry.refs)
	}
	// This has to push out many chunks of first file
	addRandomData(t, _s, 60*1024)
}

func addData(t *testing.T, s core.Storage, size int) core.Data {
	temp := s.Create(fmt.Sprintf("Adding %v bytes object", size))
	defer temp.Dispose()
	for size > 0 {
		if _, err := temp.Write([]byte{byte(size)}); err != nil {
			panic(err)
		}
		size--
	}
	if err := temp.Close(); err != nil {
		panic(err)
	}
	return temp.Data()
}

func addRandomData(t *testing.T, s core.Storage, size int) core.Data {
	temp := s.Create(fmt.Sprintf("%v random bytes", size))
	defer temp.Dispose()
	buf := make([]byte, size)
	for i, _ := range buf {
		buf[i] = byte(rand.Int())
	}
	if _, err := temp.Write(buf); err != nil {
		panic(err)
	}
	if err := temp.Close(); err != nil {
		panic(err)
	}
	return temp.Data()
}
