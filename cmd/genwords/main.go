package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	path := os.Args[1]

	buffer := bytes.NewBuffer(nil)


	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	words := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}


	fmt.Fprintf(buffer, `package main

func gWords() [%d]string { var words [%d]string; 
	`, len(words), len(words))

	for i, word := range words {
		fmt.Fprintf(buffer, "words[%d] = \"%s\"\n", i, word)
	}


	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(buffer, "return words }")

	if err := ioutil.WriteFile("words.go", buffer.Bytes(), 0644); err != nil {
		panic(err)
	}
}
