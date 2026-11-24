package main

import (
	"reflect"
	"testing"
)

func TestRa(t *testing.T) {
	pileA := []int{1, 2, 3}

	pileA = ra(pileA)

	if !reflect.DeepEqual(pileA, []int{2, 3, 1}) {
		t.Errorf("Pile A incorrecte : %v", pileA)
	}
}

func TestRb(t *testing.T) {
	pileB := []int{4, 5, 6}

	pileB = rb(pileB)

	if !reflect.DeepEqual(pileB, []int{5, 6, 4}) {
		t.Errorf("Pile B incorrecte : %v", pileB)
	}
}

func TestRr(t *testing.T) {
	pileA := []int{1, 2, 3}
	pileB := []int{4, 5, 6}

	pileA, pileB = rr(pileA, pileB)

	if !reflect.DeepEqual(pileA, []int{2, 3, 1}) {
		t.Errorf("Pile A incorrecte : %v", pileA)
	}

	if !reflect.DeepEqual(pileB, []int{5, 6, 4}) {
		t.Errorf("Pile B incorrecte : %v", pileB)
	}
}
