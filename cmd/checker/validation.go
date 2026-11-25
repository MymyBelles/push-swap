package main

import (
	"strconv"
	"strings"
)

// ========================================
// VALIDATION DE LA PILE
// ========================================

func validerPile(pile []int) bool {
	if len(pile) == 0 {
		return false
	}

	doublons := make(map[int]bool)

	for _, valeur := range pile {
		if doublons[valeur] {
			return false
		}
		doublons[valeur] = true
	}

	return true
}

// ========================================
// VALIDATION DES ARGUMENTS
// ========================================

func RecupererArguments(args []string) ([]int, bool) {
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

// ========================================
// VERIFICATION DU TRI
// ========================================

func estTrie(pileA []int) bool {
	for i := 0; i < len(pileA)-1; i++ {
		if pileA[i] > pileA[i+1] {
			return false
		}
	}

	return true
}
