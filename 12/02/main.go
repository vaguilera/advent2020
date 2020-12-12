package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type instr struct {
	opcode byte
	value  int
}

type turtle struct {
	x, y   int
	wx, wy int
}

var totalcount int
var code []instr

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
		val, _ := strconv.Atoi(line[1:])
		i := instr{
			opcode: line[0],
			value:  val,
		}
		code = append(code, i)
	}

	t := turtle{
		x:  0,
		y:  0,
		wx: 10,
		wy: 1,
	}

	fmt.Printf("totalCount: %v\n", t)
	//	fmt.Printf("c: %f s:%f\n", math.Cos(toRad(90.0)), math.Sin(toRad(90.0)))
	for _, ins := range code {
		t.processInstruction(ins)
	}
	fmt.Printf("totalCount: %v\n", t)
	fmt.Printf("Manhatan: %f\n", math.Abs(float64(t.x))+math.Abs(float64(t.y)))

}

func toRad(angle float64) float64 {
	return angle * (math.Pi / 180)
}

func (t *turtle) processInstruction(in instr) {
	switch in.opcode {
	case 'N':
		t.wy += in.value
		fmt.Printf("N %v\n", t)
	case 'S':
		t.wy -= in.value
		fmt.Printf("S %v\n", t)
	case 'E':
		t.wx += in.value
		fmt.Printf("E %v\n", t)
	case 'W':
		t.wx -= in.value
		fmt.Printf("W %v\n", t)
	case 'L':
		otx := t.wx
		oty := t.wy
		switch in.value {
		case 90:
			t.wx = -oty
			t.wy = otx
		case 180:
			t.wx = -otx
			t.wy = -oty
		case 270:
			t.wx = oty
			t.wy = -otx
		}
		fmt.Printf("R %v\n", t)
		fmt.Printf("L %v\n", t)
	case 'R':
		otx := t.wx
		oty := t.wy
		switch in.value {
		case 90:
			t.wx = oty
			t.wy = -otx
		case 180:
			t.wx = -otx
			t.wy = -oty
		case 270:
			t.wx = -oty
			t.wy = otx
		}
		fmt.Printf("R %v\n", t)
	case 'F':
		t.x += in.value * t.wx
		t.y += in.value * t.wy
		fmt.Printf("F %v\n", t)
	}

}
