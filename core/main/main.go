package main

import (
	"blockchain/core"
	"fmt"
)

func main() {
	bc := core.NewBlockChain()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PreBlockHash)
		//fmt.Printf("PoW: %s\n", strconv.FormatBool())
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
