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

	execute()
	//fmt.Printf("%v\n", code)
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

func execute() {
	ip := 0
	acc := 0

	for {
		inst := &(code[ip])
		if inst.times > 0 {
			fmt.Printf("ya pase por aqui. IP: %d - ACC: %d\n", ip, acc)
			break
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
