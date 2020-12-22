package main

import (
	"fmt"
)

var totalcount int

type value struct {
	turn  int
	first bool
}

func main() {

	numList := []int{15, 12, 0, 14, 3, 1}

	totalcount = 0
	turn := len(numList)

	for {
		turn++

		last := numList[len(numList)-1]
		lastPos := lastTime(numList, last)

		//fmt.Printf("last: %d, pos: %d\n", last, lastPos)

		if lastPos == -1 {
			numList = append(numList, 0)
		} else {
			cnum := (turn - 1) - lastPos
			numList = append(numList, cnum)
		}

		if turn == 2020 {
			break
		}
	}

	fmt.Printf("nums: %d\n", numList[2019])

}

func lastTime(list []int, n int) int {

	for i := len(list) - 2; i >= 0; i-- {
		if list[i] == n {
			return i + 1
		}
	}

	return -1
}
