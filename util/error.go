package util

import (
	"fmt"
)

func ErrorOfNonexists(t string, target string) error {
	e := fmt.Errorf("can't find %s: %s", t, target)
	// glog.Error(e)
	// debug.PrintStack()
	return e
}

func ErrorOfInvalid(t string, target string) error {
	e := fmt.Errorf("invalid %s: %s", t, target)
	// glog.Error(e)
	// debug.PrintStack()
	return e
}

func ErrorOfUnknown(t string, target string) error {
	e := fmt.Errorf("unknown %s: %s", t, target)
	// glog.Error(e)
	// debug.PrintStack()
	return e
}

func ErrorOf(reason string, t string, target string) error {
	e := fmt.Errorf("%s %s: %s", reason, t, target)
	// glog.Error(e)
	// debug.PrintStack()
	return e
}
