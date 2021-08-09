package main

import (
	"fmt"
	"github.com/MinSeo123/nomadcoin/blockchain"
)


func main(){
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")
	for _, block := range chain.AllBlocks() {
		fmt.Println("data:", block.Data)
		fmt.Println("hash:", block.Hash)
		fmt.Println("prev:" ,block.PrevHash)
	}

}

