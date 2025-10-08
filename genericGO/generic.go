package main

import "fmt"

type Number interface {
	int | float64 | string
}

func contains[T Number](slc []T, val T) T {
	var zero T
	for _, v := range slc {
		if v == val {
			return v
		}
	}
	return zero
}

func main() {
	slc := []int{10, 20, 30, 40}
	res := contains(slc, 50)
	fmt.Println(res)

	slc2 := []string{"hello", "world"}
	res2 := contains(slc2, "hello")
	fmt.Println(res2)
}
