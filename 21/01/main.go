package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var totalcount int

var finalRelation map[string]string
var alergens map[string][]string
var rawIngredients []string

func main() {

	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	alergens = make(map[string][]string)
	finalRelation = make(map[string]string)
	totalcount = 0
	for scanner.Scan() {
		line := scanner.Text()
		processLine(line)
	}

	part1()
}

func part1() {
	for al, ing := range alergens {
		fmt.Printf("al: %s %v\n", al, ing)
	}

	for len(alergens) > 0 {
		for al, ings := range alergens {
			if len(ings) == 1 {
				for a, i := range alergens {
					alergens[a] = removeItem(i, ings[0])
				}
				finalRelation[ings[0]] = al
				delete(alergens, al)
				break
			}
		}
	}

	for al, ing := range finalRelation {
		fmt.Printf("final: %s %v\n", al, ing)
	}

	for _, rawing := range rawIngredients {
		//fmt.Printf("raw: %s\n", rawing)
		if _, ok := finalRelation[rawing]; !ok {
			totalcount++
		}
	}

	fmt.Printf("total: %d\n", totalcount)
	//Part 2 is already did it solving part1. Just ordering alphabetically the
	// allergens. I did it by hand: bqkndvb,zmb,bmrmhm,snhrpv,vflms,bqtvr,qzkjrtl,rkkrx

}

func processLine(s string) {

	tok1 := strings.Split(s, " (contains ")
	cingredients := strings.Split(tok1[0], " ")
	calergens := strings.Split(tok1[1], ", ")

	for _, i := range cingredients {
		rawIngredients = append(rawIngredients, i)
	}

	for _, al := range calergens {
		if ings, ok := alergens[al]; ok {
			alergens[al] = intersec(ings, cingredients)
		} else {
			alergens[al] = cingredients
		}
	}
}

func intersec(a []string, b []string) []string {
	res := []string{}

	for _, v := range a {
		if contains(b, v) {
			res = append(res, v)
		}
	}

	return res
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func removeItem(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
