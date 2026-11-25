package main

// ========================================
// SWAP : échange des deux premiers éléments
// ========================================

func sa(pileA []int) {
	if len(pileA) < 2 {
		return
	}

	pileA[0], pileA[1] = pileA[1], pileA[0]
}

func sb(pileB []int) {
	if len(pileB) < 2 {
		return
	}

	pileB[0], pileB[1] = pileB[1], pileB[0]
}

func ss(pileA []int, pileB []int) {
	if len(pileA) < 2 && len(pileB) < 2 {
		return
	}

	if len(pileA) >= 2 {
		pileA[0], pileA[1] = pileA[1], pileA[0]
	}

	if len(pileB) >= 2 {
		pileB[0], pileB[1] = pileB[1], pileB[0]
	}
}

// ========================================
// PUSH : transfert entre les deux piles
// ========================================

func pa(a []int, b []int) ([]int, []int) {
	if len(b) == 0 {
		return a, b
	}

	sommet := b[0]
	a = append([]int{sommet}, a...)
	b = b[1:]

	return a, b
}

func pb(a []int, b []int) ([]int, []int) {
	if len(a) == 0 {
		return a, b
	}

	sommet := a[0]
	b = append([]int{sommet}, b...)
	a = a[1:]

	return a, b
}

// ========================================
// ROTATE : rotation vers le haut
// ========================================

func ra(a []int) []int {
	if len(a) < 2 {
		return a
	}

	premier := a[0]
	a = a[1:]
	a = append(a, premier)

	return a
}

func rb(b []int) []int {
	if len(b) < 2 {
		return b
	}

	premier := b[0]

	return append(b[1:], premier)
}

func rr(a []int, b []int) ([]int, []int) {
	if len(a) >= 2 {
		premierA := a[0]
		a = append(a[1:], premierA)
	}

	if len(b) >= 2 {
		premierB := b[0]
		b = append(b[1:], premierB)
	}

	return a, b
}

// ========================================
// REVERSE ROTATE : rotation vers le bas
// ========================================

func rra(a []int) []int {
	if len(a) < 2 {
		return a
	}

	dernier := a[len(a)-1]

	return append([]int{dernier}, a[:len(a)-1]...)
}

func rrb(b []int) []int {
	if len(b) < 2 {
		return b
	}

	dernier := b[len(b)-1]

	return append([]int{dernier}, b[:len(b)-1]...)
}

func rrr(a []int, b []int) ([]int, []int) {
	if len(a) >= 2 {
		dernierA := a[len(a)-1]
		a = append([]int{dernierA}, a[:len(a)-1]...)
	}

	if len(b) >= 2 {
		dernierB := b[len(b)-1]
		b = append([]int{dernierB}, b[:len(b)-1]...)
	}

	return a, b
}
