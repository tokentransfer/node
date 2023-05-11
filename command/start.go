package command

import (
	"flag"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/caivega/glog"
	"github.com/mitchellh/cli"
	libcore "github.com/tokentransfer/interfaces/core"
	"github.com/tokentransfer/node/account"
	"github.com/tokentransfer/node/config"
	"github.com/tokentransfer/node/consensus"
	"github.com/tokentransfer/node/rpc"
	"github.com/tokentransfer/node/util"
)

const (
	// gracefulTimeout controls how long we wait before forcefully terminating
	gracefulTimeout = 3 * time.Second
)

// Command is a Command implementation that runs a Node Node.
// The command will not end unless a shutdown message is sent on the
// ShutdownCh. If two messages are sent on the ShutdownCh it will forcibly
// exit.
type StartCommand struct {
	Ui         cli.Ui
	args       []string
	ShutdownCh <-chan struct{}
}

var _ cli.Command = &StartCommand{}

// readConfig is responsible for setup of our configuration using
// the command line and any file configs
func (c *StartCommand) readConfig() (*config.Config, libcore.Address) {
	var configFile string
	var accountString string

	cmdFlags := flag.NewFlagSet("start", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	cmdFlags.StringVar(&configFile, "config", "./config.json", "json file to read config from")
	cmdFlags.StringVar(&accountString, "account", "", "the account")
	if err := cmdFlags.Parse(c.args); err != nil {
		panic(err)
	}

	config, err := config.NewConfig(configFile)
	if err != nil {
		panic(err)
	}

	if len(config.GetSecret()) == 0 {
		secret, err := c.Ui.AskSecret("The secret:")
		if err != nil {
			panic(err)
		}
		config.SetSecret(secret)
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
	c.Ui.Info("Node: " + util.GetString(nodeAccount))

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
	c.Ui.Info("Account: " + util.GetString(account))

	return config, account
}

// startNode is used to start the Node and IPC
func (c *StartCommand) startNode(config *config.Config, account libcore.Address) *consensus.Node {
	c.Ui.Output("Starting node...")

	n := consensus.NewNode()
	err := n.Init(config)
	if err != nil {
		panic(err)
	}
	err = n.Start()
	if err != nil {
		panic(err)
	}

	entry := n.GetEntry(account)
	blockNumber := entry.GetBlockNumber()
	blockHash := entry.GetBlockHash()
	glog.Infof("%s: block, %d, %s", util.GetString(account), blockNumber, blockHash)

	r := rpc.NewRPCService(n)
	err = r.Init(config)
	if err != nil {
		panic(err)
	}
	go r.Start()

	return n
}

func (c *StartCommand) Run(args []string) int {
	c.Ui = &cli.PrefixedUi{
		OutputPrefix: "==> ",
		InfoPrefix:   "    ",
		ErrorPrefix:  "==> ",
		Ui:           c.Ui,
	}

	// Parse our configs
	c.args = args
	config, account := c.readConfig()
	if config == nil {
		return 1
	}

	// Start the Node
	n := c.startNode(config, account)
	if n == nil {
		return 1
	}
	defer n.Stop()

	// Wait for exit
	return c.handleSignals(config, n)
}

// handleSignals blocks until we get an exit-causing signal
func (c *StartCommand) handleSignals(config *config.Config, n *consensus.Node) int {
	signalCh := make(chan os.Signal, 4)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)

	// Wait for a signal
	var sig os.Signal
	select {
	case s := <-signalCh:
		sig = s
	case <-c.ShutdownCh:
		sig = os.Interrupt
		return 1
	}

	// Check if we should do a graceful leave
	graceful := false
	if sig == os.Interrupt {
		graceful = true
	} else if sig == syscall.SIGTERM {
		graceful = true
	}

	if !graceful {
		return 1
	}

	// Attempt a graceful leave
	gracefulCh := make(chan struct{})
	go func() {
		if err := n.Stop(); err != nil {
			panic(err)
		}
		close(gracefulCh)
	}()

	// Wait for leave or another signal
	select {
	case <-signalCh:
		return 1
	case <-time.After(gracefulTimeout):
		return 1
	case <-gracefulCh:
		return 0
	}
}

func (c *StartCommand) Synopsis() string {
	return "Runs a node"
}

func (c *StartCommand) Help() string {
	helpText := `
Usage: node start [options]

  Starts the node and runs until an interrupt is received. 

Options:

	-config=./config.json	Path to a JSON file to read configuration from.
	-account		The account
`
	return strings.TrimSpace(helpText)
}
