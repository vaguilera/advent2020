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
	loadPieces()
	t := getTile(3877)
	t.flip(true)
	t = getTile(2339)
	t.flip(true)
	t = getTile(1459)
	t.flip(true)
	t = getTile(1151)
	t.flip(true)
	//t = getTile(3943)
	//t.flip(true)

	t = getTile(3943) //51
	t.flip(true)
	t = getTile(2243)
	t.flip(true)
	//t = getTile(3259) //70
	//t.flip(true)
	//t = getTile(1223) //80
	//t.flip(true)
	//t = getTile(3877) // 90
	//t.flip(true)
	t = getTile(1877) //100
	t.flip(true)
	//t = getTile(3593) //110
	//t.flip(true)

	doMagic()

}

func loadPieces() {

	file, err := os.Open("pieces.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalcount = 0
	for scanner.Scan() {
		line := scanner.Text()
		id, rots, fx, fy := 0, 0, 0, 0
		fmt.Sscanf(line, "%d %d %d %d", &id, &rots, &fx, &fy)
		fmt.Printf("pieza: %d %d %d %d\n", id, rots, fx, fy)
		t := getTile(id)

		for i := 0; i < rots; i++ {
			t.rot()
		}

		if fx == 1 {
			t.flip(false)
		}

		if fy == 1 {
			t.flip(true)
		}

		piezas = append(piezas, pieza{
			id:      id,
			rots:    rots,
			flipedX: fx,
			flipedY: fy,
		})

	}
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

func getTile(id int) *tile {
	for i, tile := range tiles {
		if tile.id == id {
			return &(tiles[i])
		}
	}

	return nil
}

func doMagic() {

	result := make([]string, 8*12)

	linea := 0
	col := 1
	for _, pieza := range piezas {
		t := getTile(pieza.id)
		for j := 1; j < 9; j++ {
			crow := (8 * linea) + (j - 1)
			//fmt.Printf("crow: %d\n", crow)
			result[crow] += t.gridmod[j][1:9]
		}
		if col == 12 {
			linea++
			col = 1
		} else {
			col++
		}

	}

	// for i := 1; i < 9; i++ {
	// 	for j := 0; j < 12; j++ {
	// 		t := getTile(piezas[j].id)
	// 		result[i-1] += t.gridmod[i][1:9]
	// 	}
	// }

	for _, line := range result {
		fmt.Printf("%s\n", line)
	}

	// for j, line := range result {
	// 	for i := 9; i < 111; i += 10 {
	// 		if line[i:i+1] != line[i+1:i+2] {
	// 			fmt.Printf("ERROR %d %d\n", i, j)
	// 		}
	// 	}
	// }

}
