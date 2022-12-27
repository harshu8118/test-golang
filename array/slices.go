package main

import "fmt"

func main() {
	var f [5]int
	f[0] = 1
	f[1] = 2
	f[2] = 3
	f[3] = 4
	f[4] = 5
	var y = []int{1, 2, 3, 4, 5, 54, 45, 34, 54, 43}
	b := [5]int{2, 4, 6, 9, 10}
	fmt.Printf("%T", b)
	fmt.Printf("%T", y)
	a := []int{2, 4, 6, 9, 10, 7, 5, 34, 345, 23}
	fmt.Printf("%T", a)
	m := []int{42, 43, 44, 45, 46}
	n := []int{47, 48, 49, 50, 51}
	o := []int{44, 45, 46, 47, 48}
	p := []int{43, 44, 45, 46, 47}
	fmt.Println(m, n, o, p)
	fmt.Println(y[1:6])
	fmt.Println(y[2:7])
	fmt.Println(y[3:8])
	fmt.Println(y[4:9])
	var x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

}
