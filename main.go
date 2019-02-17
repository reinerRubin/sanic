package main

//go:generate go run cmd/genarray/main.go lemmad/lemmad.txt !test
//go:generate go run cmd/genwords/main.go lemmad/lemmad.txt

import (
	"fmt"
	"os"
	"time"

	"github.com/reinerRubin/sanic/internal"
)

func main() {
	started := time.Now()

	_ = os.Args[1]
	word := os.Args[2]

	wordPointerParts := [...][]int{
		gStat0()[*internal.NewWordStat(word)],
		gStat1()[*internal.NewWordStat(word)],
		gStat2()[*internal.NewWordStat(word)],
		gStat3()[*internal.NewWordStat(word)],
		gStat4()[*internal.NewWordStat(word)],
		gStat5()[*internal.NewWordStat(word)],
		gStat6()[*internal.NewWordStat(word)],
		gStat7()[*internal.NewWordStat(word)],
		gStat8()[*internal.NewWordStat(word)],
		gStat9()[*internal.NewWordStat(word)],
		gStat10()[*internal.NewWordStat(word)],
	}

	var result = ""
	words := gWords()
	for _, wordPointers := range wordPointerParts {
		for i := 0; i < len(wordPointers); i++ {
			result += words[wordPointers[i]]
			if i != len(wordPointers)-1 {
				result += ", "
			}
		}
	}

	duration := time.Since(started)
	fmt.Printf("%d, %s\n", duration.Nanoseconds()/1000, result)
}
