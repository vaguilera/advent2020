package main

import (
	"fmt"
)

var totalcount int

func main() {

	//15628416
	//11161639
	snumber := 1
	i := 1
	b := 0
	for {
		snumber = (snumber * 7) % 20201227

		if snumber == 15628416 || snumber == 11161639 {
			fmt.Printf("llala: %d %d\n", snumber, i)
			b++
			if b == 2 {
				break
			}
		}
		i++
	}

	fmt.Printf("llala2: %d\n", getCryptKey(15628416, 12435705))
	fmt.Printf("llala2: %d\n", getCryptKey(11161639, 11002971))

}

func getCryptKey(snumber, lsize int) int {
	res := 1
	for i := 0; i < lsize; i++ {
		res = (res * snumber) % 20201227
	}

	return res
}
