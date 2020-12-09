package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

const MAGIC = 1930745883

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
		buffer = append(buffer, num)

	}

	mybuf := findNumber()
	fmt.Printf("result: %d\n", calculateValue(mybuf))

}

func findNumber() []int {
	tbuf := make([]int, 0)
	i, j := 0, 0

	for {
		tbuf = append(tbuf, buffer[i+j])
		val := sumSlice(tbuf)

		if val == MAGIC {
			fmt.Printf("Encontrado: %v\n", tbuf)
			return tbuf
		}
		if val > MAGIC {
			fmt.Println("overflow")
			tbuf = []int{}
			i++
			j = 0
		} else {
			j++
		}
	}
}

func sumSlice(buffer []int) int {
	res := 0
	for i := 0; i < len(buffer); i++ {
		res += buffer[i]
	}
	return res
}

func calculateValue(buffer []int) int {
	max := 0
	min := math.MaxInt32

	for _, num := range buffer {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	fmt.Printf("min: %d max:%d\n", min, max)
	return min + max

}
