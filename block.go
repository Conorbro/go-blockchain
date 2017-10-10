package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

const randStringLength = 25

// Block - block in the blockchain
type Block struct {
	blockHash    string
	previousHash string
	transactions []string
}

// Blockchain - slice of block references
type Blockchain struct {
	blocks []*Block
}

func (b *Blockchain) addBlock(transactions []string) {
	var blockHash, previousHash string
	if len(b.blocks) == 0 {
		previousHash = "0000000000000000000000000000000000000000000000000000000000000000"
		blockHash = genSHA(strings.Join(transactions, ""))
	} else {
		previousHash = b.blocks[len(b.blocks)-1].blockHash
		blockHash = generateNextHash(*b.blocks[len(b.blocks)-1])
	}
	b.blocks = append(b.blocks,
		&Block{
			blockHash:    blockHash,
			previousHash: previousHash,
			transactions: transactions,
		})
}

func generateNextHash(b Block) string {
	for {
		answer := randStringBytesMaskImprSrc(randStringLength)
		attempt := b.previousHash + answer
		solution := genSHA(attempt)
		if strings.HasPrefix(solution, *difficulty) {
			fmt.Printf("Solutin found!\nAdding new block....\n")
			return solution
		}
	}
}

func genSHA(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	res := fmt.Sprintf("%x", h.Sum(nil))
	return res
}

func (b *Blockchain) printChain() {
	for _, block := range b.blocks {
		fmt.Println(block)
	}
}
