package main

//import "fmt"

func rra(a []int) []int {
	if len(a) < 2 {
		return a
	}
	dernier := a[len(a)-1]
	//fmt.Println("rra")
	return append([]int{dernier}, a[:len(a)-1]...)
}

func rrb(b []int) []int {
	if len(b) < 2 {
		return b
	}
	dernier := b[len(b)-1]
	//fmt.Println("rrb")
	return append([]int{dernier}, b[:len(b)-1]...)
}

func rrr(a []int, b []int) ([]int, []int) {
	//a = rra(a)
	if len(a) >= 2 {
		dernierA := a[len(a)-1]
		a = append([]int{dernierA}, a[:len(a)-1]...)
	}

	//b = rrb(b)
	if len(b) >= 2 {
		dernierB := b[len(b)-1]
		b = append([]int{dernierB}, b[:len(b)-1]...)
	}
	//b = rrb(b)

	//fmt.Println("rrr")
	return a, b

}
