package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var totalcount int
var adapters []int

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
		adapters = append(adapters, num)
	}

	adapters = append(adapters, findMax(adapters)+3)
	calculateShit()
}

func calculateShit() {
	joltage := 0
	sort.Ints(adapters)

	jmp1, jmp3 := 0, 0

	for _, val := range adapters {
		diff := val - joltage
		if diff == 1 {
			jmp1++
		}
		if diff == 3 {
			jmp3++
		}
		joltage = val
	}

	fmt.Printf("1:%d 3:%d. Result: %d\n", jmp1, jmp3, jmp1*jmp3)

}

func findMax(v []int) int {
	max := 0

	for _, val := range v {
		if val > max {
			max = val
		}
	}

	return max
}
