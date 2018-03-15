package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

type Block struct {
	Timestamp     int64
	PrevBlockHash []byte
	Hash          []byte
	Data          []byte
	Nonce         int
}

func NewBlock(data string, PrevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), PrevBlockHash, []byte{}, []byte(data), 0}
	pow := NewProofOfWork(block)
	block.Nonce, block.Hash = pow.Run()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		panic("serialize error.")
	}
	return result.Bytes()
}

func (b *Block) String() string {
	return fmt.Sprintf("Hash: %x\nPreHash: %x\nData: %s\nTimestamp: %b\n",
		b.Hash, b.PrevBlockHash, b.Data, b.Timestamp)
}
