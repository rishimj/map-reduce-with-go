# MapReduce in Go

This project is a demonstration of a simple MapReduce framework implemented in Go, using best practices, concurrency with goroutines and channels, and a word count example.

## Project Structure

- `go.mod`: Module definition.
- `main.go`: Example usage reading a text file and counting word frequency.
- `core/`: Core MapReduce logic and types.
  - `types.go`: Defines KeyValue, MapFunc, ReduceFunc.
  - `mapreduce.go`: Implements concurrent MapReduce.
  - `utils.go`: Utility functions, e.g., SplitWords.
  - `mapreduce_test.go`: Unit tests.
- `README.md`: This file.

## How to Run

```bash
go mod tidy
go test ./core
go run main.go <path_to_input_file>
```
