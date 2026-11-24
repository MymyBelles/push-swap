package main

import "testing"

func TestPileTriee(t *testing.T) {
	pileA := []int{1, 2, 3, 4, 5}

	if !estTrie(pileA) {
		t.Errorf("La pile devrait être triée")
	}
}

func TestPileNonTriee(t *testing.T) {
	pileA := []int{3, 1, 5, 2, 4}

	if estTrie(pileA) {
		t.Errorf("La pile ne devrait pas être triée")
	}
}

func TestPileVide(t *testing.T) {
	pileA := []int{}

	if !estTrie(pileA) {
		t.Errorf("Une pile vide doit être considérée comme triée")
	}
}

func TestPileUnique(t *testing.T) {
	pileA := []int{5}

	if !estTrie(pileA) {
		t.Errorf("Une pile avec un seul élément doit être triée")
	}
}
