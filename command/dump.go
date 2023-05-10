package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
	libcore "github.com/tokentransfer/interfaces/core"
	"github.com/tokentransfer/node/account"
	"github.com/tokentransfer/node/config"
	"github.com/tokentransfer/node/consensus"
	"github.com/tokentransfer/node/util"
)

// DumpCommand is a Command implementation that generating the genesis block
type DumpCommand struct {
	Ui cli.Ui
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
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	config, err := config.NewConfig(configFile)
	if err != nil {
		panic(err)
	}
	as := account.NewAccountService()

	_, nodeKey, err := as.NewKeyFromSecret(config.GetSecret())
	if err != nil {
		panic(err)
	}
	nodeAccount, err := nodeKey.GetAddress()
	if err != nil {
		panic(err)
	}
	i.Ui.Info("Node: " + util.GetString(nodeAccount))

	var account libcore.Address
	if len(accountString) > 0 {
		_, a, err := as.NewAccountFromAddress(accountString)
		if err != nil {
			panic(err)
		}
		account = a
	} else {
		account = nil
	}
	i.Ui.Info("Account: " + util.GetString(account))

	n := consensus.NewNode()
	err = n.Init(config)
	if err != nil {
		panic(err)
	}
	err = n.Load(account)
	if err != nil {
		panic(err)
	}
	n.Dump(account, i)
	return 0
}

func (i *DumpCommand) Synopsis() string {
	return "Dump the storage"
}

func (i *DumpCommand) Printf(format string, v ...interface{}) {
	i.Ui.Output(fmt.Sprintf(format, v...))
}
