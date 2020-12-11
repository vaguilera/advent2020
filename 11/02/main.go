package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var totalcount int
var mapa []string

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
		mapa = append(mapa, line)
	}

	m1 := generateMap(mapa)
	m2 := generateMap(m1)

	for {
		m1 = generateMap(m2)
		m2 = generateMap(m1)
		if equals(m1, m2) {
			break
		}
	}

	for j := 0; j < len(m1); j++ {
		for i := 0; i < len(m1[j]); i++ {
			if m1[j][i] == '#' {
				totalcount++
			}
		}
	}

	fmt.Printf("totalCount: %d\n", totalcount)

}

func decorateMap(mapa []string) []string {
	var mapa2 []string

	emptyline := strings.Repeat(".", len(mapa[0])+2)
	mapa2 = append(mapa2, emptyline)
	for _, lin := range mapa {
		mapa2 = append(mapa2, "."+lin+".")
	}
	mapa2 = append(mapa2, emptyline)
	return mapa2
}

func generateMap(mapa []string) []string {
	var mapa2 []string

	mapaDeco := decorateMap(mapa)

	for j := 1; j < len(mapaDeco)-1; j++ {
		cline := ""
		for i := 1; i < len(mapaDeco[0])-1; i++ {

			occ, _ := checkSeat(mapaDeco, i, j)

			switch mapaDeco[j][i] {
			case '.':
				cline += "."
			case 'L':
				if occ == 0 {
					cline += "#"
				} else {
					cline += "L"
				}
			case '#':
				if occ >= 5 {
					cline += "L"
				} else {
					cline += "#"
				}
			}
		}
		mapa2 = append(mapa2, cline)
	}

	return mapa2

}

func lookAt(m []string, ox, oy, dx, dy int) rune {

	dex := dx
	dey := dy
	for {

		if (ox+dex > len(m[0])-1) || (ox+dex < 0) {
			return '.'
		}
		if (oy+dey > len(m)-1) || (oy+dey < 0) {
			return '.'
		}

		if m[oy+dey][ox+dex] == 'L' {
			return 'L'
		}
		if m[oy+dey][ox+dex] == '#' {
			return '#'
		}

		dey += dy
		dex += dx
	}
}

func checkSeat(mapa []string, x, y int) (int, int) {
	occ, free := 0, 0

	for j := -1; j < 2; j++ {
		for i := -1; i < 2; i++ {
			if j == 0 && i == 0 {
				continue
			}
			r := lookAt(mapa, x, y, i, j)
			if r != '#' {
				free++
			} else {
				occ++
			}
		}
	}
	return occ, free

}

func equals(m1 []string, m2 []string) bool {

	for j := 0; j < len(m1); j++ {
		for i := 0; i < len(m1[j]); i++ {
			if m1[j][i] != m2[j][i] {
				return false
			}
		}
	}
	return true
}
