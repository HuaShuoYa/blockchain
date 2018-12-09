package core

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CLI struct {
	Bc *BlockChain
}

const usage = `
Usage:
  addblock -data BLOCK_DATA    add a block to the blockchain
  printchain                   print all the blocks of the blockchain
`

func (cli *CLI) printUsage() {
	fmt.Println(usage)
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}

}

func (cli *CLI) printChain() {
	bci := cli.Bc.Iterator()
	for {
		block := bci.Next()
		fmt.Printf("Prev hash:%x\n", block.PreBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW :%s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PreBlockHash) == 0 {
			break
		}
	}

}

func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("date", "", "Block data")
	switch os.Args[1] {
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}
	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.Bc.AddBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printUsage()
	}
}
