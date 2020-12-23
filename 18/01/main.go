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
		processLine(line)

	}

}

func processLine(s string) {
	s2 := strings.ReplaceAll(s, " ", "")

	r, _ := regexp.Compile(`\((\d+)(\+|\*|-|\/)(\d+)\)`)
	ops := r.FindAllStringSubmatch(s2, -1)

	fmt.Printf("formula: %v\n", ops)

	for _, op := range ops {
		n1, _ := strconv.Atoi(op[1])
		n2, _ := strconv.Atoi(op[3])

		res := 0
		if op[2] == "+" {
			res = n1 + n2
		} else {
			res = n1 * n2
		}

		sres := fmt.Sprintf("%d", res)
		s2 = strings.ReplaceAll(s2, op[0], sres)

		fmt.Printf("n1: %s op: %s n2: %s res: %s\n", op[1], op[2], op[3], sres)

	}

	fmt.Printf("lalalal: %s\n", s2)
}
