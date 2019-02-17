package internal

import (
	"bufio"
	"log"
	"os"
)

const CharsCnt = int(55)

type WordStat [CharsCnt]uint8

func NewWordStat(s string) *WordStat {
	st := new(WordStat)
	for _, r := range s {
		st.Register(r)
	}

	return st
}

func (s *WordStat) Register(r rune) {
	s[int(r)%CharsCnt]++
}

type WordsByStat map[WordStat][]int

func (wbs WordsByStat) Register(s string, wordNum int) {
	wordStat := NewWordStat(s)
	words, found := wbs[*wordStat]
	if !found {
		words = make([]int, 0)
	}
	words = append(words, wordNum)
	wbs[*wordStat] = words
}

func NewStatByFile(path string) WordsByStat {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stat := make(WordsByStat, 0)
	wordNum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		stat.Register(word, wordNum)
		wordNum++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return stat
}
