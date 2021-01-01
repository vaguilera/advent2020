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
var piezas []pieza

type tile struct {
	id      int
	grid    []string
	gridmod []string
	top     int
	right   int
	bottom  int
	left    int
	numrots int
	flipedX int
	flipedY int
}

type pieza struct {
	id      int
	rots    int
	flipedX int
	flipedY int
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
	//copy(grid, t.grid)
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
	piezas = []pieza{pieza{
		id:      1289,
		rots:    0,
		flipedX: 0,
		flipedY: 0,
	},
	}

	doMagic2()
	doMagic3()

}

func (t *tile) load(s []string) {
	t.grid = s
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

func (t *tile) flip() {
	to, b := t.top, t.bottom
	t.top = b
	t.bottom = to
	if t.flipedX == 0 {
		t.flipedX = 1
	} else {
		t.flipedX = 0
	}
}

func (t *tile) flipY() {
	l, r := t.left, t.right
	t.left = r
	t.right = l
	if t.flipedY == 0 {
		t.flipedY = 1
	} else {
		t.flipedY = 0
	}
}

func (t *tile) rot() {
	to, r, b, l := t.top, t.right, t.bottom, t.left

	t.top = l
	t.right = to
	t.bottom = r
	t.left = b
}

func (t *tile) reset() {
	t.gridmod = make([]string, 10)
	copy(t.gridmod, t.grid)
}

func (t *tile) getborder(border int) string {
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
			bor := t.getborder(side)
			for _, t2 := range tiles {
				if t2.id == t.id {
					continue
				}
				for side2 := 0; side2 < 4; side2++ {
					bor2 := t2.getborder(side2)
					if bor2 == bor || bor2 == reverse(bor) {
						switch side {
						case 0:
							tiles[i].top = t2.id
						case 1:
							tiles[i].right = t2.id
						case 2:
							tiles[i].bottom = t2.id
						case 3:
							tiles[i].left = t2.id
						}
					}
				}
			}
		}

	}

	result := 1
	for _, til := range tiles {
		numborders := 0
		if til.top > 0 {
			numborders++
		}
		if til.right > 0 {
			numborders++
		}
		if til.left > 0 {
			numborders++
		}
		if til.bottom > 0 {
			numborders++
		}

		if numborders == 2 {
			//fmt.Printf("id: %d - borders: %d\n", til.id, numborders)
			fmt.Printf("%v\n", til)
			result *= til.id
		}
	}

	fmt.Printf("result: %d\n", result)

}

func getTile(id int) *tile {
	for i, tile := range tiles {
		if tile.id == id {
			return &(tiles[i])
		}
	}

	return nil
}

func doMagic2() {

	for {
		lastTile := getTile(piezas[len(piezas)-1].id)
		ctile := getTile(lastTile.right)

		for ok := true; ok; ok = (ctile.left != lastTile.id) {
			ctile.rot()
			ctile.numrots++
		}

		if ctile.top != 0 {
			ctile.flip()
		}

		piezas = append(piezas,
			pieza{
				id:      ctile.id,
				rots:    ctile.numrots,
				flipedX: ctile.flipedX,
				flipedY: ctile.flipedY,
			},
		)
		if ctile.right == 0 {
			break
		}
	}

}

func doMagic3() {

	lastTileID := 0
	for j := 0; j < 11; j++ {
		for i := 0; i < 12; i++ {
			til1 := getTile(piezas[i+(j*12)].id)
			ctile := getTile(til1.bottom)

			for ok := true; ok; ok = (ctile.top != til1.id) {
				ctile.rot()
				ctile.numrots++
			}

			if ctile.left != lastTileID {
				ctile.flipY()
			}
			//fmt.Printf("%d %d %d %d\n", ctile.id, ctile.numrots, ctile.flipedX, ctile.flipedY)
			lastTileID = ctile.id
			piezas = append(piezas,
				pieza{
					id:      ctile.id,
					rots:    ctile.numrots,
					flipedX: ctile.flipedX,
					flipedY: ctile.flipedY,
				},
			)
		}
	}

	for _, p := range piezas {
		if p.id == 1289 {
			fmt.Printf("AQUI\n")
		}
		fmt.Printf("%d %d %d %d\n", p.id, p.rots, p.flipedX, p.flipedY)
	}

	fmt.Printf("LEN: %d\n", len(piezas))

}
