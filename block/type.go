package block

import (
	"github.com/tokentransfer/node/account"
	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/crypto"

	libblock "github.com/tokentransfer/interfaces/block"
)

var cs = &crypto.CryptoService{}
var as = account.NewAccountService()

// TransactionType
const (
	TRANSACTION libblock.TransactionType = libblock.TransactionType(core.CORE_TRANSACTION)
)

// StateType
const (
	ACCOUNT_STATE libblock.StateType = libblock.StateType(core.CORE_ACCOUNT_STATE)
)

func init() {
	// TransactionType
	TRANSACTION.Register("Transaction", func(t libblock.TransactionType) libblock.Transaction {
		tx := &Transaction{}
		tx.TransactionType = t
		return tx
	})

	// StateType
	ACCOUNT_STATE.Register("AccountState", func(t libblock.StateType) libblock.State {
		info := &AccountState{}
		info.StateType = t
		return info
	})
}
