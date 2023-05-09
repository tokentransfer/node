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

// GenesisCommand is a Command implementation that generating the genesis block
type GenesisCommand struct {
	Ui cli.Ui

	args []string
}

var _ cli.Command = &GenesisCommand{}

func (i *GenesisCommand) Help() string {
	helpText := `
Usage: node genesis [options]
	Provides debugging information for operators
Options:
  -config                  Path to a JSON file to read configuration from.
  -account				   The account
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
	if n.GetBlockNumber(a) >= 0 {
		i.Ui.Error(fmt.Sprintf("=== load block %d, %s, %d\n", n.GetBlockNumber(a), n.GetBlockHash(a), len(n.GetBlock(a).GetTransactions())))
		return 1
	}

	block, err := n.GenerateBlock(a)
	if err != nil {
		panic(err)
	}
	_, err = n.HashBlock(block)
	if err != nil {
		panic(err)
	}
	_, err = n.VerifyBlock(a, block)
	if err != nil {
		panic(err)
	}
	err = n.AddBlock(a, block)
	if err != nil {
		panic(err)
	}
	i.Ui.Output(fmt.Sprintf("=== generate block %d, %s, %d\n", block.GetIndex(), block.GetHash().String(), len(block.GetTransactions())))
	return 0
}

func (i *GenesisCommand) Synopsis() string {
	return "Generating the genesis block"
}
