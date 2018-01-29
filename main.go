package main
import (
	"fmt"
	"go-block/block"
	"strconv"
)
	
func main() { 
	bc := block.NewBlockchain()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")
	bc.AddBlock("这是第三个块")
	fmt.Println(bc)
	for _, b:= range bc.Blocks {
		pow := block.NewProofOfWork(b)
		fmt.Printf("Prev. hash: %x\n", b.PrevBlockHash)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
        fmt.Printf("Data: %s\n", b.Data)
        fmt.Printf("Hash: \r%x", b.Hash)
        fmt.Println()
	}
}