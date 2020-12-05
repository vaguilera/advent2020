package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var totalcount int
var forest []string

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
		forest = append(forest, line)
	}

	fmt.Printf("WIDTH: %d\n", len(forest[0]))
	fmt.Printf("ARBOLES1: %d\n", start(1, 1))
	fmt.Printf("ARBOLES2: %d\n", start(3, 1))
	fmt.Printf("ARBOLES3: %d\n", start(5, 1))
	fmt.Printf("ARBOLES4: %d\n", start(7, 1))
	fmt.Printf("ARBOLES5: %d\n", start(1, 2))

	fmt.Printf("MULTIPLY: %d\n", start(1, 1)*start(3, 1)*start(5, 1)*start(7, 1)*start(1, 2))
}

func start(slopeX int, slopeY int) int {
	x, y, arboles := 0, 0, 0

	for y = slopeY; y < len(forest); y += slopeY {
		x = (x + slopeX) % 31

		if forest[y][x] == 35 {
			arboles++
		}
	}

	return arboles
}
