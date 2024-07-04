package parser

import (
	"eth-parser/pkg/ethereum"
	"testing"

	"eth-parser/internal/storage"
)

type MockEthereumClient struct{}

func (m *MockEthereumClient) GetLatestBlockNumber() (int, error) {
	return 100, nil
}

func (m *MockEthereumClient) GetBlockByNumber(blockNumber int) (*ethereum.Block, error) {
	return &ethereum.Block{
		Number: "0x1",
		Transactions: []ethereum.Transaction{
			{From: "0x123", To: "0x456", Value: "100", Hash: "0xabc"},
		},
	}, nil
}

func TestEthereumParser(t *testing.T) {
	s := storage.NewMemoryStorage()
	c := &MockEthereumClient{}
	p := NewEthereumParser(s, c)

	// Test Subscribe
	if !p.Subscribe("0x123") {
		t.Error("Failed to subscribe new address")
	}

	// Test GetCurrentBlock
	s.SetCurrentBlock(10)
	if p.GetCurrentBlock() != 10 {
		t.Error("GetCurrentBlock failed")
	}

	// Test ProcessBlock
	err := p.ProcessBlock(1)
	if err != nil {
		t.Errorf("ProcessBlock failed: %v", err)
	}

	// Test GetTransactions
	transactions := p.GetTransactions("0x123")
	if len(transactions) != 1 {
		t.Error("GetTransactions failed")
	}
}
