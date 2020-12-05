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
	min, _ := strconv.Atoi(splited[1])
	max, _ := strconv.Atoi(splited[2])
	letter := splited[3]
	text := splited[4]

	times := countTimes(letter[0], text)
	if times >= min && times <= max {
		totalcount++
	}

	//Printf("%d - %d - %s - %d -- %s\n", min, max, letter, times, text)
}

func countTimes(letter byte, text string) int {
	times := 0
	for i := 0; i < len(text); i++ {
		if text[i] == letter {
			times++
		}
	}

	return times
}
