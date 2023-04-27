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

func ErrorOfEmpty(t string, target string) error {
	e := fmt.Errorf("empty %s: %s", t, target)
	if IsDebug() {
		glog.Error(e)
		debug.PrintStack()
	}
	return e
}

func ErrorOfNotFound(t string, target string) error {
	e := fmt.Errorf("not found %s: %s", t, target)
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

func ErrorOfUnmatched(t, target string, expected interface{}, actual interface{}) error {
	e := fmt.Errorf("unmatched %s: %s, %v != %v", t, target, expected, actual)
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
