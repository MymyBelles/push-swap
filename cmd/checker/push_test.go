package main

import (
	"reflect"
	"testing"
)

func TestPa(t *testing.T) {
	pileA := []int{2, 3}
	pileB := []int{1, 4}

	pileA, pileB = pa(pileA, pileB)

	if !reflect.DeepEqual(pileA, []int{1, 2, 3}) {
		t.Errorf("Pile A incorrecte : %v", pileA)
	}

	if !reflect.DeepEqual(pileB, []int{4}) {
		t.Errorf("Pile B incorrecte : %v", pileB)
	}
}

func TestPb(t *testing.T) {
	pileA := []int{1, 2, 3}
	pileB := []int{4}

	pileA, pileB = pb(pileA, pileB)

	if !reflect.DeepEqual(pileA, []int{2, 3}) {
		t.Errorf("Pile A incorrecte : %v", pileA)
	}

	if !reflect.DeepEqual(pileB, []int{1, 4}) {
		t.Errorf("Pile B incorrecte : %v", pileB)
	}
}
