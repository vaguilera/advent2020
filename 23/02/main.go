package main

import (
	"fmt"
)

var totalcount int

var initData = []int{6, 4, 3, 7, 1, 9, 2, 5, 8}

//var initData = []int{3, 8, 9, 1, 2, 5, 4, 6, 7}

var db map[int]int
var current int

const MAXVALUE = 1000000

func main() {

	db = make(map[int]int)

	for i := 0; i < len(initData)-1; i++ {
		db[initData[i]] = initData[i+1]
	}
	db[8] = 10
	for i := 10; i < 1000000; i++ {
		db[i] = i + 1
	}
	db[1000000] = 6

	current = 6
	turn(10000000)

	a, b := db[1], db[db[1]]
	fmt.Printf("a: %d b: %d - mult: %d\n", a, b, a*b)

}

func printPos(p int) {
	fmt.Printf("db: %d ", p)
	pos := p
	for i := 0; i < len(db)-1; i++ {
		fmt.Printf("%d ", db[pos])
		pos = db[pos]
	}
	fmt.Println("")

}

func turn(rounds int) {
	for {
		p1 := db[current]
		p2 := db[p1]
		p3 := db[p2]
		final := db[p3]

		des := current - 1
		if des == 0 {
			des = MAXVALUE
		}

		for des == p1 || des == p2 || des == p3 {
			des--
			if des == 0 {
				des = MAXVALUE
			}
		}

		db[current] = final
		tmp := db[des]
		db[des] = p1
		db[p3] = tmp

		current = final
		rounds--
		if rounds == 0 {
			break
		}
	}

}
