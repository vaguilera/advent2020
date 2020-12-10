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

	sort.Ints(adapters)
	adapters = append(adapters, findMax(adapters)+3)
	fmt.Printf("adapters: %v\n", adapters)
	calculateShit()
}

func calculateShit() {
	count := 0
	quitables := []int{1}
	jumptable := []int{}
	finalnumber := 1

	for i := 0; i < len(adapters); i++ {
		if (i < len(adapters)-2) && ((adapters[i+2] - adapters[i]) <= 3) {
			quitables = append(quitables, adapters[i+1])
		}
		jumptable = append(jumptable, adapters[i]-count)
		count = adapters[i]

	}

	fmt.Printf("quitables: %v\n", quitables)
	fmt.Printf("jumptable: %v\n", jumptable)
	fmt.Printf("COUNT:%d\n", count)

	indexes := []int{}
	for i := 0; i < len(jumptable); i++ {

		if jumptable[i] == 1 {
			indexes = append(indexes, i)
		} else {

			switch len(indexes) {
			case 2:
				finalnumber *= 2
			case 3:
				finalnumber *= 4
			case 4:
				finalnumber *= 7
			default:

			}
			indexes = []int{}
		}
	}

	fmt.Printf("TEST:%d\n", finalnumber)
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
