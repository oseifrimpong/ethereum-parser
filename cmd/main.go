package main

import (
	"bufio"
	"eth-parser/internal/parser"
	"eth-parser/internal/storage"
	"eth-parser/pkg/ethereum"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	memoryStorage := storage.NewMemoryStorage()
	ethereum := &ethereum.DefaultEthereumClient{}
	p := parser.NewEthereumParser(memoryStorage, ethereum)

	// Simple command-line interface
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")

		switch parts[0] {
		case "start":
			go func() {
				for {
					currentBlock, err := ethereum.GetLatestBlockNumber()
					if err != nil {
						fmt.Println("Error getting latest block:", err)
						time.Sleep(15 * time.Second)
						continue
					}

					for i := p.GetCurrentBlock() + 1; i <= currentBlock; i++ {
						err := p.ProcessBlock(i)
						if err != nil {
							fmt.Printf("Error processing block %d: %v\n", i, err)
							continue
						}
						fmt.Printf("Processed block %d\n", i)
					}

					time.Sleep(15 * time.Second)
				}
			}()
			fmt.Println("Block processing started")

		case "subscribe":
			if len(parts) != 2 {
				fmt.Println("Usage: subscribe <address>")
				continue
			}
			success := p.Subscribe(parts[1])
			if success {
				fmt.Println("Subscribed successfully")
			} else {
				fmt.Println("Address already subscribed")
			}

		case "transactions":
			if len(parts) != 2 {
				fmt.Println("Usage: transactions <address>")
				continue
			}
			transactions := p.GetTransactions(parts[1])
			for _, tx := range transactions {
				fmt.Printf("From: %s, To: %s, Value: %s, Hash: %s, Block: %d\n", tx.From, tx.To, tx.Value, tx.Hash, tx.Block)
			}

		case "current":
			fmt.Printf("Current block: %d\n", p.GetCurrentBlock())

		case "exit":
			return

		default:
			fmt.Println("Unknown command. Available commands: start, subscribe, transactions, current, exit")
		}
	}
}
