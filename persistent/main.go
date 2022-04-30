package blockchain

func main() {
	bc, _ := persistent.NewBlockchain()
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}
