package command

import (
	"flag"
	"strings"

	"github.com/mitchellh/cli"

	"github.com/tokentransfer/node/account"
	"github.com/tokentransfer/node/config"
	"github.com/tokentransfer/node/util"

	libcore "github.com/tokentransfer/interfaces/core"
)

// InfoCommand is a Command implementation that queries a running
// Node agent for various debugging statistics for operators
type InfoCommand struct {
	Ui cli.Ui
}

var _ cli.Command = &InfoCommand{}

func (i *InfoCommand) Help() string {
	helpText := `
Usage: node info [options]
	Provides debugging information for operators
Options:
	-config		Path to a JSON file to read configuration from.
	-account		The account
`
	return strings.TrimSpace(helpText)
}

func (i *InfoCommand) Run(args []string) int {
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

	if len(config.GetSecret()) == 0 {
		secret, err := i.Ui.AskSecret("The secret:")
		if err != nil {
			panic(err)
		}
		config.SetSecret(secret)
	}

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
	return 0
}

func (i *InfoCommand) Synopsis() string {
	return "Provides debugging information for operators"
}
