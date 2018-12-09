package core

import (
	"github.com/boltdb/bolt"
	"log"
)

//区块链表
type BlockChain struct {
	Tip []byte
	Db  *bolt.DB
}

//添加一个区块
func (bc *BlockChain) AddBlock(data string) {
	var lastHash []byte
	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		panic(err)
	}
	newBlock := NewBlock(data, lastHash)

	err = bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err = b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		bc.Tip = newBlock.Hash
		return nil
	})
}

const dbFile = "blockchain.Db"
const blocksBucket = "blocks"

func NewBlockChain() *BlockChain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))

		if bucket == nil {
			genesis := NewGenesisBlock()
			bucket, err = tx.CreateBucket([]byte(blocksBucket))
			err = bucket.Put(genesis.Hash, genesis.Serialize())
			err = bucket.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = bucket.Get([]byte("l"))
		}
		return err
	})
	bc := BlockChain{
		Tip: tip,
		Db:  db,
	}
	return &bc
}

type BlockChainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (bc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{
		currentHash: bc.Tip,
		db:          bc.Db,
	}
}

func (iter *BlockChainIterator) Next() (block *Block) {
	_ = iter.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		encodeBlock := bucket.Get(iter.currentHash)
		block = DeserializeBlock(encodeBlock)
		return nil
	})
	iter.currentHash = block.PreBlockHash
	return block
}
