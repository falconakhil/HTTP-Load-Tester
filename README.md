# HTTP Load Tester

A command-line utility for HTTP load testing that allows you to benchmark the performance of web servers and applications. This tool supports testing individual URLs or multiple URLs from a CSV file with configurable concurrency and request parameters.

## Features

- Test single URLs with customizable request count and concurrency
- Process multiple URLs from a CSV file
- Configurable concurrency levels for both URL and file-based testing
- Comprehensive performance metrics including response time statistics
- Detailed logging to file

## Installation

### Prerequisites

- Go 1.13 or higher

### Building from source

```bash
# Clone the repository
git clone https://github.com/yourusername/http-load-tester.git
cd http-load-tester

# Build the binary
go build -o loadtest
```

## Usage

### Testing a single URL

```bash
# Basic usage
./loadtest test https://example.com

# With custom parameters
./loadtest test https://example.com --requests 100 --concurrency 10 --logfile test.log
```

### Testing multiple URLs from a CSV file

```bash
# Basic usage
./loadtest file urls.csv

# With custom parameters
./loadtest file urls.csv --concurrent_urls 5 --logfile batch_test.log
```

## Command Line Options

### Global Options

- `--logfile` - Path to log file (default: stdout)

### `test` Command Options

- `--requests, -n` - Number of requests to send (default: 1)
- `--concurrency, -c` - Number of concurrent requests (default: 1)

### `file` Command Options

- `--concurrent_urls, -c` - Number of concurrent URL tests (default: 1)

## CSV File Format

The CSV file for batch testing should have the following format:

```
URL,REQUESTS,CONCURRENCY
https://example.com,100,10
https://another-example.com,50,5
```

Each line contains:

1. URL to test
2. Number of requests to make
3. Concurrency level for that URL

## Project Structure

```
HTTP Load Tester/
├── cmd/            # Command definitions and CLI handlers
│   ├── file.go     # Command for file-based testing
│   ├── root.go     # Root command definition
│   └── test.go     # Command for single URL testing
├── lib/            # Core library code
│   ├── tester.go   # Main testing functionality
│   └── testUrlWorker.go # Worker implementation for URL testing
├── models/         # Data models and display functions
└── main.go         # Application entry point
```

## Output Metrics

The tool provides the following performance metrics:

- Total Time: Overall time taken for all requests
- Average Response Time: Mean response time across all requests
- Min/Max Response Time: Fastest and slowest responses
- Success Rate: Percentage of successful responses
- Status Code Distribution: Count of responses by HTTP status code
