package main

import "sort"

// ========================================
// VERIFICATION DU TRI
// ========================================

// Vérifie si la pile est triée dans l'ordre croissant.
func estTrie(pileA []int) bool {
	for i := 0; i < len(pileA)-1; i++ {
		if pileA[i] > pileA[i+1] {
			return false
		}
	}
	return true
}

// ========================================
// TRI DE DEUX ELEMENTS
// ========================================

func triDeux(a []int) {
	if len(a) == 2 && a[0] > a[1] {
		sa(a)
	}
}

// ========================================
// TRI DE TROIS ELEMENTS
// ========================================

func triTrois(a []int) []int {
	if len(a) != 3 {
		return a
	}

	if estTrie(a) {
		return a
	}

	if a[0] > a[1] && a[1] < a[2] && a[0] < a[2] {
		sa(a)
	} else if a[0] > a[1] && a[1] > a[2] {
		sa(a)
		a = rra(a)
	} else if a[0] > a[1] && a[1] < a[2] && a[0] > a[2] {
		a = ra(a)
	} else if a[0] < a[1] && a[1] > a[2] && a[0] < a[2] {
		sa(a)
		a = ra(a)
	} else if a[0] < a[1] && a[1] > a[2] && a[0] > a[2] {
		a = rra(a)
	}

	return a
}

// ========================================
// TRI DE QUATRE ELEMENTS
// ========================================

func triQuatre(a []int, b []int) ([]int, []int) {
	if len(a) != 4 {
		return a, b
	}

	if estTrie(a) {
		return a, b
	}

	// Recherche de la plus petite valeur.
	min := a[0]

	for _, valeur := range a {
		if valeur < min {
			min = valeur
		}
	}

	// Recherche de la position du minimum.
	positionMin := 0

	for i, valeur := range a {
		if valeur == min {
			positionMin = i
			break
		}
	}

	// Choix du sens de rotation.
	if positionMin <= 2 {
		for a[0] != min {
			a = ra(a)
		}
	} else {
		for a[0] != min {
			a = rra(a)
		}
	}

	// Transfert du minimum dans B.
	if a[0] == min {
		a, b = pb(a, b)
	}

	// Tri des trois éléments restants.
	a = triTrois(a)

	// Retour du minimum dans A.
	a, b = pa(a, b)

	return a, b
}

// ========================================
// TRI DE CINQ ELEMENTS
// ========================================

func triCinq(a []int, b []int) ([]int, []int) {
	if len(a) != 5 {
		return a, b
	}

	if estTrie(a) {
		return a, b
	}

	// Recherche de la plus petite valeur.
	min := a[0]

	for _, valeur := range a {
		if valeur < min {
			min = valeur
		}
	}

	// Recherche de la position du minimum.
	positionMin := 0

	for i, valeur := range a {
		if valeur == min {
			positionMin = i
		}
	}

	// Choix du sens de rotation.
	if positionMin <= 2 {
		for a[0] != min {
			a = ra(a)
		}
	} else {
		for a[0] != min {
			a = rra(a)
		}
	}

	// Transfert du minimum dans B.
	a, b = pb(a, b)

	// Tri des quatre éléments restants.
	a, b = triQuatre(a, b)

	// Retour du minimum dans A.
	a, b = pa(a, b)

	return a, b
}

// ========================================
// TRI PAR BLOCS : PLUS DE CINQ ELEMENTS
// ========================================

func triBlocs(a []int, b []int) ([]int, []int) {

	// Vérification de l'état initial de la pile.
	if estTrie(a) {
		return a, b
	}

	// Définition de la taille des blocs.
	taille := len(a)
	tailleBloc := 20

	if taille > 100 {
		tailleBloc = 55
	}

	// Création et tri d'une copie de la pile A.
	valeurTriees := append([]int{}, a...)
	sort.Ints(valeurTriees)

	// Association de chaque valeur à son rang.
	rangs := make(map[int]int)

	for i, valeur := range valeurTriees {
		rangs[valeur] = i
	}

	limite := tailleBloc

	// PHASE 1 : transfert des éléments de A vers B.
	for len(a) > 0 {

		if rangs[a[0]] < limite {
			a, b = pb(a, b)

			// Placement des petits rangs vers le bas de B.
			if len(b) > 1 && rangs[b[0]] < limite-tailleBloc/2 {
				b = rb(b)
			}

			// Passage au bloc suivant.
			if len(b) == limite {
				limite += tailleBloc
			}

		} else {
			a = ra(a)
		}
	}

	// PHASE 2 : reconstruction de la pile A triée.
	for len(b) > 0 {

		positionMax := 0

		// Recherche du plus grand élément de B.
		for i := 1; i < len(b); i++ {
			if b[i] > b[positionMax] {
				positionMax = i
			}
		}

		// Choix du sens de rotation le plus court.
		if positionMax <= len(b)/2 {

			for positionMax > 0 {
				b = rb(b)
				positionMax--
			}

		} else {

			for positionMax < len(b) {
				b = rrb(b)
				positionMax++
			}
		}

		// Transfert du maximum vers A.
		a, b = pa(a, b)
	}

	return a, b
}
