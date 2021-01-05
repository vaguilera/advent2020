package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var totalcount int
var r *regexp.Regexp

type tile struct {
	x, y int
}

var mapa []tile

func main() {

	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf("failed to open")
	}
	r, _ = regexp.Compile(`(e|se|sw|w|nw|ne)`)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalcount = 0
	for scanner.Scan() {
		line := scanner.Text()
		processLine(line)

	}

	//fmt.Printf("%v\n", mapa)
	fmt.Printf("total: %d\n", len(mapa))
}

func processLine(s string) {

	x, y := 0, 0

	splited := r.FindAllStringSubmatch(s, -1)
	//fmt.Printf("%v\n", splited)

	for _, v := range splited {
		switch v[0] {
		case "se":
			if (y % 2) != 0 {
				x++
			}
			y++

		case "sw":
			if (y % 2) == 0 {
				x--
			}
			y++
		case "w":
			x--
		case "e":
			x++
		case "ne":
			if (y % 2) != 0 {
				x++
			}
			y--
		case "nw":
			if (y % 2) == 0 {
				x--
			}
			y--

		}
	}

	searchAndUpdate(x, y)

}

func searchAndUpdate(x, y int) {

	for i, v := range mapa {
		if v.x == x && v.y == y {
			mapa[i] = mapa[len(mapa)-1]
			mapa = mapa[:len(mapa)-1]
			return
		}
	}

	mapa = append(mapa, tile{
		x: x, y: y,
	})
}
