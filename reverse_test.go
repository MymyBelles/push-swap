package main

import (
	"reflect"
	"testing"
)

func TestRra(t *testing.T) {
	pileA := []int{1, 2, 3}

	pileA = rra(pileA)

	if !reflect.DeepEqual(pileA, []int{3, 1, 2}) {
		t.Errorf("Pile A incorrecte : %v", pileA)
	}
}

func TestRrb(t *testing.T) {
	pileB := []int{4, 5, 6}

	pileB = rrb(pileB)

	if !reflect.DeepEqual(pileB, []int{6, 4, 5}) {
		t.Errorf("Pile B incorrecte : %v", pileB)
	}
}

func TestRrr(t *testing.T) {
	pileA := []int{1, 2, 3}
	pileB := []int{4, 5, 6}

	pileA, pileB = rrr(pileA, pileB)

	if !reflect.DeepEqual(pileA, []int{3, 1, 2}) {
		t.Errorf("Pile A incorrecte : %v", pileA)
	}

	if !reflect.DeepEqual(pileB, []int{6, 4, 5}) {
		t.Errorf("Pile B incorrecte : %v", pileB)
	}
}
