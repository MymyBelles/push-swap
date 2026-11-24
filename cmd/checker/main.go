package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	pileA, valide := RecupererArguments(os.Args[1:])

	if !valide {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	if len(pileA) == 0 {
		return
	}

	pileB := []int{}

	reader := bufio.NewReader(os.Stdin)

	for {
		ligne, err := reader.ReadString('\n')

		if err == io.EOF && len(ligne) == 0 {
			break
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}

		instruction := ligne[:len(ligne)-1]

		switch instruction {
		case "sa":
			sa(pileA)
		case "sb":
			sb(pileB)
		case "ss":
			ss(pileA, pileB)
		case "pa":
			pileA, pileB = pa(pileA, pileB)
		case "pb":
			pileA, pileB = pb(pileA, pileB)
		case "ra":
			pileA = ra(pileA)
		case "rb":
			pileB = rb(pileB)
		case "rr":
			pileA, pileB = rr(pileA, pileB)
		case "rra":
			pileA = rra(pileA)
		case "rrb":
			pileB = rrb(pileB)
		case "rrr":
			pileA, pileB = rrr(pileA, pileB)
		default:
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
	}

	if estTrie(pileA) && len(pileB) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
