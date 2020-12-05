package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var totalcount int

func main() {

	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var line string

	totalcount = 0
	r, _ := regexp.Compile(`(\d*)-(\d*)\s(.):\s(.*)`)
	for scanner.Scan() {
		line = scanner.Text()
		processLine(line, r)
	}

	fmt.Printf("TOTAL: %d\n", totalcount)
}

func processLine(line string, r *regexp.Regexp) {
	splited := r.FindStringSubmatch(line)
	pos1, _ := strconv.Atoi(splited[1])
	pos2, _ := strconv.Atoi(splited[2])
	letter := splited[3]
	text := splited[4]

	letterPos1 := text[pos1-1]
	letterPos2 := text[pos2-1]

	count := 0
	if letterPos1 == letter[0] {
		count++
	}
	if letterPos2 == letter[0] {
		count++
	}

	if count == 1 {
		totalcount++
	}

	fmt.Printf("%d - %d - %d - %d -- %s -- %d - %d\n", letterPos1, letterPos2, letter[0], count, text, pos1, pos2)
}
