package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const PREVSIZE = 25

var totalcount int
var buffer []int

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
		num, _ := strconv.Atoi(line)

		if len(buffer) < PREVSIZE {
			buffer = append(buffer, num)
			continue
		}

		if !checkNumber(num) {
			fmt.Printf("NOT FOUND SOLUTION: %d\n", num)
			os.Exit(0)
		}
		fmt.Printf("FOUND SOLUTION: %d\n", num)

		buffer = append(buffer[1:], num)
	}
}

func checkNumber(num int) bool {

	for i := 0; i < len(buffer); i++ {
		for j := 0; j < len(buffer); j++ {
			if i == j {
				continue
			}
			if buffer[i]+buffer[j] == num {
				return true
			}
		}
	}

	return false
}
