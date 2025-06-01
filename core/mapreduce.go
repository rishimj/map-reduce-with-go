package core

import (
    "sort"
    "sync"
)

// MapReduce struct holds map and reduce functions
type MapReduce struct {
    MapFunc    MapFunc
    ReduceFunc ReduceFunc
}

// NewMapReduce creates a new MapReduce instance
func NewMapReduce(m MapFunc, r ReduceFunc) *MapReduce {
    return &MapReduce{MapFunc: m, ReduceFunc: r}
}

// Execute runs the MapReduce algorithm concurrently: map phase, shuffle, and reduce phase.
func (mr *MapReduce) Execute(inputs []string) []KeyValue {
    // 1. Map phase (concurrent)
    intermediateChan := make(chan KeyValue)
    var mapWg sync.WaitGroup
    for _, input := range inputs {
        mapWg.Add(1)
        go func(in string) {
            defer mapWg.Done()
            kva := mr.MapFunc(in)
            for _, kv := range kva {
                intermediateChan <- kv
            }
        }(input)
    }

    // Close channel when done
    go func() {
        mapWg.Wait()
        close(intermediateChan)
    }()

    // 2. Shuffle phase: group values by key
    grouped := make(map[string][]int)
    for kv := range intermediateChan {
        grouped[kv.Key] = append(grouped[kv.Key], kv.Value)
    }

    // 3. Sort keys
    keys := make([]string, 0, len(grouped))
    for k := range grouped {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    // 4. Reduce phase (concurrent)
    var reduceWg sync.WaitGroup
    resultChan := make(chan KeyValue)

    for _, k := range keys {
        reduceWg.Add(1)
        go func(key string, values []int) {
            defer reduceWg.Done()
            result := mr.ReduceFunc(key, values)
            resultChan <- result
        }(k, grouped[k])
    }

    // Close results channel when done
    go func() {
        reduceWg.Wait()
        close(resultChan)
    }()

    // 5. Collect results
    var results []KeyValue
    for res := range resultChan {
        results = append(results, res)
    }

    // Sort results by Key to ensure deterministic order
    sort.Slice(results, func(i, j int) bool {
        return results[i].Key < results[j].Key
    })

    return results
}
