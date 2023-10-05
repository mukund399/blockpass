package main

import (
    "fmt"
    "testing"
)

func TestMultipleBlocks(t *testing.T) {
    blockchain := NewBlockchain()
    initialBlockCount := len(blockchain.Blocks)
    fmt.Printf("Initial blockchain length: %d\n", initialBlockCount)

    // Add two valid blocks
    for i := 0; i < 2; i++ {
        newBlock := NewBlock(fmt.Sprintf("valid transaction %d", i), blockchain.Blocks[len(blockchain.Blocks)-1].MyBlockHash, 6)
        //blockchain.mineBlock(newBlock)
        blockchain.Blocks = append(blockchain.Blocks, newBlock)
    }

    // Add an invalid block
    invalidBlock := NewBlock("invalid transaction", blockchain.Blocks[len(blockchain.Blocks)-1].MyBlockHash, 5)
    invalidBlock.MyBlockHash = []byte("invalidhashnotendingwithzeros")
    blockchain.Blocks = append(blockchain.Blocks, invalidBlock)

    // Call the TestHashEndsWithSixZeros() function to check the blocks
    TestHashEndsWithSixZeros(t)

    finalBlockCount := len(blockchain.Blocks)
    fmt.Printf("Final blockchain length: %d\n", finalBlockCount)
}


func TestHashEndsWithSixZeros(t *testing.T) {
    blockchain := NewBlockchain()

    // Loop through the blocks and check hash endings
    for i, block := range blockchain.Blocks {
        // Convert the hash to hexadecimal and get the last 6 characters
        hashedSuffix := fmt.Sprintf("%x", block.MyBlockHash)[len(block.MyBlockHash)*2-6:]

        // Check that the hashed suffix is all zeros
        if hashedSuffix != "000000" {
            t.Errorf("Block %d: Hashed suffix doesn't end with 6 zeros", i)
        }
    }
}