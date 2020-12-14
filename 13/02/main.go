package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var totalcount int
var data = "13,x,x,41,x,x,x,x,x,x,x,x,x,641,x,x,x,x,x,x,x,x,x,x,x,19,x,x,x,x,17,x,x,x,x,x,x,x,x,x,x,x,29,x,661,x,x,x,x,x,37,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,23"
var expected []int

var buslist []int

func main() {

	buses := strings.Split(data, ",")
	timestamp := 1

	for _, busID := range buses {
		var id int
		if busID == "x" {
			id = 1
		} else {
			id, _ = strconv.Atoi(busID)
		}
		buslist = append(buslist, id)
	}

	for {
		taccum := 1
		found := true

		for offset := 0; offset < len(buslist); offset++ {
			if (timestamp+offset)%buslist[offset] != 0 {
				found = false
				break
			}

			taccum *= buslist[offset]
		}

		if found {
			fmt.Printf("time:%d\n", timestamp)
			os.Exit(0)
		}

		timestamp += taccum
	}

}
