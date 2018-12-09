package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp    int64
	Data         []byte
	Hash         []byte
	PreBlockHash []byte
	Nonce        int
}

func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{block.PreBlockHash, block.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	block.Hash = hash[:]
}

func NewBlock(data string, preBlockHash []byte) (block *Block) {
	block = &Block{Timestamp: time.Now().Unix(), Data: []byte(data), PreBlockHash: preBlockHash}
	pow := NewProofOfWork(block)
	none, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = none
	block.SetHash()
	return
}

//创世块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		fmt.Println(err)
	}
	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		fmt.Println(err)
	}
	return &block
}
