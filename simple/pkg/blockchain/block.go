package blockchain

import (
	"time"
)

type Block struct {
	Timestamp    int64
	Data         []byte
	PreviousHash []byte
	Hash         []byte
	Nonce        int
}

func NewGeneisBlock() *Block {
	return NewBlock("Genesis block", []byte{})
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Timestamp:    time.Now().Unix(),
		Hash:         []byte{},
		Data:         []byte(data),
		PreviousHash: prevHash,
		Nonce:        0,
	}

	proof := NewProof(block)
	nonce, hash := proof.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
