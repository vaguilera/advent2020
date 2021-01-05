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

	fmt.Printf("%v\n", mapa)
	fmt.Printf("total: %d\n", len(mapa))

	for i := 0; i < 100; i++ {
		mapa = passDay()
		fmt.Printf("total: %d\n", len(mapa))
	}

}

func processLine(s string) {

	x, y := 0, 0

	splited := r.FindAllStringSubmatch(s, -1)
	//fmt.Printf("%v\n", splited)

	for _, v := range splited {
		x, y = getCoords(v[0], x, y)
	}

	searchAndUpdate(x, y)

}

func getCoords(adr string, x, y int) (int, int) {
	switch adr {
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

	return x, y
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

func search(x, y int) bool {
	for _, v := range mapa {
		if v.x == x && v.y == y {
			return true
		}
	}

	return false

}

func blackcount(x, y int) int {

	dirs := []string{"e", "se", "sw", "w", "nw", "ne"}
	count := 0

	for _, dir := range dirs {
		tx, ty := getCoords(dir, x, y)
		if search(tx, ty) {
			count++
		}
	}

	return count

}

func passDay() []tile {
	var res = []tile{}
	minX, maxX, minY, maxY := 0, 0, 0, 0

	for _, v := range mapa {
		if v.x < minX {
			minX = v.x
		}
		if v.x > maxX {
			maxX = v.x
		}
		if v.y < minY {
			minY = v.y
		}
		if v.y > maxY {
			maxY = v.y
		}
	}

	fmt.Printf("minX: %d, maxX: %d, minY: %d, maxY: %d", minX, maxX, minY, maxY)
	minX--
	minY--
	maxX++
	maxY++

	for j := minY; j <= maxY; j++ {
		for i := minX; i <= maxX; i++ {
			isblack := search(i, j)
			count := blackcount(i, j)

			if isblack {
				if count == 1 || count == 2 {
					res = append(res, tile{i, j})
				}
			} else {
				if count == 2 {
					res = append(res, tile{i, j})
				}
			}

			//fmt.Printf("x: %d, y:%d %t\n", i, j, search(i, j))
		}

	}

	return res

}
