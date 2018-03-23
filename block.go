package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

type Block struct {
	Timestamp     int64
	PrevBlockHash []byte
	Hash          []byte
	Transactions  []*Transaction
	Nonce         int
	Height        int
}

func NewBlock(transactions []*Transaction, PrevBlockHash []byte, height int) *Block {
	block := &Block{time.Now().Unix(), PrevBlockHash, []byte{}, transactions, 0, height}
	pow := NewProofOfWork(block)
	block.Nonce, block.Hash = pow.Run()
	return block
}

func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{}, 0)
}

func (b *Block) HashTransactions() []byte {
	var transactions [][]byte
	for _, tx := range b.Transactions {
		transactions = append(transactions, tx.Serialize())
	}
	mTree := NewMerkleTree(transactions)
	return mTree.RootNode.Data
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		panic(err)
	}
	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

func (b *Block) String() string {
	return fmt.Sprintf("Hash: %x\nPreHash: %x\nTransactions: %s\nTimestamp: %b\n",
		b.Hash, b.PrevBlockHash, b.Transactions, b.Timestamp)
}
