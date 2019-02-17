package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/reinerRubin/sanic/internal"
)

func main() {
	path := os.Args[1]
	dataSet := os.Args[2]

	statBuffer := bytes.NewBuffer(nil)

	fmt.Fprint(statBuffer, `
package main

import	"github.com/reinerRubin/sanic/internal"
	`)

	mainStat := internal.NewStatByFile(path)

	testSize := 8
	testStat := make(internal.WordsByStat, testSize)
	for k, v := range mainStat {
		testSize--
		testStat[k] = v
		if testSize == 0 {
			break
		}
	}

	stat := mainStat
	if dataSet == "test" {
		stat = testStat
	}

	partSize := 10
	statParts := make([]internal.WordsByStat, partSize+1)
	currentPart := 0
	currentInPart := 0
	for k, v := range stat {
		st := statParts[currentPart]
		if len(st) == 0 {
			st = make(internal.WordsByStat, 0)
		}
		st[k] = v
		statParts[currentPart] = st

		currentInPart++
		if currentInPart >= len(stat)/partSize {
			currentPart++
			currentInPart = 0
		}
	}

	for i, part := range statParts {
		fmt.Fprintf(statBuffer, `func gStat%d() internal.WordsByStat {`, i)
		fmt.Fprint(statBuffer, "\nstat := internal.WordsByStat{}\n")
		for k, v := range part {
			fmt.Fprintf(statBuffer, "stat[%#v] = %#v\n", k, v)
		}

		fmt.Fprint(statBuffer, "\nreturn stat }\n")
	}


	if err := ioutil.WriteFile("dictionary.go", statBuffer.Bytes(), 0644); err != nil {
		panic(err)
	}
}
