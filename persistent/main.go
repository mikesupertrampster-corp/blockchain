package main

import (
	"github.com/mikesupertrampster-corp/blockchain/simple/pkg/blockchain"
	"time"
)

func main() {
	chain := blockchain.NewBlockChain()

	chain.AddBlock("Send 1 BTC to Mike")
	time.Sleep(3 * time.Second)
	chain.AddBlock("Send 2 more BTC to Mike")

	chain.Print()
}
