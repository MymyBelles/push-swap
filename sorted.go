package main

func estTrie(pileA []int) bool {
	for i := 0; i < len(pileA)-1; i++ {
		if pileA[i] > pileA[i+1] {
			return false
		}
	}
	return true
}
