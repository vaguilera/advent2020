package main

import (
	"fmt"
)

var totalcount int
var initData = []int{6, 4, 3, 7, 1, 9, 2, 5, 8}

const MINVALUE = 1
const MAXVALUE = 9

func main() {

	init := 6
	initData := removeStep1(initData, init)

	for i := 0; i < 99; i++ {
		pos := (findPos(initData, init) + 1) % len(initData)

		init = initData[pos]
		initData = removeStep1(initData, init)
	}
	fmt.Printf("Final: %v\n", initData)

}

func findPos(c []int, val int) int {
	for i, v := range c {
		if v == val {
			return i
		}
	}
	return -1
}
func removeStep1(c []int, init int) []int {
	pick := make([]int, 3)

	fmt.Printf("Current: %d\n", init)
	pos := findPos(c, init)

	pick[0] = c[(pos+1)%len(c)]
	pick[1] = c[(pos+2)%len(c)]
	pick[2] = c[(pos+3)%len(c)]

	//fmt.Printf("pick: %v\n", pick)

	des := init - 1
	if des <= MINVALUE-1 {
		des = MAXVALUE
	}
	for des == pick[0] || des == pick[1] || des == pick[2] {
		des--
		if des <= MINVALUE-1 {
			des = MAXVALUE
		}
	}
	fmt.Printf("cups: %v\n", c)
	fmt.Printf("pick: %v\n", pick)
	fmt.Printf("destination: %d\n", des)
	rsub := removeSub(c, pick)
	res := injectSlice(rsub, pick, des)
	//fmt.Printf("res : %v\n", res)
	return res
}

func removeSub(c []int, pick []int) []int {

	res := []int{}

	for _, v := range c {
		found := false
		for _, v2 := range pick {
			if v == v2 {
				found = true
				break
			}
		}
		if !found {
			res = append(res, v)
		}
	}
	return res
}

func injectSlice(c []int, pick []int, pos int) []int {
	res := []int{}

	for _, v := range c {
		if pos != v {
			res = append(res, v)
			continue
		}
		res = append(res, v)
		for _, v2 := range pick {
			res = append(res, v2)
		}

	}
	return res
}
