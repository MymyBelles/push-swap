package main

import (
	"strconv"
	"strings"
)

func validerPile(pile []int) bool {
	// VERIFIER SI LA PILE EST VIDE
	if len(pile) == 0 {
		return false
	}

	// VERIFIER SI LA PILE CONTIENT DES DOUBLONS
	doublons := make(map[int]bool)

	for _, valeur := range pile {
		if doublons[valeur] {
			return false
		}
		doublons[valeur] = true
	}

	return true
}

func RecupererArguments(args []string) ([]int, bool) {
	// VERIFIER LES ARGUMENTS VIDES
	for _, arg := range args {
		if strings.TrimSpace(arg) == "" {
			return nil, false
		}
	}

	arguments := strings.Fields(strings.Join(args, " "))
	pile := []int{}

	for _, argument := range arguments {
		nombre, err := strconv.Atoi(argument)

		if err != nil {
			return nil, false
		}

		pile = append(pile, nombre)
	}

	if len(pile) > 0 && !validerPile(pile) {
		return nil, false
	}

	return pile, true
}
