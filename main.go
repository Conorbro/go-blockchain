package main

import (
	"fmt"
	"time"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	difficulty = kingpin.Flag("difficulty", "Number of proceeding 0's in answer").Default("000").String()
)

func main() {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()
	fmt.Println("Simple Blockchain Demo!")

	bChain := blockChain{}
	startTime := time.Now()
	bChain.addBlock([]string{"Conor sent Dan 20 bitcoin", "Gerry sent Bob 0.01 bitcoin", "Fred sent Mary 243 bitcoin"})
	bChain.addBlock([]string{"Conor sent Jamie 1 bitcoin", "Barry sent Sean 0.01 bitcoin", "Shane sent Mary 0.29 bitcoin"})

	fmt.Println("Blockchain updated!")
	fmt.Printf("Time to add complete chain: %v\n", time.Since(startTime))
}
