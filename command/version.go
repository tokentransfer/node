package command

import (
	"flag"
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/mitchellh/cli"
)

var (
	Version = "dev"
	Commit  = "-"
	Date    = ""
)

// VersionCommand
type VersionCommand struct {
	Ui cli.Ui

	args []string
}

var _ cli.Command = &VersionCommand{}

func (i *VersionCommand) Help() string {
	helpText := `
Usage: node version [options]
Options:
`
	return strings.TrimSpace(helpText)
}

func (i *VersionCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("version", flag.ContinueOnError)
	cmdFlags.Usage = func() { i.Ui.Output(i.Help()) }
	if err := cmdFlags.Parse(i.args); err != nil {
		return 1
	}

	if info, available := debug.ReadBuildInfo(); available {
		if Date == "" {
			Version = info.Main.Version
			Commit = fmt.Sprintf("(unknown, mod sum: %q)", info.Main.Sum)
			Date = "(unknown)"
		}
	}

	fmt.Printf(` version     : %s
 commit      : %s
 build date  : %s
 go version  : %s
 go compiler : %s
 platform    : %s/%s
`, Version, Commit, Date, runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH)

	return 0
}

func (i *VersionCommand) Synopsis() string {
	return "show the version"
}

func (i *VersionCommand) Printf(format string, v ...interface{}) {
	i.Ui.Output(fmt.Sprintf(format, v...))
}
