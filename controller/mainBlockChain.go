package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index      int
	Timestamp  string
	BPM        int
	Hash       string
	PrevHash   string
	minerIndex int
}

var Blockchain []Block

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.BPM) + block.PrevHash + strconv.Itoa(block.minerIndex)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, BPM int, minerIndex int) (Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.minerIndex = minerIndex
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

var isRunning bool = true

func stopMainBlockChain() {
	isRunning = false
}

func continueMainBlockChain() {
	isRunning = true
}

func runMainBlockChain() {
	t := time.Now()
	genesisBlock := Block{0, t.String(), 0, "", "", 0}
	genesisBlock = Block{0, t.String(), 0, calculateHash(genesisBlock), "", 0}
	Blockchain = append(Blockchain, genesisBlock)

	tick := time.Tick(20 * time.Second)

	for range tick {
		if !isRunning {
			break
		}
		runPosNet()
		newMinerIndex := SelectWalletSlice[len(SelectWalletSlice)-1]
		block, err := generateBlock(Blockchain[len(Blockchain)-1], len(Blockchain), newMinerIndex)
		if err != nil {
			fmt.Println(err)
			return
		}
		if isBlockValid(block, Blockchain[len(Blockchain)-1]) {
			Blockchain = append(Blockchain, block)
		}
	}
}

//func RunMainBlockChain(runTime int) {
//	t := time.Now()
//	genesisBlock := Block{0, t.String(), 0, "", "", 0}
//	genesisBlock = Block{0, t.String(), 0, calculateHash(genesisBlock), "", 0}
//	Blockchain = append(Blockchain, genesisBlock)
//
//	tick := time.Tick(2 * time.Second)
//	stop := time.After(time.Duration(runTime) * time.Second)
//
//	for {
//		select {
//		case <-tick:
//			runPosNet()
//			newMinerIndex := SelectWalletSlice[len(SelectWalletSlice)-1]
//			block, err := generateBlock(Blockchain[len(Blockchain)-1], len(Blockchain), newMinerIndex)
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//			if isBlockValid(block, Blockchain[len(Blockchain)-1]) {
//				Blockchain = append(Blockchain, block)
//			}
//		case <-stop:
//			return
//		}
//	}
//}

func init() {
	go runMainBlockChain()

	// 模拟停止和继续运行的操作
	time.Sleep(5 * time.Second)
	stopMainBlockChain()
	time.Sleep(5 * time.Second)
	continueMainBlockChain()
}
