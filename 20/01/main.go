package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var totalcount int

var tiles []tile

type tile struct {
	id      int
	grid    []string
	gridmod []string
	top     []int
	left    []int
	right   []int
	bottom  []int
}

func main() {

	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalcount = 0
	nline := 0

	cgrid := []string{}
	id := 0
	for scanner.Scan() {
		line := scanner.Text()
		if nline == 0 {
			id, _ = strconv.Atoi(line[5:9])
			cgrid = []string{}
			nline++
			continue
		}

		if nline <= 10 {
			cgrid = append(cgrid, line)
			nline++
			continue
		}

		til := tile{
			id: id,
		}
		til.load(cgrid)
		tiles = append(tiles, til)
		nline = 0
	}

	fmt.Printf("len: %d\n", len(tiles))
	doMagic()
}

func (t *tile) load(s []string) {
	t.grid = s
	t.top = []int{}
	t.gridmod = make([]string, 10)
	copy(t.gridmod, t.grid)
}

func (t *tile) print() {
	fmt.Printf("Tile: %d\n", t.id)
	for _, line := range t.gridmod {
		fmt.Printf("%s\n", line)
	}
	fmt.Printf("\n")

}

func (t *tile) flip(x bool) {
	tmp := []string{}
	if x {
		for _, line := range t.gridmod {
			s := ""
			for i := 0; i < 10; i++ {
				s += line[9-i : 9-i+1]
			}
			tmp = append(tmp, s)
		}
	} else {
		for nline := range t.gridmod {
			tmp = append(tmp, t.gridmod[9-nline])
		}
	}

	t.gridmod = make([]string, 10)
	copy(t.gridmod, tmp)

}

func (t *tile) rot() {
	tmp := []string{}
	for i := 0; i < 10; i++ {
		s := ""
		for j := 9; j >= 0; j-- {
			s += t.gridmod[j][i : i+1]
		}
		tmp = append(tmp, s)
	}
	t.gridmod = make([]string, 10)
	copy(t.gridmod, tmp)
}

func (t *tile) reset() {
	t.gridmod = make([]string, 10)
	copy(t.gridmod, t.grid)
}

func (t *tile) getborder(border, rottimes int, flipX, flipY bool) string {

	if flipX {
		t.flip(true)
	}
	if flipY {
		t.flip(false)
	}

	for i := 0; i < rottimes; i++ {
		t.rot()
	}

	switch border {
	case 0:
		return t.gridmod[0]
	case 1:
		s := ""
		for i := 0; i < 10; i++ {
			s += t.gridmod[i][9:10]
		}
		return s
	case 2:
		return t.gridmod[9]
	case 3:
		s := ""
		for i := 0; i < 10; i++ {
			s += t.gridmod[i][0:1]
		}
		return s
	}

	return "ERROR"
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func doMagic() {

	for i, t := range tiles {
		for side := 0; side < 4; side++ {
			bor := t.getborder(side, 0, false, false)
			for _, t2 := range tiles {
				if t2.id == t.id {
					continue
				}
				for side2 := 0; side2 < 4; side2++ {
					bor2 := t2.getborder(side2, 0, false, false)
					if bor2 == bor || bor2 == reverse(bor) {
						switch side {
						case 0:
							tiles[i].top = append(tiles[i].top, t2.id)
						case 1:
							tiles[i].right = append(tiles[i].right, t2.id)
						case 2:
							tiles[i].bottom = append(tiles[i].bottom, t2.id)
						case 3:
							tiles[i].left = append(tiles[i].left, t2.id)
						}
					}
				}
			}
		}

	}

	result := 1
	for _, til := range tiles {
		t := len(til.top)
		l := len(til.left)
		r := len(til.right)
		b := len(til.bottom)
		numborders := t + l + r + b

		if numborders == 2 {
			fmt.Printf("id: %d - borders: %d\n", til.id, numborders)
			result *= til.id
		}
	}

	fmt.Printf("result: %d\n", result)

}

func getTile(id int) *tile {
	for _, tile := range tiles {
		if tile.id == id {
			return &tile
		}
	}

	return nil
}
