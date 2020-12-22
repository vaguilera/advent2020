package main

import (
	"fmt"
	"os"
)

var totalcount int

type value struct {
	turn     int
	lastTurn int
}

func main() {

	numList := []int{15, 12, 0, 14, 3, 1}

	db := make(map[int]value)

	for i, val := range numList {
		db[val] = value{
			turn:     i + 1,
			lastTurn: -1,
		}
	}

	turn := len(numList)
	last := numList[len(numList)-1]

	for {
		turn++
		previous := db[last]

		if previous.lastTurn == -1 {

			var lastTurn0 int
			db0, ok := db[0]
			if !ok {
				lastTurn0 = -1
			} else {
				lastTurn0 = db0.turn
			}
			db[0] = value{
				turn:     turn,
				lastTurn: lastTurn0,
			}
			last = 0
		} else {
			cnum := (turn - 1) - previous.lastTurn
			//fmt.Printf("cnum: %d\n", cnum)
			dbX, ok := db[cnum]
			var lastTurnX int
			if !ok {
				lastTurnX = -1
			} else {
				lastTurnX = dbX.turn
			}
			//fmt.Printf("dbX, ok, lastTurnX: %v %t %d\n", dbX, ok, lastTurnX)
			db[cnum] = value{
				turn:     turn,
				lastTurn: lastTurnX,
			}
			last = cnum

			if turn == 30000000 {
				fmt.Printf("cnum: %d\n", cnum)
				os.Exit(0)
			}

		}
	}
}

func lastTime(list []int, n int) int {

	for i := len(list) - 2; i >= 0; i-- {
		if list[i] == n {
			return i + 1
		}
	}

	return -1
}
