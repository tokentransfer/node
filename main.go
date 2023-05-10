package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	args := os.Args[1:]
	for index, arg := range args {
		if arg == "-v" || arg == "--version" {
			newArgs := make([]string, len(args)+1)
			newArgs[0] = "version"
			copy(newArgs[1:], args)
			args = newArgs
			break
		} else if index == 0 && (arg == "-h" || arg == "--help") {
			newArgs := make([]string, len(args)+1)
			newArgs[0] = "help"
			copy(newArgs[1:], args)
			args = newArgs
			break
		}
	}
	err := flag.CommandLine.Parse(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parse arguments: %s\n", err.Error())
		os.Exit(1)
	}

	cli := &cli.CLI{
		Args:     args,
		Commands: Commands,
		HelpFunc: cli.BasicHelpFunc("node"),
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		os.Exit(1)
	}

	os.Exit(exitCode)
}
