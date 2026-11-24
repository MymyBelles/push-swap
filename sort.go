package main

func triDeux(a []int) {
	if len(a) == 2 && a[0] > a[1] {
		sa(a)
	}
}

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

func triQuatre(a []int, b []int) ([]int, []int) {
	if len(a) != 4 {
		return a, b
	}

	if estTrie(a) {
		return a, b
	}

	min := a[0]

	for _, valeur := range a {
		if valeur < min {
			min = valeur
		}
	}

	positionMin := 0

	for i, valeur := range a {
		if valeur == min {
			positionMin = i
			break
		}
	}
	if positionMin <= 2 {
		for a[0] != min {
			a = ra(a)
		}
	} else {
		for a[0] != min {
			a = rra(a)
		}
	}

	if a[0] == min {
		a, b = pb(a, b)
	}

	a = triTrois(a)

	a, b = pa(a, b)

	return a, b
}

func triCinq(a []int, b []int) ([]int, []int) {
	if len(a) != 5 {
		return a, b
	}
	if estTrie(a) {
		return a, b
	}
	min := a[0]

	for _, valeur := range a {
		if valeur < min {
			min = valeur
		}
	}

	positionMin := 0

	for i, valeur := range a {
		if valeur == min {
			positionMin = i
		}
	}

	if positionMin <= 2 {
		for a[0] != min {
			a = ra(a)
		}
	} else {
		for a[0] != min {
			a = rra(a)
		}
	}
	a, b = pb(a, b)
	a, b = triQuatre(a, b)
	a, b = pa(a, b)

	return a, b
}
