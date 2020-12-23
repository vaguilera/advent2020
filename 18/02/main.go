package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var totalcount int
var r, r2 *regexp.Regexp

func main() {

	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalcount = 0
	r, _ = regexp.Compile(`\(([^()]*)\)`)
	r2, _ = regexp.Compile(`((\d+)(\D?))`)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, " ", "")
		proc := processLine(line)
		res := calculate(proc)
		fmt.Printf("Result: %d\n", res)
		totalcount += res
	}

	fmt.Printf("Total: %d\n", totalcount)
}

func calculate(s string) int {
	tokens := strings.Split(s, "*")

	accum := 1
	for _, token := range tokens {
		ops := strings.Split(token, "+")
		res := 0
		for _, n := range ops {
			ni, _ := strconv.Atoi(n)
			res += ni
		}

		accum *= res
	}

	return accum

}

func processLine(s string) string {

	ops := r.FindAllStringSubmatch(s, -1)

	if len(ops) == 0 {
		return s
	}

	for _, op := range ops {
		res := calculate(op[1])
		sres := fmt.Sprintf("%d", res)
		s = strings.ReplaceAll(s, op[0], sres)
	}

	return processLine(s)

}
