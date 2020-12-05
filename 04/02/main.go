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

	//fmt.Printf("RES: %t\n", processLine("pid:337605855 cid:249 byr:1952 hgt:155cm ecl:grn iyr:2017 eyr:2026 hcl:#866857"))

	fmt.Printf("TOTAL: %d\n", totalcount)

}

func processLine(s string) bool {
	fields := strings.Split(s, " ")

	ret := false

	if len(fields) == 7 && hasCID(fields) {
		return false
	}

	if (len(fields) >= 7) && validateFields(fields) {
		ret = true
	}

	//fmt.Printf("%s ---- %d %t\n", s, len(fields), ret)
	// if ret {
	// 	fmt.Printf(" %t\n", ret)
	// }
	return ret

}

func validateFields(fields []string) bool {

	for _, field := range fields {
		keys := strings.Split(field, ":")
		if !performValidation(keys[0], keys[1]) {
			return false
		}
	}
	return true
}

func performValidation(key string, value string) bool {
	eclAllowed := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	switch key {
	case "byr":
		v := validateNumber(value, 1920, 2002)
		return v
	case "iyr":
		return validateNumber(value, 2010, 2020)
	case "eyr":
		return validateNumber(value, 2020, 2030)
	case "hgt":
		v := validateHeigh(value)
		return v
	case "hcl":
		r, _ := regexp.Compile("#(?:[0-9a-f]{6})")
		v := r.MatchString(value)
		return v
	case "ecl":
		for _, item := range eclAllowed {
			if item == value {
				return true
			}
		}
		return false
	case "pid":
		r, _ := regexp.Compile(`^(\d){9}$`)
		v := r.MatchString(value)
		v2 := v && (len(value) == 9)
		fmt.Printf("%s %t %d\n", value, v2, len(value))
		return v2
	case "cid":
		return true
	}

	return false
}

func validateHeigh(s string) bool {
	if len(s) < 3 {
		return false
	}

	n := s[0 : len(s)-2]
	unit := s[len(s)-2:]

	switch unit {
	case "cm":
		return validateNumber(n, 150, 193)
	case "in":
		return validateNumber(n, 59, 76)
	default:
		//fmt.Printf("NO HEIGHT EN CM O IN\n")
		return false
	}

}

func validateNumber(s string, min int, max int) bool {
	i, err := strconv.Atoi(s)
	// if s == "2009" {
	// 	fmt.Printf("\n\n2009!!!!!!!!!!!!!!!!!!!! %d %d %d \n\n\n", i, min, max)
	// }
	if err != nil {
		return false
	}
	if (i < min) || (i > max) {
		return false
	}
	return true
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
