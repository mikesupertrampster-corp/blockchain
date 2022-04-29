package blockchain

import (
	"fmt"
	"time"
)

type BlockChain struct {
	blocks []*Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGeneisBlock()}}
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc BlockChain) Print() {
	for i, block := range bc.blocks {
		fmt.Println("----------------")
		fmt.Printf("block #%d\n", i)
		fmt.Println("----------------")
		fmt.Printf("  TIMESTAMP: %s\n", time.Unix(block.Timestamp, 0))
		fmt.Printf("  HASH: %x\n", block.Hash)
		fmt.Printf("  PREVIOUS: %x\n", block.PreviousHash)
		fmt.Printf("  DATA: %s\n", block.Data)
	}
}
