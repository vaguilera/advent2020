package main

import (
	"fmt"
)

var totalcount int

var deck1 = []int{38, 39, 42, 17, 13, 37, 4, 10, 2, 34, 43, 41, 22, 24, 46, 19, 30, 50, 6, 44, 28, 27, 36, 5, 45}
var deck2 = []int{31, 40, 25, 11, 3, 48, 16, 9, 33, 7, 12, 35, 49, 32, 26, 47, 14, 8, 20, 23, 1, 29, 15, 21, 18}

//var deck1 = []int{9, 2, 6, 3, 1}
//var deck2 = []int{5, 8, 4, 7, 10}

type turn struct {
	deck1, deck2 []int
}

type game struct {
	history []turn
}

func main() {

	totalcount = 0

	maingame := newGame(deck1, deck2)
	// maingame.round()
	// maingame.round()
	// maingame.round()
	// maingame.round()
	// maingame.round()
	// maingame.round()
	// maingame.round()
	// maingame.round()
	// maingame.round()

	win := maingame.resolveGame()
	fmt.Printf("%d - %v\n", win, maingame)

	d1, _ := maingame.cturn()
	for i := 0; i < len(d1); i++ {
		totalcount += d1[i] * (len(d1) - i)
	}

	fmt.Printf("total: %d\n", totalcount)
}

func newGame(deck1 []int, deck2 []int) game {
	var fturn turn
	fturn.deck1 = make([]int, len(deck1))
	fturn.deck2 = make([]int, len(deck2))
	copy(fturn.deck1, deck1)
	copy(fturn.deck2, deck2)
	return game{
		history: []turn{fturn},
	}
}

func (g *game) resolveGame() int {
	d1, d2 := g.cturn()
	for (len(d1) > 0) && (len(d2) > 0) {
		r := g.round()
		//fmt.Printf("return: %d %v\n", r, g.history)
		d1, d2 = g.cturn()
		if r == 2 {
			return 0
		}
	}

	if len(d1) == 0 {
		return 1
	}
	return 0
}

func (g *game) round() int {

	if g.checkHistory() {
		return 2 // win player 1
	}

	d1, d2 := g.cturn()

	a := d1[0]
	b := d2[0]

	ndeck1 := make([]int, len(d1)-1)
	ndeck2 := make([]int, len(d2)-1)
	copy(ndeck1, d1[1:])
	copy(ndeck2, d2[1:])

	winner := 0

	if a <= len(ndeck1) && b <= len(ndeck2) && len(ndeck1) > 0 && len(ndeck2) > 0 {
		sdeck1 := ndeck1[:a]
		sdeck2 := ndeck2[:b]
		n := newGame(sdeck1, sdeck2)
		if n.resolveGame() == 1 {
			winner = 1
		}
		//	fmt.Printf("SUB: %d\n", winner)
	} else {
		if b > a {
			winner = 1
		}
		//fmt.Printf("NORMAL: %d\n", winner)
	}

	if winner == 0 {
		ndeck1 = append(ndeck1, []int{a, b}...)
	} else {
		ndeck2 = append(ndeck2, []int{b, a}...)
	}

	g.history = append(g.history, turn{deck1: ndeck1, deck2: ndeck2})
	//fmt.Printf("%d\n", len(g.history))
	return winner

}

func (g *game) cturn() ([]int, []int) {
	return g.history[len(g.history)-1].deck1, g.history[len(g.history)-1].deck2
}

func equalSlice(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func (g *game) checkHistory() bool {
	if len(g.history) == 1 {
		//fmt.Println("primer juego")
		return false
	}

	d1, d2 := g.cturn()

	for i := 0; i < len(g.history)-1; i++ {
		//fmt.Printf("Check: %v %v %v %v \n", d1, g.history[i].deck1, d2, g.history[i].deck2)
		if equalSlice(g.history[i].deck1, d1) && equalSlice(g.history[i].deck2, d2) {
			return true
		}
	}

	return false
}
