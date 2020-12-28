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

type rule struct {
	final   string
	branchA []int
	branchB []int
}

var rulez map[int]rule

func main() {

	file, err := os.Open("rulez.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	rulez = make(map[int]rule)
	r, _ = regexp.Compile(`\(([^()]*)\)`)
	r2, _ = regexp.Compile(`((\d+)(\D?))`)
	for scanner.Scan() {
		line := scanner.Text()
		parseRule(line)
	}

	file, err = os.Open("input.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalcount = 0
	for scanner.Scan() {
		line := scanner.Text()

		s1, e := checkrule(0, line)
		if s1 == "" && e == 0 {
			totalcount++
		}
	}
	fmt.Printf("RES: %d\n", totalcount)

}

func parseRule(s string) {
	tokens := strings.Split(s, ": ")
	nrul, _ := strconv.Atoi(tokens[0])

	if tokens[1] == "\"a\"" || tokens[1] == "\"b\"" {
		rulez[nrul] = rule{
			final: tokens[1][1:2],
		}
		return
	}

	tokens = strings.Split(tokens[1], " | ")
	rulesA := []int{}
	rulesB := []int{}

	values := strings.Split(tokens[0], " ")
	for _, val := range values {
		n, _ := strconv.Atoi(val)
		rulesA = append(rulesA, n)
	}

	if len(tokens) > 1 {
		values = strings.Split(tokens[1], " ")
		for _, val := range values {
			n, _ := strconv.Atoi(val)
			rulesB = append(rulesB, n)
		}
	}

	rulez[nrul] = rule{
		final:   "",
		branchA: rulesA,
		branchB: rulesB,
	}
}

func checkrule(r int, s string) (string, int) {
	if len(s) == 0 {
		return "", -1
	}
	if rulez[r].final != "" {
		if s[0:1] == rulez[r].final {
			return s[1:], 0
		}
		return "", -1
	}

	res := s
	err := 0
	for _, n := range rulez[r].branchA {
		res, err = checkrule(n, res)
		if err == -1 {
			break
		}
	}

	if err == 0 {
		return res, 0
	}

	if len(rulez[r].branchB) == 0 {
		return "", -1
	}

	res = s
	err = 0
	for _, n := range rulez[r].branchB {
		res, err = checkrule(n, res)
		if err == -1 {
			break
		}
	}

	if err == 0 {
		return res, 0
	}

	return "", -1

}
