package core

import (
    "reflect"
    "testing"
)

func TestMapReduceWordCount(t *testing.T) {
    inputs := []string{
        "Hello world",
        "world of Go",
    }

    mapF := func(line string) []KeyValue {
        words := SplitWords(line)
        kva := []KeyValue{}
        for _, w := range words {
            kva = append(kva, KeyValue{Key: w, Value: 1})
        }
        return kva
    }

    reduceF := func(key string, values []int) KeyValue {
        sum := 0
        for _, v := range values {
            sum += v
        }
        return KeyValue{Key: key, Value: sum}
    }

    mr := NewMapReduce(mapF, reduceF)
    got := mr.Execute(inputs)

    want := []KeyValue{
        {"go", 1},
        {"hello", 1},
        {"of", 1},
        {"world", 2},
    }

    if !reflect.DeepEqual(got, want) {
        t.Errorf("Got %v, want %v", got, want)
    }
}
