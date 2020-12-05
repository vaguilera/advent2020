package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	totalcount = 0

	for scanner.Scan() {
		line := scanner.Text()
		x, y := processSeat(line)
		val := (y * 8) + x
		if val > totalcount {
			totalcount = val
		}
	}

	//x, y := processSeat("BFFFBBFRRR")
	fmt.Printf("max: %d\n", totalcount)
}

func processSeat(s string) (x int, y int) {

	row := s[0:7]
	seat := s[7:10]

	rows := findPosition(row, 127)
	cols := findPosition(seat, 7)
	//fmt.Printf("X: %s, Y: %s - pos: %d - pos2: %d\n", row, seat, pos, pos2)

	return cols, rows
}

func findPosition(s string, max int) int {
	division := max
	pos := 0
	for i := 0; i < len(s); i++ {
		letter := s[i]
		division = (division / 2)
		if letter == 'B' || letter == 'R' {
			pos = (pos + division) + 1
		}
	}
	return pos
}
