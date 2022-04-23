package main

import (
	"blockchain/pkg/blockchain"
	"fmt"
)

func main() {
	blockchain := blockchain.CreateBlockchain(2)

	blockchain.AddBlock("Alice", "Bob", 5)
	blockchain.AddBlock("John", "Bob", 2)

	blockchain.Print()

	fmt.Println("-------------------------------------------------------")
	fmt.Printf("Chain is Valid: %t\n", blockchain.IsValid())
	fmt.Println("-------------------------------------------------------")
}
