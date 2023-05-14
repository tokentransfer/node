package block

import (
	"log"

	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/core/pb"
	"github.com/tokentransfer/node/util"

	libblock "github.com/tokentransfer/interfaces/block"
	libcore "github.com/tokentransfer/interfaces/core"
	libcrypto "github.com/tokentransfer/interfaces/crypto"
)

type Transaction struct {
	Hash libcore.Hash

	TransactionType libblock.TransactionType

	Account  libcore.Address
	Sequence uint64
	Gas      uint64
	Payload  *PayloadInfo

	PublicKey libcore.PublicKey
	Signature libcore.Signature
}

func (tx *Transaction) GetIndex() uint64 {
	return tx.Sequence
}

func (tx *Transaction) GetHash() libcore.Hash {
	return tx.Hash
}

func (tx *Transaction) SetHash(h libcore.Hash) {
	tx.Hash = h
}

func (tx *Transaction) UnmarshalBinary(data []byte) error {
	var err error

	meta, msg, err := core.Unmarshal(data)
	if err != nil {
		return err
	}
	if meta != core.CORE_TRANSACTION {
		return util.ErrorOfUnmatched("data type", "transaction with data", core.CORE_TRANSACTION, meta)
	}
	t := msg.(*pb.Transaction)

	tx.TransactionType = libblock.TransactionType(t.TransactionType)

	tx.Account, err = ByteToAddress(t.Account)
	if err != nil {
		return err
	}
	tx.Sequence = t.Sequence
	tx.Gas = t.Gas
	tx.Payload = fromPayloadInfo(t.Payload)

	tx.PublicKey = libcore.PublicKey(t.PublicKey)
	tx.Signature = libcore.Signature(t.Signature)

	return nil
}

func (tx *Transaction) MarshalBinary() ([]byte, error) {
	return tx.Raw(libcrypto.RawBinary)
}

func (tx *Transaction) Raw(rt libcrypto.RawType) ([]byte, error) {
	switch rt {
	case libcrypto.RawBinary:
		{
			accountData, err := AddressToByte(tx.Account)
			if err != nil {
				return nil, err
			}

			t := &pb.Transaction{
				TransactionType: uint32(tx.TransactionType),

				Account:  accountData,
				Sequence: tx.Sequence,
				Gas:      tx.Gas,
				Payload:  toPayloadInfo(tx.Payload, libcrypto.RawBinary),

				PublicKey: []byte(tx.PublicKey),
				Signature: []byte(tx.Signature),
			}
			return core.Marshal(t)
		}
	case libcrypto.RawIgnoreSigningFields:
		{
			account, err := AddressToByte(tx.Account)
			if err != nil {
				return nil, err
			}
			t := &pb.Transaction{
				TransactionType: uint32(tx.TransactionType),

				Account:  account,
				Sequence: tx.Sequence,
				Gas:      tx.Gas,
				Payload:  toPayloadInfo(tx.Payload, libcrypto.RawIgnoreSigningFields),

				PublicKey: []byte(tx.PublicKey),
			}
			return core.Marshal(t)
		}
	case libcrypto.RawIgnoreVariableFields:
		{
			account, err := AddressToByte(tx.Account)
			if err != nil {
				return nil, err
			}

			t := &pb.Transaction{
				TransactionType: uint32(tx.TransactionType),

				Account:  account,
				Sequence: tx.Sequence,
				Gas:      tx.Gas,
				Payload:  toPayloadInfo(tx.Payload, libcrypto.RawIgnoreVariableFields),

				PublicKey: []byte(tx.PublicKey),
				Signature: []byte(tx.Signature),
			}
			return core.Marshal(t)
		}
	}
	account, err := AddressToByte(tx.Account)
	if err != nil {
		return nil, err
	}

	t := &pb.Transaction{
		TransactionType: uint32(tx.TransactionType),

		Account:  account,
		Sequence: tx.Sequence,
		Gas:      tx.Gas,
		Payload:  toPayloadInfo(tx.Payload, rt),

		PublicKey: []byte(tx.PublicKey),
		Signature: []byte(tx.Signature),
	}
	return core.Marshal(t)
}

func (tx *Transaction) GetTransactionType() libblock.TransactionType {
	return tx.TransactionType
}

func (tx *Transaction) GetAccount() libcore.Address {
	return tx.Account
}

func (tx *Transaction) GetPublicKey() libcore.PublicKey {
	return tx.PublicKey
}

func (tx *Transaction) SetPublicKey(p libcore.PublicKey) {
	tx.PublicKey = p
}

func (tx *Transaction) GetSignature() libcore.Signature {
	return tx.Signature
}

func (tx *Transaction) SetSignature(s libcore.Signature) {
	tx.Signature = s
}

type TransactionWithData struct {
	libblock.Transaction

	Receipt libblock.Receipt

	Date int64
	Hash libcore.Hash
}

func (txWithData *TransactionWithData) SetHash(h libcore.Hash) {
	txWithData.Hash = h
}

func (txWithData *TransactionWithData) GetTransaction() libblock.Transaction {
	return txWithData.Transaction
}

func (txWithData *TransactionWithData) GetReceipt() libblock.Receipt {
	return txWithData.Receipt
}

func (txWithData *TransactionWithData) UnmarshalBinary(data []byte) error {
	meta, msg, err := core.Unmarshal(data)
	if err != nil {
		return err
	}
	if meta != core.CORE_TRANSACTION_WITH_DATA {
		return util.ErrorOfUnmatched("data type", "transaction with data", core.CORE_TRANSACTION_WITH_DATA, meta)
	}

	td := msg.(*pb.TransactionWithData)

	txData, err := core.Marshal(td.Transaction)
	if err != nil {
		return err
	}
	tx := &Transaction{}
	err = tx.UnmarshalBinary(txData)
	if err != nil {
		return err
	}

	receiptData, err := core.Marshal(td.Receipt)
	if err != nil {
		log.Println(err)
		return err
	}
	receipt := &Receipt{}
	err = receipt.UnmarshalBinary(receiptData)
	if err != nil {
		log.Println(err)
		return err
	}

	txWithData.Transaction = tx
	txWithData.Receipt = receipt
	txWithData.Date = td.Date
	return nil
}

func (txWithData *TransactionWithData) MarshalBinary() ([]byte, error) {
	return txWithData.Raw(libcrypto.RawBinary)
}

func (txWithData *TransactionWithData) Raw(rt libcrypto.RawType) ([]byte, error) {
	txData, err := txWithData.Transaction.Raw(rt)
	if err != nil {
		return nil, err
	}
	_, msg, err := core.Unmarshal(txData)
	if err != nil {
		return nil, err
	}
	tx := msg.(*pb.Transaction)

	receiptData, err := txWithData.Receipt.Raw(rt)
	if err != nil {
		return nil, err
	}
	_, msg, err = core.Unmarshal(receiptData)
	if err != nil {
		return nil, err
	}
	receipt := msg.(*pb.Receipt)

	td := &pb.TransactionWithData{
		Transaction: tx,
		Receipt:     receipt,
		Date:        txWithData.Date,
	}
	data, err := core.Marshal(td)
	if err != nil {
		return nil, err
	}
	return data, nil
}
