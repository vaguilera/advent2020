package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	opcode string
	param  int
	times  int
}

var code []instruction

var totalcount int

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
		code = append(code, processLine(line))
	}

	for i, instr := range code {
		code2 := make([]instruction, len(code))
		copy(code2, code)

		if instr.opcode == "nop" {
			code2[i].opcode = "jmp"
		} else if instr.opcode == "jmp" {
			code2[i].opcode = "nop"
		}

		res := execute(code2)
		//fmt.Printf("res: %d\n", res)
		if res > -1 {
			os.Exit(0) // a lo bruto
		}

	}
}

func processLine(s string) instruction {
	tokens := strings.Split(s, " ")
	param, _ := strconv.Atoi(tokens[1])

	return instruction{
		opcode: tokens[0],
		param:  param,
		times:  0,
	}
}

func execute(code []instruction) int {
	ip := 0
	acc := 0

	for {
		if ip >= len(code) {
			fmt.Printf("SE FINI. IP: %d - ACC: %d\n", ip, acc)
			return acc
		}

		inst := &(code[ip])
		if inst.times > 0 {
			fmt.Printf("ya pase por aqui. IP: %d - ACC: %d\n", ip, acc)
			return -1
		}

		switch inst.opcode {
		case "nop":
			inst.times++
			ip++
		case "acc":
			acc += inst.param
			inst.times++
			ip++
		case "jmp":
			ip += inst.param
			inst.times++
		default:
			err := errors.New("INVALID OPCODE")
			panic(err)
		}

	}
}
