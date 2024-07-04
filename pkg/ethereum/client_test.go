package ethereum

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDefaultEthereumClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var response JSONRPCResponse
		if r.Method == "POST" {
			var request JSONRPCRequest
			json.NewDecoder(r.Body).Decode(&request)
			switch request.Method {
			case "eth_blockNumber":
				response = JSONRPCResponse{
					JSONRPC: "2.0",
					Result:  "0x4b7", // 1207 in decimal
					ID:      1,
				}
			case "eth_getBlockByNumber":
				response = JSONRPCResponse{
					JSONRPC: "2.0",
					Result: map[string]interface{}{
						"number": "0x1",
						"transactions": []map[string]interface{}{
							{
								"from":  "0x123",
								"to":    "0x456",
								"value": "0x64",
								"hash":  "0xabc",
							},
						},
					},
					ID: 1,
				}
			}
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	originalEndpoint := ethereumRPCEndpoint
	ethereumRPCEndpoint = server.URL
	defer func() { ethereumRPCEndpoint = originalEndpoint }()

	client := &DefaultEthereumClient{}

	t.Run("GetLatestBlockNumber", func(t *testing.T) {
		blockNumber, err := client.GetLatestBlockNumber()
		if err != nil {
			t.Errorf("GetLatestBlockNumber failed: %v", err)
		}
		if blockNumber != 1207 {
			t.Errorf("Expected block number 1207, got %d", blockNumber)
		}
	})

	t.Run("GetBlockByNumber", func(t *testing.T) {
		block, err := client.GetBlockByNumber(1)
		if err != nil {
			t.Errorf("GetBlockByNumber failed: %v", err)
		}
		if block.Number != "0x1" {
			t.Errorf("Expected block number 0x1, got %s", block.Number)
		}
		if len(block.Transactions) != 1 {
			t.Errorf("Expected 1 transaction, got %d", len(block.Transactions))
		}
	})
}
