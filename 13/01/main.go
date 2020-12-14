package main

import (
	"fmt"
	"math"
)

var totalcount int
var buses = []int{13, 41, 641, 19, 17, 29, 661, 37, 23}

const mytime = 1006697

func main() {

	for _, bus := range buses {
		t := math.Ceil(mytime / float64(bus))
		diff := (int(t) * bus) - mytime
		fmt.Printf("val: %d - %d - %d\n", int(t), diff, diff*bus)

	}

}
