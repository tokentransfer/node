package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/tokentransfer/node/conf"
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
`
	return strings.TrimSpace(helpText)
}

func (i *GenesisCommand) Run(args []string) int {
	var configFile string

	cmdFlags := flag.NewFlagSet("info", flag.ContinueOnError)
	cmdFlags.Usage = func() { i.Ui.Output(i.Help()) }
	cmdFlags.StringVar(&configFile, "config", "./config.json", "json file to read config from")
	if err := cmdFlags.Parse(i.args); err != nil {
		return 1
	}

	config, err := conf.NewConfig(configFile)
	if err != nil {
		panic(err)
	}

	n := consensus.NewNode()
	err = n.Init(config)
	if err != nil {
		panic(err)
	}

	n.Load()
	if n.GetBlockNumber() >= 0 {
		i.Ui.Error(fmt.Sprintf("=== load block %d, %s, %d\n", n.GetBlockNumber(), n.GetBlockHash(), len(n.GetBlock().GetTransactions())))
		return 1
	}

	block, err := n.GenerateBlock()
	if err != nil {
		panic(err)
	}
	_, err = n.HashBlock(block)
	if err != nil {
		panic(err)
	}
	_, err = n.VerifyBlock(block)
	if err != nil {
		panic(err)
	}
	err = n.AddBlock(block)
	if err != nil {
		panic(err)
	}
	i.Ui.Output(fmt.Sprintf("=== generate block %d, %s, %d\n", block.GetIndex(), block.GetHash().String(), len(block.GetTransactions())))
	return 0
}

func (i *GenesisCommand) Synopsis() string {
	return "Generating the genesis block"
}
