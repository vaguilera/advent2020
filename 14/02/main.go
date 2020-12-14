package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var totalcount int
var cmask string
var addresses []string
var mem map[int64]int

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalcount = 0
	r, _ := regexp.Compile(`mem\[(\d*)\]\s=\s(\d*)`)
	mem = make(map[int64]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0:4] == "mask" {
			cmask = line[7:]
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

func processMem(smem, val string) {

	fmt.Printf("%s - %s - ", smem, val)
	imem, _ := strconv.Atoi(smem)
	bmem := fmt.Sprintf("%b", imem)
	ival, _ := strconv.Atoi(val)

	padding := strings.Repeat("0", len(cmask)-len(bmem))
	maskedAddr := getMaskedAddr(padding+bmem, cmask)
	fmt.Printf("%s\n", maskedAddr)

	addresses = []string{}
	generateAddresses(maskedAddr)
	for _, v := range addresses {
		iaddr, _ := strconv.ParseInt(v, 2, 64)
		mem[iaddr] = ival
	}
}

func getMaskedAddr(mem, mask string) string {
	res := ""
	for i, v := range mask {
		switch v {
		case '0':
			res += string(mem[i])
		default:
			res += string(v)
		}
	}

	return res
}

func generateAddresses(mask string) {
	var opA, opB string
	var i int
	for i = 0; i < len(mask); i++ {
		if mask[i] == 'X' {
			opA = mask[:i] + "0" + mask[i+1:]
			opB = mask[:i] + "1" + mask[i+1:]
			generateAddresses(opA)
			generateAddresses(opB)
			break
		}
	}
	if i == 36 {
		addresses = append(addresses, mask)
	}
}
