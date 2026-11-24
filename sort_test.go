package main

import (
	"reflect"
	"testing"
)

func TestTriDeux(t *testing.T) {
	pileA := []int{2, 1}

	triDeux(pileA)

	if !reflect.DeepEqual(pileA, []int{1, 2}) {
		t.Errorf("Tri incorrect : %v", pileA)
	}
}

func TestTriTrois(t *testing.T) {
	pileA := []int{3, 1, 2}

	pileA = triTrois(pileA)

	if !reflect.DeepEqual(pileA, []int{1, 2, 3}) {
		t.Errorf("Tri incorrect : %v", pileA)
	}
}

func TestTriQuatre(t *testing.T) {
	pileA := []int{4, 2, 1, 3}
	pileB := []int{}

	pileA, pileB = triQuatre(pileA, pileB)

	if !reflect.DeepEqual(pileA, []int{1, 2, 3, 4}) || len(pileB) != 0 {
		t.Errorf("Tri incorrect : A=%v, B=%v", pileA, pileB)
	}
}

func TestTriCinq(t *testing.T) {
	pileA := []int{5, 2, 4, 1, 3}
	pileB := []int{}

	pileA, pileB = triCinq(pileA, pileB)

	if !reflect.DeepEqual(pileA, []int{1, 2, 3, 4, 5}) || len(pileB) != 0 {
		t.Errorf("Tri incorrect : A=%v, B=%v", pileA, pileB)
	}
}

func TestTriBlocs(t *testing.T) {
	pileA := []int{10, 3, 7, 1, 9, 2, 8, 4, 6, 5}
	pileB := []int{}

	pileA, pileB = triBlocs(pileA, pileB)

	attendu := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if !reflect.DeepEqual(pileA, attendu) || len(pileB) != 0 {
		t.Errorf("Tri incorrect : A=%v, B=%v", pileA, pileB)
	}
}
