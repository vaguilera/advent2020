package main

import "fmt"

var totalcount int

var deck1 = []int{38, 39, 42, 17, 13, 37, 4, 10, 2, 34, 43, 41, 22, 24, 46, 19, 30, 50, 6, 44, 28, 27, 36, 5, 45}
var deck2 = []int{31, 40, 25, 11, 3, 48, 16, 9, 33, 7, 12, 35, 49, 32, 26, 47, 14, 8, 20, 23, 1, 29, 15, 21, 18}

func main() {

	// file, err := os.Open("input.txt")
	// defer file.Close()

	// if err != nil {
	// 	log.Fatalf("failed to open")
	// }

	// scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanLines)

	totalcount = 0
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	processLine(line)
	// }

	totalcards := len(deck1)
	for (len(deck1) > 0) && (len(deck2) > 0) {
		round()
	}
	fmt.Printf("%v\n%v\n", deck1, deck2)

	wdeck := make([]int, totalcards*2)
	if len(deck1) != 0 {
		copy(wdeck, deck1)

	} else {
		copy(wdeck, deck2)
	}
	for i := 0; i < len(wdeck); i++ {
		totalcount += wdeck[i] * (len(wdeck) - i)
	}

	fmt.Printf("total: %d\n", totalcount)
}

func round() {
	a := deck1[0]
	b := deck2[0]

	deck1 = deck1[1:]
	deck2 = deck2[1:]

	if a > b {
		deck1 = append(deck1, []int{a, b}...)
	} else {
		deck2 = append(deck2, []int{b, a}...)
	}

}
