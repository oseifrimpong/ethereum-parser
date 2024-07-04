package ethereum

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var ethereumRPCEndpoint = "https://cloudflare-eth.com"

type EthereumClient interface {
	GetLatestBlockNumber() (int, error)
	GetBlockByNumber(blockNumber int) (*Block, error)
}

type DefaultEthereumClient struct{}

func getEthereumRPCEndpoint() string {
	return ethereumRPCEndpoint
}
func SetEthereumRPCEndpoint(url string) {
	ethereumRPCEndpoint = url
}

func (c *DefaultEthereumClient) GetLatestBlockNumber() (int, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_blockNumber",
		Params:  []interface{}{},
		ID:      1,
	}

	response, err := sendRequest(request)
	if err != nil {
		return 0, err
	}

	blockNumberHex, ok := response.Result.(string)
	if !ok {
		return 0, fmt.Errorf("unexpected result type: %T", response.Result)
	}

	return hexToInt(blockNumberHex)
}

func (c *DefaultEthereumClient) GetBlockByNumber(blockNumber int) (*Block, error) {
	hexBlockNumber := fmt.Sprintf("0x%x", blockNumber)
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  []interface{}{hexBlockNumber, true},
		ID:      1,
	}

	response, err := sendRequest(request)
	if err != nil {
		return nil, err
	}

	var block Block
	resultBytes, err := json.Marshal(response.Result)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resultBytes, &block)
	if err != nil {
		return nil, err
	}

	return &block, nil
}

func sendRequest(request JSONRPCRequest) (*JSONRPCResponse, error) {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(getEthereumRPCEndpoint(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response JSONRPCResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func hexToInt(hex string) (int, error) {
	// Ensure the hex string has the "0x" prefix
	hex = strings.TrimPrefix(hex, "0x")
	result, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return 0, err
	}
	return int(result), nil
}
