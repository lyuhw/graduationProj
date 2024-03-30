package server

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

func calculateHash(block Block) string {
	if block.Index == 0 || block.Timestamp == "" || block.Data == "" || block.PrevHash == "" {
		return ""
	}
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func CreateBlock(data string, blockchain []Block) Block {
	if data == "" {
		return Block{}
	}
	block := Block{}
	block.Index = len(blockchain) + 1
	block.Timestamp = time.Now().String()
	block.Data = data
	if len(blockchain) == 0 {
		block.PrevHash = "0"
	} else {
		prevBlock := blockchain[len(blockchain)-1]
		block.PrevHash = prevBlock.Hash
	}
	block.Hash = calculateHash(block)
	return block
}

func CreateGenesisBlock(blockchain []Block) Block {
	if len(blockchain) != 0 {
		newBlock := CreateBlock("Genesis Block", []Block{}) //
		blockchain = append(blockchain, newBlock)
		return newBlock
	} else {
		return Block{}
	}
}

func AddBlock(data string, blockchain []Block) Block {
	if len(blockchain) == 0 {
		newBlock := CreateBlock("Genesis Block", []Block{}) //
		blockchain = append(blockchain, newBlock)
		return newBlock
	}
	newBlock := CreateBlock(data, blockchain) //
	blockchain = append(blockchain, newBlock)
	return newBlock

}

func PrintAllBlock(blockchain []Block) {
	for _, block := range blockchain {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println("-------------")
	}
}
