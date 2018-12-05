package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp    int64
	Data         []byte
	Hash         []byte
	PreBlockHash []byte
}

func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{block.PreBlockHash, block.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	block.Hash=hash[:]
}

func NewBlock(data string, preBlockHash []byte) (block *Block) {
	block = &Block{Timestamp: time.Now().Unix(), Data: []byte(data), PreBlockHash: preBlockHash}
	block.SetHash()
	return
}

//创世块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block",[]byte{})
}
