package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var totalcount int
var ids []int
var possibleIds []int

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
		ids = append(ids, val)
		if val > totalcount {
			totalcount = val
		}
	}
	fmt.Printf("max: %d\n", totalcount)
	findMySeat()
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

func findMySeat() {
	//seat := ids[0]

	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			val := (i * 8) + j
			if !Contains(ids, val) {
				fmt.Printf("No esta: x: %d, y: %d\n", j, i)
			}
		}
	}

	// id := ids[0]
	// lala, idx := Contains(ids, id+1)
	// fmt.Printf("N: %d N2:%d - %t\n", id, idx, lala)
	count := 0
	for _, n := range ids {
		if Contains(ids, n+1) && Contains(ids, n-1) {
			//fmt.Printf("N: %d\n", n)
			count++
		}

	}

	//fmt.Printf("total: %d", count)

}

func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
