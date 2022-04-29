package main

import "simple/pkg/blockchain"

func main() {
	chain := blockchain.NewBlockChain()

	chain.AddBlock("Send 1 BTC to Mike")
	chain.AddBlock("Send 2 more BTC to Mike")

	chain.Print()
}
