package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wc := make(map[string]int)
	wl := len(words)
	
	for i := 0; i < wl; i++ {
		if _, ok := wc[words[i]]; ok == true {
			wc[words[i]] += 1
		} else {
			wc[words[i]] = 1
		}
	}
	
	return wc
}

func main() {
	wc.Test(WordCount)
}
