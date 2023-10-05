//block.go file
package main

import (
	"bytes"
	"crypto/sha256" //crypto library to hash the data
	"encoding/hex"
	"strconv" // for conversion
	"strings"
	"time" // the time for our timestamp
)

func (block *Block) calculateHash() []byte {
    headers := bytes.Join([][]byte{
        []byte(strconv.FormatInt(block.Timestamp, 10)),
        block.PreviousBlockHash,
        block.AllData,
        []byte(strconv.FormatInt(int64(block.Nonce), 10)),
    }, []byte{})
    hash := sha256.Sum256(headers)
    return hash[:]
}

// Create a function for new blcok generation and return that block
func NewBlock(data string, prevBlockHash []byte, difficulty int) *Block {

	block := &Block{
		Timestamp: time.Now().Unix(), 
		PreviousBlockHash: prevBlockHash, 
		AllData: []byte(data), 
		Nonce: 0, 
		Difficulty: difficulty,
	}
	block.mineBlock()
	return block
}

func (block *Block) mineBlock() {
	tragetPrefix := strings.Repeat("0", block.Difficulty)
	for {
		hash := block.calculateHash()
		hashStr := hex.EncodeToString(hash)

		if strings.HasPrefix(hashStr, tragetPrefix) {
			block.MyBlockHash = hash
			break
		}
		block.Nonce++
	}
}

// THE GENESIS BLCOK
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}, 4)
}