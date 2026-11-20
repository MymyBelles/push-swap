package main

import "fmt"

func ra(a []int) []int {
	if len(a) < 2 {
		return a
	}

	premier := a[0]
	a = a[1:]
	a = append(a, premier)

	fmt.Println("ra")

	return a
}

func rb(b []int) []int {
	if len(b) < 2 {
		return b
	}
	premier := b[0]
	fmt.Println("rb")
	return append(b[1:], premier)

	//return append([]int{premier}, b[1:]...)
	//return b
}

func rr(a []int, b []int) ([]int, []int) {
	//a = ra(a)

	if len(a) >= 2 {
		premierA := a[0]
		//a = a[1:]
		a = append(a[1:], premierA)
	}

	//b = rb(b)
	if len(b) >= 2 {
		premierB := b[0]
		b = append(b[1:], premierB)
	}

	fmt.Println("rr")

	return a, b
}
