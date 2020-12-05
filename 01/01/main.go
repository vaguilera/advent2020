package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var numbers []int

func main() {

	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	for _, strnum := range text {
		i, _ := strconv.Atoi(strnum)
		numbers = append(numbers, i)
	}

	fmt.Println(len(numbers))
	findNumbers()
}

func findNumbers() {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == 2020 {
				log.Printf("%d - %d - %d - %d - %d", i, j, numbers[i], numbers[j], numbers[i]*numbers[j])
			}
		}

	}
}
