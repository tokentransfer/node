package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/tokentransfer/node/account"
	"github.com/tokentransfer/node/config"
	"github.com/tokentransfer/node/consensus"
)

// DumpCommand is a Command implementation that generating the genesis block
type DumpCommand struct {
	Ui cli.Ui

	args []string
}

var _ cli.Command = &DumpCommand{}

func (i *DumpCommand) Help() string {
	helpText := `
Usage: node dump [options]
	Provides debugging information for operators
Options:
  -config                  Path to a JSON file to read configuration from.
  -account				   The account
`
	return strings.TrimSpace(helpText)
}

func (i *DumpCommand) Run(args []string) int {
	var configFile string
	var accountString string

	cmdFlags := flag.NewFlagSet("info", flag.ContinueOnError)
	cmdFlags.Usage = func() { i.Ui.Output(i.Help()) }
	cmdFlags.StringVar(&configFile, "config", "./config.json", "json file to read config from")
	cmdFlags.StringVar(&accountString, "account", "", "the account")
	if err := cmdFlags.Parse(i.args); err != nil {
		return 1
	}
	config, err := config.NewConfig(configFile)
	if err != nil {
		panic(err)
	}
	as := account.NewAccountService()
	_, a, err := as.NewAccountFromAddress(accountString)
	if err != nil {
		panic(err)
	}
	n := consensus.NewNode()
	err = n.Init(config)
	if err != nil {
		panic(err)
	}
	err = n.Load(a)
	if err != nil {
		panic(err)
	}
	n.Dump(a, i)
	return 0
}

func (i *DumpCommand) Synopsis() string {
	return "Dump the storage"
}

func (i *DumpCommand) Printf(format string, v ...interface{}) {
	i.Ui.Output(fmt.Sprintf(format, v...))
}
