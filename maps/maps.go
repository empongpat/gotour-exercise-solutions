package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		elem := m[v]
		m[v] = elem + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}