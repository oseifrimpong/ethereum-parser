# Ethereum Blockchain Parser

This application implements an Ethereum blockchain parser that allows querying transactions for subscribed addresses.

## Getting Started

### Prerequisites

- Go 1.15 or higher

### Installation

1. Clone the repository:
   `git clone https://github.com/oseifrimpong/ethereum-parser.git`
2. cd ethereum-blockchain-parser
3. Install dependencies (if any): `go mod tidy`

### Running the Application

To run the application, use the following command from the root directory of the project:
`go run main.go`

This will start the Ethereum parser. You can interact with it through the command-line interface.

### Running Tests

To run the tests, use the following command: `go test ./...`

This will run all tests in all packages of the project.

## Usage

The application provides a command-line interface with the following commands:

- `subscribe <address>`: Subscribe to an Ethereum address
- `transactions <address>`: Get all transactions for a subscribed address
- `current`: Get the current processed block number
- `exit`: Quit the program

## Approach and Implementation

The solution is implemented using the following approach:

1. **Modular Design**: The application is divided into several packages:
   - `ethereum`: Handles interaction with the Ethereum blockchain
   - `parser`: Implements the main parsing logic
   - `storage`: Manages data storage

2. **Interfaces**: Key components are defined as interfaces to allow for easy mocking in tests and future extensibility:
   - `Parser`: Defines the main parsing operations
   - `Storage`: Abstracts the data storage operations
   - `EthereumClient`: Abstracts the Ethereum blockchain interaction

3. **In-Memory Storage**: The current implementation uses in-memory storage for simplicity, but the design allows for easy extension to other storage types in the future.

4. **Ethereum JSON-RPC**: The application interacts with the Ethereum blockchain using the JSON-RPC protocol, specifically the `eth_blockNumber` and `eth_getBlockByNumber` methods.

5. **Concurrent Processing**: The application processes blocks in the background, allowing for real-time updates and interaction through the command-line interface.

6. **Error Handling**: Comprehensive error handling is implemented throughout the application to ensure robustness.

7. **Testing**: Unit tests are provided for key components, with mocking of external dependencies where appropriate.

## Future Improvements

- Implement persistent storage
- Add more comprehensive error logging
- Implement rate limiting for Ethereum RPC calls
- Add support for websocket subscriptions for real-time updates
- Implement a REST API for interacting with the parser
