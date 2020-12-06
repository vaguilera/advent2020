package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var totalcount int
var current []byte
var totals []int
var numlinea = 0

func main() {

	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalcount = 0
	numlinea = 1
	for scanner.Scan() {
		line := scanner.Text()
		processLine(line)
		numlinea++
	}
	//fmt.Printf("current: %s\n", string(current))

	//fmt.Printf("llala: %v\n", totals)

	for i := 0; i < len(totals); i++ {
		totalcount += totals[i]
	}
	fmt.Printf("llala: %d\n", totalcount)

}

func processLine(s string) {
	s2 := []byte(s)
	s3 := []byte{}

	if s == "" {
		totals = append(totals, len(current))
		current = []byte{}
		numlinea = 0
		return
	}

	if numlinea == 1 {
		current = s2
		return
	}

	for i := 0; i < len(s2); i++ {
		_, f := Find(current, s2[i])
		if f {
			s3 = append(s3, s[i])
		}
	}
	current = s3
}

func Find(slice []byte, val byte) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
