// Structure.go file
package main

type Block struct {
	Timestamp			int64 // the time when the block was created
	PreviousBlockHash 	[]byte //the hash of the previous block
	MyBlockHash 		[]byte // the hash of the current block
	AllData 			[]byte // the data or transscation (bosy info)
	Nonce				int
	Difficulty			int
}

type Blockchain struct {
	Blocks []*Block
}