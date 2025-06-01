package core

import (
    "regexp"
    "strings"
)

// SplitWords splits a string into lowercase words using regex.
func SplitWords(text string) []string {
    // use regex to find word characters
    re := regexp.MustCompile(`[a-zA-Z]+`)
    rawWords := re.FindAllString(text, -1)
    var words []string
    for _, w := range rawWords {
        lower := strings.ToLower(w)
        words = append(words, lower)
    }
    return words
}
