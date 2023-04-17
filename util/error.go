package util

import (
	"fmt"
	"runtime/debug"

	"github.com/caivega/glog"
)

var Mode string

func Init(m string) {
	Mode = m
}

func IsTest() bool {
	return Mode == "test"
}

func IsDebug() bool {
	return Mode == "debug"
}

func ErrorOfNonexists(t string, target string) error {
	e := fmt.Errorf("can't find %s: %s", t, target)
	if IsDebug() {
		glog.Error(e)
		debug.PrintStack()
	}
	return e
}

func ErrorOfInvalid(t string, target string) error {
	e := fmt.Errorf("invalid %s: %s", t, target)
	if IsDebug() {
		glog.Error(e)
		debug.PrintStack()
	}
	return e
}

func ErrorOfUnknown(t string, target string) error {
	e := fmt.Errorf("unknown %s: %s", t, target)
	if IsDebug() {
		glog.Error(e)
		debug.PrintStack()
	}
	return e
}

func ErrorOf(reason string, t string, target string) error {
	e := fmt.Errorf("%s %s: %s", reason, t, target)
	if IsDebug() {
		glog.Error(e)
		debug.PrintStack()
	}
	return e
}
