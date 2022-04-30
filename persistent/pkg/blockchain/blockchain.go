package blockchain

import (
	"github.com/boltdb/bolt"
	simple "github.com/mikesupertrampster-corp/blockchain/simple/pkg/blockchain"
)

const dbFile = "blockchain.DB"
const blocksBucket = "blocks"

var lastBlockPointer = []byte("l")

type Blockchain struct {
	lastBlock []byte
	DB        *bolt.DB
}

type Iterator struct {
	currentHash []byte
	db          *bolt.DB
}

func NewBlockchain() (*Blockchain, error) {
	var lastBlock []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := simple.NewGeneisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				return err
			}

			SerializeBlock, err := Serialize(genesis)
			if err != nil {
				return err
			}

			err = b.Put(genesis.Hash, SerializeBlock)
			if err != nil {
				return err
			}

			err = b.Put(lastBlockPointer, genesis.Hash)
			if err != nil {
				return err
			}

			lastBlock = genesis.Hash
		} else {
			lastBlock = b.Get(lastBlockPointer)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	bc := Blockchain{lastBlock, db}
	return &bc, nil
}

func (bc *Blockchain) AddBlock(data string) error {
	var lastHash []byte

	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get(lastBlockPointer)
		return nil
	})
	if err != nil {
		return err
	}

	newBlock := simple.NewBlock(data, lastHash)

	err = bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		SerializeBlock, err := Serialize(newBlock)
		if err != nil {
			return err
		}

		err = b.Put(newBlock.Hash, SerializeBlock)
		if err != nil {
			return err
		}

		err = b.Put(lastBlockPointer, newBlock.Hash)
		if err != nil {
			return err
		}

		bc.lastBlock = newBlock.Hash
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (bc Blockchain) Iterator() *Iterator {
	bci := &Iterator{
		currentHash: bc.lastBlock,
		db:          bc.DB,
	}
	return bci
}

func (bci Iterator) Next() (*simple.Block, error) {
	var block *simple.Block

	err := bci.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(bci.currentHash)
		block, _ = DeserializeBlock(encodedBlock)
		return nil
	})
	if err != nil {
		return nil, err
	}

	bci.currentHash = block.PreviousHash
	return block, nil
}
