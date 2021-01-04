package main

import (
	"testing"
)

func TestEqualSlice(t *testing.T) {

	var tests = []struct {
		deck1    []int
		deck2    []int
		expected bool
	}{
		{[]int{3, 4}, []int{}, false},
		{[]int{3, 4}, []int{2, 3}, false},
		{[]int{3, 4}, []int{3, 4}, true},
	}

	for _, test := range tests {
		if output := equalSlice(test.deck1, test.deck2); output != test.expected {
			t.Error("Test Failed: {} inputted, {} {} expected, recieved: {}", test.deck1, test.deck2, test.expected, output)
		}
	}

}

func TestCheckHistory(t *testing.T) {
	g := newGame([]int{1, 1, 1}, []int{11, 11, 11})

	g.history = append(g.history, turn{deck1: []int{2, 2, 2}, deck2: []int{22, 22, 22}})
	g.history = append(g.history, turn{deck1: []int{3, 3, 3}, deck2: []int{33, 33, 33}})
	g.history = append(g.history, turn{deck1: []int{4, 4, 4}, deck2: []int{44, 44, 44}})

	r := g.checkHistory()
	if r {
		t.Error("Test Failed: No previous match should be found")
	}

	///fmt.Printf("HIS: %v\n", g.history)
	g.history = append(g.history, turn{deck1: []int{1, 1, 1}, deck2: []int{11, 11, 11}})
	r = g.checkHistory()
	if !r {
		t.Error("Test Failed: No previous match should be found")
	}

}
