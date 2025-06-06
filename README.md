# MapReduce in Go

This project is a demonstration of a simple MapReduce framework implemented in Go with a sample text file for educational purposes.

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
go run main.go sample.txt
```
