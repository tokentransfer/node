package core

// Original from https://github.com/indyjo/cafs

import (
	"io"
)

type Storage interface {
	Root() Key
	Update() error
	Create(string) Temporary
	Group(string) (Group, error)
	Exists(Key) bool
	Get(Key) (Data, error)
	Chunk(Key) ([]byte, error)
	Size(Key) (int64, error)
	Reference(Key) (int64, error)
	Commit(Storage, Key) error
	Remove(Key) error
	Close() error
	DumpStatistics(Printer)
	DumpLog(Printer)

	Version() int64
	GetVersion(int64) (Key, error)

	GetUsed() int64
	GetCapacity() int64
}

type Group interface {
	Name() string
	Key() Key
	Parent() Group
	Group(string) (Group, error)
	AddData(string, Key) (bool, error)
	AddGroup(string) (Group, error)
	Remove(string) error
	Exists(string) bool
	List() []string
	GetKey(string) (Key, error)
	IsGroup(string) (bool, error)
	Commit() error
	Dispose()
	Dump(Printer)
}

type Data interface {
	Dispose()
	Remove()
	Key() Key
	Open() io.ReadCloser
	Size() int64
	Duplicate() Data
	IsChunked() bool
	Chunks() DataIterator
	NumChunks() int64
}

type DataIterator interface {
	Dispose() error
	Duplicate() (DataIterator, error)
	Next() (bool, error)
	Key() (Key, error)
	Size() (int64, error)
	Data() (Data, error)
}

type Temporary interface {
	io.WriteCloser
	Data() (Data, error)
	Dispose()
}

type Printer interface {
	Printf(format string, v ...interface{})
}
