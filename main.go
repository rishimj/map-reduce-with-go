package main

import (
    "bufio"
    "fmt"
    "log"
    "mapreduce/core"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: mapreduce <input_file>")
        os.Exit(1)
    }
    inputFile := os.Args[1]
    file, err := os.Open(inputFile)
    if err != nil {
        log.Fatalf("Failed to open file: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatalf("Error reading file: %v", err)
    }

    // Define map function: split line into words
    mapF := func(line string) []core.KeyValue {
        words := core.SplitWords(line)
        kva := []core.KeyValue{}
        for _, w := range words {
            kva = append(kva, core.KeyValue{Key: w, Value: 1})
        }
        return kva
    }

    // Define reduce function: sum counts
    reduceF := func(key string, values []int) core.KeyValue {
        sum := 0
        for _, v := range values {
            sum += v
        }
        return core.KeyValue{Key: key, Value: sum}
    }

    mr := core.NewMapReduce(mapF, reduceF)
    results := mr.Execute(lines)

    for _, kv := range results {
        fmt.Printf("%s: %d\n", kv.Key, kv.Value)
    }
}
