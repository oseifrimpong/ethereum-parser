package storage

import (
	"testing"
)

func TestMemoryStorage(t *testing.T) {
	s := NewMemoryStorage()

	// Test Subscribe
	if !s.Subscribe("0x123") {
		t.Error("Failed to subscribe new address")
	}
	if s.Subscribe("0x123") {
		t.Error("Subscribed to already subscribed address")
	}

	// Test IsSubscribed
	if !s.IsSubscribed("0x123") {
		t.Error("IsSubscribed failed for subscribed address")
	}
	if s.IsSubscribed("0x456") {
		t.Error("IsSubscribed returned true for unsubscribed address")
	}

	// Test GetSubscribers
	subscribers := s.GetSubscribers()
	if len(subscribers) != 1 || subscribers[0] != "0x123" {
		t.Error("GetSubscribers failed")
	}

	// Test AddTransaction and GetTransactions
	tx := Transaction{From: "0x123", To: "0x456", Value: "100", Hash: "0xabc", Block: 1}
	s.AddTransaction("0x123", tx)
	transactions := s.GetTransactions("0x123")
	if len(transactions) != 1 || transactions[0] != tx {
		t.Error("AddTransaction or GetTransactions failed")
	}

	// Test SetCurrentBlock and GetCurrentBlock
	s.SetCurrentBlock(100)
	if s.GetCurrentBlock() != 100 {
		t.Error("SetCurrentBlock or GetCurrentBlock failed")
	}
}
