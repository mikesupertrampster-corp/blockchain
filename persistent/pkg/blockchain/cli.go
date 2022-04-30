package blockchain

import (
	"fmt"
	simple "github.com/mikesupertrampster-corp/blockchain/simple/pkg/blockchain"
	"log"
	"os"
	"strconv"
)

type CLI struct {
	bc *Blockchain
}

func (cli CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("  printchain - print all the blocks of the blockchain")
}

func (cli CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli CLI) addBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("Success!")
}

func (cli CLI) printChain() {
	bci := cli.bc.Iterator()

	for {
		block, err := bci.Next()
		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("Prev. hash %x\n", block.PreviousHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash %x\n", block.Hash)
		pow := simple.NewProof(block)
		fmt.Printf("PoW %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PreviousHash) == 0 {
			break
		}
	}
}

func (cli CLI) Run() {
	cli.validateArgs()

}
