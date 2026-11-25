package main

import (
	"reflect"
	"testing"
)

// ========================================
// TESTS DES OPERATIONS SWAP
// ========================================

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

// ========================================
// TESTS DES OPERATIONS PUSH
// ========================================

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

// ========================================
// TESTS DES OPERATIONS ROTATE
// ========================================

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

// ========================================
// TESTS DES OPERATIONS REVERSE ROTATE
// ========================================

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

// ========================================
// TESTS DE VERIFICATION DU TRI
// ========================================

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

// ========================================
// TESTS DES ALGORITHMES DE TRI
// ========================================

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

// ========================================
// TESTS DE VALIDATION DES ARGUMENTS
// ========================================

func TestArgumentsValides(t *testing.T) {
	pile, valide := RecupererArguments([]string{"3 2 1"})

	if !valide || !reflect.DeepEqual(pile, []int{3, 2, 1}) {
		t.Errorf("Arguments valides mal traités : %v", pile)
	}
}

func TestArgumentsNegatifs(t *testing.T) {
	pile, valide := RecupererArguments([]string{"-3 0 2"})

	if !valide || !reflect.DeepEqual(pile, []int{-3, 0, 2}) {
		t.Errorf("Nombres négatifs mal traités : %v", pile)
	}
}

func TestArgumentsDoublons(t *testing.T) {
	_, valide := RecupererArguments([]string{"3 2 3"})

	if valide {
		t.Errorf("Les doublons doivent être refusés")
	}
}

func TestArgumentsNonNumeriques(t *testing.T) {
	_, valide := RecupererArguments([]string{"3 abc 1"})

	if valide {
		t.Errorf("Les caractères non numériques doivent être refusés")
	}
}

func TestArgumentVide(t *testing.T) {
	_, valide := RecupererArguments([]string{""})

	if valide {
		t.Errorf("Un argument vide doit être refusé")
	}
}

func TestSansArguments(t *testing.T) {
	pile, valide := RecupererArguments([]string{})

	if !valide || len(pile) != 0 {
		t.Errorf("L'absence d'arguments doit être acceptée")
	}
}
