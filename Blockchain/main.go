// main File
package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// Init Blockxhain
	newblockchain := NewBlockchain()

	// Define API routea
	router.HandleFunc("/addData", newblockchain.HandleAddData).Methods("POST") // curl -X POST -H "Content-Type: application/json" -d "{\"data\":\"first\"}" http://localhost:8080/addData
	router.HandleFunc("/getBlockchain", newblockchain.HandleGetBlockchain).Methods("GET") // curl http://localhost:8080/getBlockchain

	http.Handle("/", router)

	// curl http://localhost:8080/print
	http.HandleFunc("/print", func(w http.ResponseWriter, r *http.Request) {
		for i, block :=  range newblockchain.Blocks {
			fmt.Printf("\nBlock ID: %d\n", i)
			fmt.Printf("Timestamp: %d\n", block.Timestamp+int64(i))
			fmt.Printf("Hash of the block: %x\n", block.MyBlockHash)
			fmt.Printf("Hash of the previous Block: %x\n", block.PreviousBlockHash)
			fmt.Printf("Data: %s\n\n", block.AllData)
		}
	})

	// Health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	//Start the http server
	fmt.Println("Listening on port :8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}