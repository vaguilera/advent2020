package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var totalcount int

var data []string

var drake []string

type pos struct {
	x, y int
}

var dpos = [15]pos{
	pos{x: 18, y: 0},
	pos{x: 0, y: 1},
	pos{x: 5, y: 1},
	pos{x: 6, y: 1},
	pos{x: 11, y: 1},
	pos{x: 12, y: 1},
	pos{x: 17, y: 1},
	pos{x: 18, y: 1},
	pos{x: 19, y: 1},
	pos{x: 1, y: 2},
	pos{x: 4, y: 2},
	pos{x: 7, y: 2},
	pos{x: 10, y: 2},
	pos{x: 13, y: 2},
	pos{x: 16, y: 2},
}

func main() {

	drake = []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}

	file, err := os.Open("puzzle.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalcount = 0

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	for _, l := range data {
		for _, c := range l {
			if c == '#' {
				totalcount++
			}
		}
		//fmt.Printf("%s\n", l)
	}

	// for _, l := range data {
	// 	fmt.Printf("%s\n", l)
	// }
	// data = rot()
	// fmt.Printf("\n\n")
	// for _, l := range data {
	// 	fmt.Printf("%s\n", l)
	// }

	data = flip(false)

	data = rot()

	data = rot()
	//data = rot()

	drakes := 0

	fmt.Printf("total %d\n", totalcount)
	for y := 0; y < len(data)-3; y++ {
		for x := 0; x < len(data[0])-20; x++ {
			d := checkDrake(x, y)
			if d {
				fmt.Printf("true\n")
				drakes++
			}

		}
	}

	fmt.Printf("%d %d %d\n", totalcount, drakes*15, totalcount-(drakes*15))

}

func checkDrake(x, y int) bool {
	region := []string{
		data[y][x : x+20],
		data[y+1][x : x+20],
		data[y+2][x : x+20],
	}

	for _, dp := range dpos {
		if region[dp.y][dp.x] != '#' {
			return false
		}
	}

	return true
}

func rot() []string {
	tmp := []string{}
	for i := 0; i < 96; i++ {
		s := ""
		for j := 95; j >= 0; j-- {
			s += data[j][i : i+1]
		}
		tmp = append(tmp, s)
	}
	return tmp
}

func flip(x bool) []string {
	tmp := []string{}
	if x {
		for _, line := range data {
			s := ""
			for i := 0; i < 96; i++ {
				s += line[95-i : 95-i+1]
			}
			tmp = append(tmp, s)
		}
	} else {
		for nline := range data {
			tmp = append(tmp, data[95-nline])
		}
	}

	return tmp

}
