package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var totalcount int

type rule struct {
	name                   string
	minA, minB, maxA, maxB int
}

var rules = []rule{
	rule{
		name: "departure location",
		minA: 25, maxA: 568,
		minB: 594, maxB: 957,
	},
	rule{
		name: "departure station",
		minA: 33, maxA: 447,
		minB: 466, maxB: 952,
	},
	rule{
		name: "departure platform",
		minA: 31, maxA: 700,
		minB: 725, maxB: 956,
	},
	rule{
		name: "departure track",
		minA: 43, maxA: 124,
		minB: 141, maxB: 952,
	},
	rule{
		name: "departure date",
		minA: 26, maxA: 290,
		minB: 306, maxB: 962,
	},
	rule{
		name: "departure time",
		minA: 34, maxA: 754,
		minB: 763, maxB: 960},
	rule{
		name: "arrival location",
		minA: 29, maxA: 208,
		minB: 217, maxB: 958,
	},
	rule{
		name: "arrival station",
		minA: 48, maxA: 118,
		minB: 124, maxB: 973,
	},
	rule{
		name: "arrival platform",
		minA: 35, maxA: 368,
		minB: 389, maxB: 972,
	},
	rule{
		name: "arrival track",
		minA: 47, maxA: 91,
		minB: 106, maxB: 970,
	},
	rule{
		name: "class",
		minA: 35, maxA: 521,
		minB: 528, maxB: 960,
	},
	rule{
		name: "duration",
		minA: 27, maxA: 833,
		minB: 855, maxB: 965,
	},
	rule{
		name: "price",
		minA: 25, maxA: 870,
		minB: 895, maxB: 957,
	},
	rule{
		name: "route",
		minA: 31, maxA: 140,
		minB: 146, maxB: 965,
	},
	rule{
		name: "row",
		minA: 35, maxA: 736,
		minB: 743, maxB: 957,
	},
	rule{
		name: "seat",
		minA: 33, maxA: 227,
		minB: 249, maxB: 961,
	},
	rule{
		name: "train",
		minA: 27, maxA: 763,
		minB: 788, maxB: 961,
	},
	rule{
		name: "type: ",
		minA: 34, maxA: 167,
		minB: 193, maxB: 950,
	},
	rule{
		name: "wagon",
		minA: 47, maxA: 437,
		minB: 443, maxB: 952,
	},
	rule{
		name: "zone",
		minA: 48, maxA: 928,
		minB: 940, maxB: 955,
	},
}

var ticket = []int{113, 197, 59, 167, 151, 107, 79, 73, 109, 157, 199, 193, 83, 53, 89, 71, 149, 61, 67, 163}
var validLines [][]int

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
		v := processLine(line)
		if v == -1 {
			validLines = append(validLines, parse(line))
		}
	}
	fmt.Printf("llala: %d\n", len(validLines))
	doMagic()
}

func processLine(s string) int {
	nums := strings.Split(s, ",")

	for _, val := range nums {
		i, _ := strconv.Atoi(val)
		if checkNumber(i) == false {
			return i
		}
	}
	return -1
}

func checkNumber(n int) bool {
	for _, r := range rules {
		if n >= r.minA && n <= r.maxA {
			return true
		}
		if n >= r.minB && n <= r.maxB {
			return true
		}
	}

	return false
}

func parse(s string) []int {
	nums := strings.Split(s, ",")
	var ret []int
	for _, val := range nums {
		i, _ := strconv.Atoi(val)
		ret = append(ret, i)
	}
	return ret
}

func checkNum(n int, r rule) int {
	//for i, r := range rules {
	if (n >= r.minA && n <= r.maxA) || (n >= r.minB && n <= r.maxB) {
		return 1
		//	}
	}

	return 0
}

func doMagic() {

	var cols []int
	var rulz []int

	for i := 0; i < len(ticket); i++ {
		for col := 0; col < len(ticket); col++ {
			//fmt.Printf("col:%d rules: ", col)
			if find(cols, col) != -1 {
				continue
			}
			rulpasses := []int{}
			for rul := 0; rul < len(rules); rul++ {
				if find(rulz, rul) != -1 {
					continue
				}
				valid := allPass(col, rul)
				if valid {
					rulpasses = append(rulpasses, rul)
				}
			}

			if len(rulpasses) == 1 {
				cols = append(cols, col)
				rulz = append(rulz, rulpasses[0])
			}

		}
	}

	for i, v := range cols {
		fmt.Printf("[%d - %d]\n", v, rulz[i])
	}

}

func allPass(col int, rule int) bool {
	for _, lin := range validLines {
		numrules := checkNum(lin[col], rules[rule])
		if numrules == 0 {
			return false
		}
	}

	return true

}

func find(slice []int, val int) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}
	return -1
}
