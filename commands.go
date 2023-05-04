package main

import (
	"os"
	"os/signal"

	"github.com/mitchellh/cli"

	cmd "github.com/tokentransfer/node/command"
)

// Commands is the mapping of all the available node commands.
var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	Commands = map[string]cli.CommandFactory{
		"version": func() (cli.Command, error) {
			return &cmd.VersionCommand{
				Ui: ui,
			}, nil
		},
		"info": func() (cli.Command, error) {
			return &cmd.InfoCommand{
				Ui: ui,
			}, nil
		},
		"start": func() (cli.Command, error) {
			return &cmd.StartCommand{
				Ui:         ui,
				ShutdownCh: makeShutdownCh(),
			}, nil
		},
		"genesis": func() (cli.Command, error) {
			return &cmd.GenesisCommand{
				Ui: ui,
			}, nil
		},
		"dump": func() (cli.Command, error) {
			return &cmd.DumpCommand{
				Ui: ui,
			}, nil
		},
	}
}

// makeShutdownCh returns a channel that can be used for shutdown
// notifications for commands. This channel will send a message for every
// interrupt received.
func makeShutdownCh() <-chan struct{} {
	resultCh := make(chan struct{})

	signalCh := make(chan os.Signal, 4)
	signal.Notify(signalCh, os.Interrupt)
	go func() {
		for {
			<-signalCh
			resultCh <- struct{}{}
		}
	}()

	return resultCh
}
