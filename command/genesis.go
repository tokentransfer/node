package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/tokentransfer/node/account"
	"github.com/tokentransfer/node/config"
	"github.com/tokentransfer/node/consensus"
	"github.com/tokentransfer/node/util"

	libcore "github.com/tokentransfer/interfaces/core"
)

// GenesisCommand is a Command implementation that generating the genesis block
type GenesisCommand struct {
	Ui cli.Ui
}

var _ cli.Command = &GenesisCommand{}

func (i *GenesisCommand) Help() string {
	helpText := `
Usage: 	node genesis [options]
	Generate the genesis block
Options:
  -config		Path to a JSON file to read configuration from.
  -account		The account
`
	return strings.TrimSpace(helpText)
}

func (i *GenesisCommand) Run(args []string) int {
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

	n := consensus.NewNode()
	err = n.Init(config)
	if err != nil {
		panic(err)
	}
	err = n.Load(account)
	if err != nil {
		panic(err)
	}
	if n.GetBlockNumber(account) >= 0 {
		i.Ui.Error(fmt.Sprintf("=== %s: load block %d, %s, %d\n", util.GetString(account), n.GetBlockNumber(account), n.GetBlockHash(account), len(n.GetBlock(account).GetTransactions())))
		return 1
	}

	block, err := n.GenerateBlock(account)
	if err != nil {
		panic(err)
	}
	_, err = n.HashBlock(block)
	if err != nil {
		panic(err)
	}
	_, err = n.VerifyBlock(account, block)
	if err != nil {
		panic(err)
	}
	err = n.AddBlock(account, block)
	if err != nil {
		panic(err)
	}
	i.Ui.Output(fmt.Sprintf("=== %s: generate block %d, %s, %d\n", util.GetString(account), block.GetIndex(), block.GetHash().String(), len(block.GetTransactions())))
	return 0
}

func (i *GenesisCommand) Synopsis() string {
	return "Generating the genesis block"
}
