package parser

import (
	"eth-parser/internal/storage"
	"eth-parser/pkg/ethereum"
	"strings"
)

type Parser interface {
	GetCurrentBlock() int
	Subscribe(address string) bool
	GetTransactions(address string) []storage.Transaction
	ProcessBlock(blockNumber int) error
}

type EthereumParser struct {
	storage storage.Storage
	client  ethereum.EthereumClient
}

func NewEthereumParser(s storage.Storage, c ethereum.EthereumClient) *EthereumParser {
	if s == nil {
		s = storage.NewMemoryStorage()
	}
	if c == nil {
		c = &ethereum.DefaultEthereumClient{}
	}
	return &EthereumParser{
		storage: s,
		client:  c,
	}
}

func (p *EthereumParser) GetCurrentBlock() int {
	return p.storage.GetCurrentBlock()
}

func (p *EthereumParser) Subscribe(address string) bool {
	return p.storage.Subscribe(strings.ToLower(address))
}

func (p *EthereumParser) GetTransactions(address string) []storage.Transaction {
	return p.storage.GetTransactions(strings.ToLower(address))
}

func (p *EthereumParser) ProcessBlock(blockNumber int) error {
	block, err := p.client.GetBlockByNumber(blockNumber)
	if err != nil {
		return err
	}

	for _, tx := range block.Transactions {
		p.processTransaction(tx, blockNumber)
	}

	p.storage.SetCurrentBlock(blockNumber)
	return nil
}

func (p *EthereumParser) processTransaction(tx ethereum.Transaction, blockNumber int) {
	transaction := storage.Transaction{
		From:  tx.From,
		To:    tx.To,
		Value: tx.Value,
		Hash:  tx.Hash,
		Block: blockNumber,
	}

	if p.storage.IsSubscribed(strings.ToLower(tx.From)) {
		p.storage.AddTransaction(strings.ToLower(tx.From), transaction)
	}

	if p.storage.IsSubscribed(strings.ToLower(tx.To)) {
		p.storage.AddTransaction(strings.ToLower(tx.To), transaction)
	}
}
