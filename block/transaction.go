package block

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/core/pb"

	libblock "github.com/tokentransfer/interfaces/block"
	libcore "github.com/tokentransfer/interfaces/core"
)

type Transaction struct {
	Hash libcore.Hash

	TransactionType libblock.TransactionType

	Account     libcore.Address
	Sequence    uint64
	Amount      core.Amount
	Gas         int64
	Destination libcore.Address
	Payload     libcore.Bytes
	PublicKey   libcore.PublicKey
	Signature   libcore.Signature
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

func byteToAddress(b []byte) (libcore.Address, error) {
	_, a, err := as.NewAccountFromBytes(b)
	if err != nil {
		fmt.Println("bytes", len(b), hex.EncodeToString(b))
		return nil, err
	}
	return a, nil
}

func (tx *Transaction) UnmarshalBinary(data []byte) error {
	var err error

	meta, msg, err := core.Unmarshal(data)
	if err != nil {
		return err
	}
	if meta != core.CORE_TRANSACTION {
		return errors.New("error transaction data")
	}
	t := msg.(*pb.Transaction)

	tx.TransactionType = libblock.TransactionType(t.TransactionType)

	tx.Account, err = byteToAddress(t.Account)
	if err != nil {
		return err
	}

	a, err := core.NewAmount(t.Amount)
	if err != nil {
		return err
	}

	tx.Sequence = t.Sequence
	tx.Amount = *a
	tx.Gas = t.Gas

	tx.Destination, err = byteToAddress(t.Destination)
	if err != nil {
		return err
	}

	tx.Payload = t.Payload
	tx.PublicKey = libcore.PublicKey(t.PublicKey)
	tx.Signature = libcore.Signature(t.Signature)

	return nil
}

func addressToByte(a libcore.Address) ([]byte, error) {
	return a.MarshalBinary()
}

func (tx *Transaction) MarshalBinary() ([]byte, error) {
	fromData, err := addressToByte(tx.Account)
	if err != nil {
		return nil, err
	}
	toData, err := addressToByte(tx.Destination)
	if err != nil {
		return nil, err
	}

	t := &pb.Transaction{
		TransactionType: uint32(tx.TransactionType),

		Account:     fromData,
		Sequence:    tx.Sequence,
		Amount:      tx.Amount.String(),
		Gas:         tx.Gas,
		Destination: toData,
		Payload:     tx.Payload,
		PublicKey:   []byte(tx.PublicKey),
		Signature:   []byte(tx.Signature),
	}
	return core.Marshal(t)
}

func (tx *Transaction) Raw(ignoreSigningFields bool) ([]byte, error) {
	if ignoreSigningFields {
		fromAccount, err := addressToByte(tx.Account)
		if err != nil {
			return nil, err
		}
		toAccount, err := addressToByte(tx.Destination)
		if err != nil {
			return nil, err
		}
		t := &pb.Transaction{
			TransactionType: uint32(tx.TransactionType),

			Account:     fromAccount,
			Sequence:    tx.Sequence,
			Amount:      tx.Amount.String(),
			Gas:         tx.Gas,
			Destination: toAccount,
			Payload:     tx.Payload,
			PublicKey:   []byte(tx.PublicKey),
		}
		return core.Marshal(t)
	}
	return tx.MarshalBinary()
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
		return errors.New("error transaction with data")
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
	txData, err := txWithData.Transaction.MarshalBinary()
	if err != nil {
		return nil, err
	}
	_, msg, err := core.Unmarshal(txData)
	if err != nil {
		return nil, err
	}
	tx := msg.(*pb.Transaction)

	receiptData, err := txWithData.Receipt.MarshalBinary()
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

func (txWithData *TransactionWithData) Raw(ignoreSigningFields bool) ([]byte, error) {
	txData, err := txWithData.Transaction.Raw(ignoreSigningFields)
	if err != nil {
		return nil, err
	}
	_, msg, err := core.Unmarshal(txData)
	if err != nil {
		return nil, err
	}
	tx := msg.(*pb.Transaction)

	receiptData, err := txWithData.Receipt.Raw(ignoreSigningFields)
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

func ReadTransaction(data []byte) (libblock.Transaction, error) {
	if len(data) == 0 {
		return nil, errors.New("error transaction")
	}
	meta := core.GetMeta(data)
	switch meta {
	case core.CORE_TRANSACTION:
		tx := &Transaction{}
		err := tx.UnmarshalBinary(data)
		if err != nil {
			return nil, err
		}
		return tx, nil
	default:
		return nil, errors.New("error data")
	}
}
