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
	x, y  int
	angle int
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
		x:     0,
		y:     0,
		angle: 90,
	}
	fmt.Printf("totalCount: %v\n", t)
	fmt.Printf("c: %f s:%f\n", math.Cos(toRad(180.0)), math.Sin(toRad(180.0)))
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
		t.y += in.value
		fmt.Printf("N %v\n", t)
	case 'S':
		t.y -= in.value
		fmt.Printf("S %v\n", t)
	case 'E':
		t.x += in.value
		fmt.Printf("E %v\n", t)
	case 'W':
		t.x -= in.value
		fmt.Printf("W %v\n", t)
	case 'L':
		t.angle -= in.value % 360
		fmt.Printf("L %v\n", t)
	case 'R':
		t.angle += in.value % 360
		fmt.Printf("R %v\n", t)
	case 'F':
		t.y += (in.value * int(math.Cos(toRad(float64(t.angle)))))
		t.x += (in.value * int(math.Sin(toRad(float64(t.angle)))))
		fmt.Printf("F %v\n", t)
	}

}
