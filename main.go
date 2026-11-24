package main

import (
	"fmt"
	"os"
)

func main() {
	pileA, valide := RecupererArguments(os.Args[1:])
	pileB := []int{}

	//arguments := strings.Fields(strings.Join(os.Args[1:], " "))

	if !valide {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	pileB = []int{}

	//CHOIX AUTOMATIQUE DE TRI - BLOCK SWITCH
	switch len(pileA) {
	case 2:
		triDeux(pileA)
	case 3:
		pileA = triTrois(pileA)
	case 4:
		pileA, pileB = triQuatre(pileA, pileB)
	case 5:
		pileA, pileB = triCinq(pileA, pileB)
	default:
		//fmt.Println("Le programme ne gère pas les piles de plus de 5 éléments.")
		if len(pileA) > 5 {
			pileA, pileB = triBlocs(pileA, pileB)
		}
	}
}
