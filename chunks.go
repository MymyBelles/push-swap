package main

import "sort"

// TRIBLOCS TRIE LES PILES DE PLUS DE 5 ELEMENTS EN UTILISANT LES OPERATIONS PUSH-SWAP
func triBlocs(a []int, b []int) ([]int, []int) {

	//PERMET DE VERIFIER SI LA PILE A EST DEJA TRIEE
	if estTrie(a) {
		return a, b
	}

	taille := len(a)

	tailleBloc := 20

	if taille > 100 {
		//tailleBloc = 45
		//tailleBloc = 60
		//tailleBloc = 50
		//tailleBloc = 40
		//tailleBloc = 50
		tailleBloc = 55
	}

	valeurTriees := append([]int{}, a...)

	//COPY DE LA PILE A DANS UNE NOUVELLE PILE POUR LA TRIER
	sort.Ints(valeurTriees)

	//PERMET D'ASSOCIER CHAQUE NB À SON RANG DANS LA PILE TRIEE
	rangs := make(map[int]int)

	for i, valeur := range valeurTriees {
		rangs[valeur] = i
	}

	limite := tailleBloc

	//PERMET DE TRANFERER LES ELEMENTS DE A VERS B
	for len(a) > 0 {

		if rangs[a[0]] < limite {
			a, b = pb(a, b)

			// PERMET DE PASSER AU BLOC SUIVANT
			if len(b) > 1 && rangs[b[0]] < limite-tailleBloc/2 {
				b = rb(b)
			}
			if len(b) == limite {
				limite += tailleBloc
			}

		} else {
			a = ra(a)
		}
	}

	for len(b) > 0 {

		positionMax := 0

		for i := 1; i < len(b); i++ {
			if b[i] > b[positionMax] {
				positionMax = i
			}
		}

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

		a, b = pa(a, b)
	}

	//_ = taille
	//_ = tailleBloc
	//_ = limite
	//_ = valeurTriees

	return a, b
}
