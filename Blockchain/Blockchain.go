// Blockchain.go file
package main

import (
    "bytes"
	"encoding/json"
	"fmt"
	"net/http"
    "strings"

	//"github.com/gorilla/mux"
)

// Create the function that returns the whole blockchain and add teh genesis to it first ever mined blcok. Function that will return it since it does not exist yet
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func (blockchain *Blockchain) HandleAddData(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Parse the request body to extract the data
    var requestBody struct {
        Data string `json:"data"`
    }

    if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Error decoding request body: %s", err)
        return
    }

    var m string
    // Add the data to the blockchain
    if blockchain.AddBlock(requestBody.Data) {
        m = "Data added to the blockchain successfully."
    } else {
        m = "Block contains duplicate data."
    }

    // Send a '200' response with a success message
    response := struct {
        Message string `json:"message"`
    }{
        Message: m,
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Error encoding JSON response: %s", err)
    }
}

// HandleGetBlockchain Handles request to retrieve the entire blockchain
func (bc * Blockchain) HandleGetBlockchain(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    // Send the blockchain as the response
    response := struct {
        Blocks []*Block `json:"Blocks"`
    }{
        Blocks: bc.Blocks,
    }
    
    formattedJSON, err := json.MarshalIndent(response, "", "    ") // Indent with four spaces
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Error encoding JSON response: %s", err)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(formattedJSON)
}


// creeat the method that adds a new block to blocchain
func (blockchain *Blockchain) AddBlock(data string) bool {
	// Create a new block with the provided data
    prevBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
    newBlock := NewBlock(data, prevBlock.MyBlockHash, prevBlock.Difficulty)

    // Mine the new block
    if blockchain.mineBlock(newBlock) {
        blockchain.Blocks = append(blockchain.Blocks, newBlock)
        return true
    } else {
        return false
    }
}

func (blockchain *Blockchain) mineBlock(newBlock *Block) bool {
	// Continue with mining process
    targetPrefix := strings.Repeat("0", newBlock.Difficulty)
    for {
        newBlock.Nonce++
        hash := newBlock.calculateHash()
        if strings.HasPrefix(fmt.Sprintf("%x", hash), targetPrefix) {
            // Iterate through existing blocks in reverse order to check for duplicate data
            for i := len(blockchain.Blocks) - 1; i >= 0; i-- {
                block := blockchain.Blocks[i]
                if bytes.Equal(block.AllData, newBlock.AllData) {
                    fmt.Println("Block contains duplicate data, removing invalid block")
                    return false
                }
            }

            newBlock.MyBlockHash = hash
            return true // Successfully mined and added to the blockchain
        }
    }
}