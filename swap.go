package main

import "fmt"

func sa(pileA []int) {
	if len(pileA) < 2 {
		return
	}

	pileA[0], pileA[1] = pileA[1], pileA[0]
	fmt.Println("sa")
}

func sb(pileB []int) {
	if len(pileB) < 2 {
		return
	}

	pileB[0], pileB[1] = pileB[1], pileB[0]
	fmt.Println("sb")
}

func ss(pileA []int, pileB []int) {
	if len(pileA) < 2 && len(pileB) < 2 {
		return
	}

	//fmt.Println("ss")
	//sa(pileA)
	if len(pileA) >= 2 {
		pileA[0], pileA[1] = pileA[1], pileA[0]
	}
	//sb(pileB)
	if len(pileB) >= 2 {
		pileB[0], pileB[1] = pileB[1], pileB[0]
	}
	fmt.Println("ss")
}
