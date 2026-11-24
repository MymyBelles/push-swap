package main

import "testing"

func TestSa(t *testing.T) {
	pileA := []int{2, 1, 3}

	sa(pileA)

	if pileA[0] != 1 || pileA[1] != 2 || pileA[2] != 3 {
		t.Errorf("Résultat incorrect : %v, attendu : [1 2 3]", pileA)
	}
}
func TestSaCasParticuliers(t *testing.T) {
	pileVide := []int{}
	sa(pileVide)

	if len(pileVide) != 0 {
		t.Errorf("La pile doit rester vide")
	}

	pileUnique := []int{5}
	sa(pileUnique)

	if len(pileUnique) != 1 || pileUnique[0] != 5 {
		t.Errorf("La pile doit rester [5]")
	}
}
func TestSb(t *testing.T) {
	pileB := []int{5, 4, 6}

	sb(pileB)

	if pileB[0] != 4 || pileB[1] != 5 || pileB[2] != 6 {
		t.Errorf("Résultat incorrect : %v, attendu : [4 5 6]", pileB)
	}
}

func TestSs(t *testing.T) {
	pileA := []int{2, 1, 3}
	pileB := []int{5, 4, 6}

	ss(pileA, pileB)

	if pileA[0] != 1 || pileA[1] != 2 || pileA[2] != 3 {
		t.Errorf("Pile A incorrecte : %v", pileA)
	}

	if pileB[0] != 4 || pileB[1] != 5 || pileB[2] != 6 {
		t.Errorf("Pile B incorrecte : %v", pileB)
	}
}
