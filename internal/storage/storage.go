package storage

import (
	"sync"
)

type Storage interface {
	Subscribe(address string) bool
	IsSubscribed(address string) bool
	GetSubscribers() []string
	AddTransaction(address string, transaction Transaction)
	GetTransactions(address string) []Transaction
	SetCurrentBlock(blockNumber int)
	GetCurrentBlock() int
}

type Transaction struct {
	From  string
	To    string
	Value string
	Hash  string
	Block int
}

type MemoryStorage struct {
	subscribers  map[string]bool
	transactions map[string][]Transaction
	currentBlock int
	mu           sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		subscribers:  make(map[string]bool),
		transactions: make(map[string][]Transaction),
		currentBlock: 0,
	}
}

func (s *MemoryStorage) Subscribe(address string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.subscribers[address] {
		return false
	}
	s.subscribers[address] = true
	return true
}

func (s *MemoryStorage) IsSubscribed(address string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.subscribers[address]
}

func (s *MemoryStorage) GetSubscribers() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	subscribers := make([]string, 0, len(s.subscribers))
	for address := range s.subscribers {
		subscribers = append(subscribers, address)
	}
	return subscribers
}

func (s *MemoryStorage) AddTransaction(address string, transaction Transaction) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.transactions[address] = append(s.transactions[address], transaction)
}

func (s *MemoryStorage) GetTransactions(address string) []Transaction {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.transactions[address]
}

func (s *MemoryStorage) SetCurrentBlock(blockNumber int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.currentBlock = blockNumber
}

func (s *MemoryStorage) GetCurrentBlock() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.currentBlock
}
