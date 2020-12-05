package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
		if processLine(line) {
			totalcount++
		}
	}

	fmt.Printf("TOTAL: %d\n", totalcount)

}

func processLine(s string) bool {
	fields := strings.Split(s, " ")

	ret := false
	if (len(fields) == 8) || (len(fields) == 7 && !hasCID(fields)) {
		ret = true
	}

	fmt.Printf("%s ---- %d %t\n", s, len(fields), ret)
	return ret

}

func hasCID(fields []string) bool {

	for _, field := range fields {
		keys := strings.Split(field, ":")
		if keys[0] == "cid" {
			return true
		}
	}

	return false
}
