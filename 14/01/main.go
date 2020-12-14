package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var totalcount int64

var maskOn int64
var maskOff int64

var mem map[int]int64

func main() {

	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalcount = 0
	mem = make(map[int]int64)
	r, _ := regexp.Compile(`mem\[(\d*)\]\s=\s(\d*)`)

	for scanner.Scan() {
		line := scanner.Text()
		if line[0:4] == "mask" {
			processMask(line[7:])
		} else {
			splited := r.FindStringSubmatch(line)
			processMem(splited[1], splited[2])
		}
	}

	for _, el := range mem {
		totalcount += el
	}

	fmt.Printf("total: %d\n", totalcount)

}

func processMask(mask string) {
	fmt.Printf("%s\n", mask)

	var smaskOn, smaskOff string

	for _, bit := range mask {
		switch bit {
		case 'X':
			smaskOff += "1"
			smaskOn += "0"
		case '0':
			smaskOff += "0"
			smaskOn += "0"
		case '1':
			smaskOff += "1"
			smaskOn += "1"

		}

	}

	maskOn, _ = strconv.ParseInt(smaskOn, 2, 64)
	maskOff, _ = strconv.ParseInt(smaskOff, 2, 64)
}

func processMem(smem, val string) {

	fmt.Printf("%s - %s,", smem, val)
	imem, _ := strconv.Atoi(smem)

	ival, _ := strconv.Atoi(val)
	val64 := (int64(ival) | maskOn) & maskOff

	fmt.Printf("%d\n", val64)
	mem[imem] = val64
}
