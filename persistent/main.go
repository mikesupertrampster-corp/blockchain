package main

import (
	persistent "github.com/mikesupertrampster-corp/blockchain/persistent/pkg/blockchain"
	"log"
)

func main() {
	bc, err := persistent.NewBlockchain()
	if err != nil {
		log.Panic(err)
	}

	defer bc.DB.Close()

	cli := persistent.CLI{Blockchain: bc}
	cli.Run()
}
