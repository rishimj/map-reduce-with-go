package core

// KeyValue type holds a key-value pair
type KeyValue struct {
    Key   string
    Value int
}

// MapFunc defines the signature of a map function
type MapFunc func(input string) []KeyValue

// ReduceFunc defines the signature of a reduce function
type ReduceFunc func(key string, values []int) KeyValue
