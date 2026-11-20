package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//pileA := []int{}
	pileA := []int{5, 2, 8}
	//pileB := []int{7, 2, 4}
	//pileB := []int{}
	pileB := []int{3, 1, 2}

	for i := 1; i < len(os.Args); i++ {
		nombre, err := strconv.Atoi(os.Args[i])

		if err != nil {
			fmt.Fprint(os.Stderr, "Error")
			return
		}

		for _, valeur := range pileA {
			if valeur == nombre {
				fmt.Fprintln(os.Stderr, "Error")
				return
			}
		}
		pileA = append(pileA, nombre)
	}

	//sa(pileA)
	//sb(pileB)
	//ss(pileA, pileB)
	//pileA, pileB = pa(pileA, pileB)
	//pileA, pileB = pb(pileA, pileB)
	//pileA = ra(pileA)
	//pileB = rb(pileB)
	//pileA, pileB = rr(pileA, pileB)
	//pileA = rra(pileA)
	//pileB = rrb(pileB)
	pileA, pileB = rrr(pileA, pileB)

	//triDeux(pileA)
	//triTrois(pileA)
	//pileA = triDeux(pileA)
	//pileA = triTrois(pileA)
	//pileA, pileB = triQuatre(pileA, pileB)
	//pileA, pileB = triCinq(pileA, pileB)

	fmt.Println(pileA)
	fmt.Println(pileB)
	fmt.Print("Pile A triée :", estTrie(pileA))
}
