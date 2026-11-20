package main

import "fmt"

func pa(a []int, b []int) ([]int, []int) {
	if len(b) == 0 {
		return a, b
	}

	sommet := b[0]
	//fmt.Println("Sommet de B :", sommet)
	a = append([]int{sommet}, a...)
	b = b[1:]

	fmt.Println("pa")

	return a, b
}

func pb(a []int, b []int) ([]int, []int) {
	if len(a) == 0 {
		return a, b
	}

	sommet := a[0]
	b = append([]int{sommet}, b...)
	a = a[1:]

	//fmt.Println("Sommet de A :", sommet)
	fmt.Println("pb")

	return a, b
}
