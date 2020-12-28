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

		if checkrule2(line) {
			totalcount++
		}
	}

	fmt.Printf("RES TOTAL: %d\n", totalcount)

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
	//fmt.Printf("checking: %d, s: %s\n", r, s)
	if len(s) == 0 {
		return s, -1
	}
	if rulez[r].final != "" {
		if s[0:1] == rulez[r].final {
			return s[1:], 0
		}
		return s, -1
	}

	res := s
	err := -1
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
		return s, -1
	}

	res = s
	err = -1
	for _, n := range rulez[r].branchB {
		res, err = checkrule(n, res)
		if err == -1 {
			return s, -1
		}
	}

	return res, 0

}

func checkrule2(s string) bool {
	/*
		1: 42, 42, 31
		2: 42, 42 * x, 31 * x
		3: 42 * x, 42, 31
		4: 42 * x, 42 * y, 31 * y
	*/
	fmt.Printf("%s ", s)
	s1 := s
	e := 0
	count42 := 0
	for {
		s1, e = checkrule(42, s1)
		if e == -1 {
			break
		}
		count42++
	}
	fmt.Printf("count 42/31: %d - ", count42)

	if count42 < 2 {
		fmt.Printf("\n")
		return false
	}

	count31 := 0
	for {
		s1, e = checkrule(31, s1)
		if e == 0 {
			count31++
		}
		if e == -1 || s1 == "" {
			break
		}
	}

	fmt.Printf("%d\n", count31)
	if count31 > 0 && s1 == "" && count31 < count42 {
		return true
	}

	return false

}
