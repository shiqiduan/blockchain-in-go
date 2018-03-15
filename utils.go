package main

import (
	"bytes"
	"encoding/gob"
	"strconv"
)

func IntToHex(i int64) []byte {
	return []byte(strconv.FormatInt(i, 16))
}
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		panic("DeserializeBlock error.")
	}
	return &block
}
