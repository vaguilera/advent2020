package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type content struct {
	count int
	name  string
}

type bag struct {
	name    string
	content []content
}

var totalcount int
var bags []bag

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
		bags = append(bags, processLine(line))
	}

	for _, b := range bags {
		totalcount += findGold(b)
	}

	fmt.Printf("%d\n", totalcount)
}

func processLine(s string) bag {
	r, _ := regexp.Compile(`(\w+ \w+) bags contain (((?:(?:, )?\d \w+ \w+ bags?)+)|(no other bags))\.`)
	r2, _ := regexp.Compile(`(\d) (\w+ \w+) bags?`)
	splited := r.FindStringSubmatch(s)

	cont := []content{}

	if splited[2] == "no other bags" {
		return bag{
			splited[1],
			cont,
		}
	}

	subbags := strings.Split(splited[2], ", ")

	for _, b := range subbags {
		subbag := r2.FindStringSubmatch(b)
		i, _ := strconv.Atoi(subbag[1])
		cont = append(cont, content{i, subbag[2]})
	}

	return bag{
		splited[1],
		cont,
	}

}

func findGold(b bag) int {
	//fmt.Printf("looking into: %s\n", b.name)
	if len(b.content) == 0 {
		//fmt.Printf("%s no more bags\n", b.name)
		return 0
	}

	for _, subbag := range b.content {
		if subbag.name == "shiny gold" {
			//fmt.Printf("%s contains shiny gold\n", b.name)
			return 1
		}
		bb := findBag(subbag.name)
		if bb == nil {
			err := errors.New(fmt.Sprint("NO ENCUENTRO BOLSA: %s\n", subbag.name))
			panic(err)
		}
		if findGold(*bb) > 0 {
			return 1
		}
	}

	return 0
}

func findBag(s string) *bag {
	for _, b := range bags {
		if b.name == s {
			return &b
		}
	}
	return nil
}
