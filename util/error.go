package util

import "fmt"

func ErrorOfNonexists(t string, target string) error {
	return fmt.Errorf("can't find %s: %s", t, target)
}

func ErrorOfInvalid(t string, target string) error {
	return fmt.Errorf("invalid %s: %s", t, target)
}

func ErrorOfUnknown(t string, target string) error {
	return fmt.Errorf("unknown %s: %s", t, target)
}

func ErrorOf(reason string, t string, target string) error {
	return fmt.Errorf("%s %s: %s", reason, t, target)
}
