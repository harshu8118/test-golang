package main

import "fmt"

func main() {
	var f [5]int
	f[0] = 1
	f[1] = 2
	f[2] = 3
	f[3] = 4
	f[4] = 5
	var g = []int{1, 2, 3, 4, 5, 54, 45, 34, 54, 43}
	b := [5]int{2, 4, 6, 9, 10}
	fmt.Printf("%T", b)
	fmt.Printf("%T", g)
	a := []int{2, 4, 6, 9, 10, 7, 5, 34, 345, 23}
	fmt.Printf("%T", a)
	m := []int{42, 43, 44, 45, 46}
	n := []int{47, 48, 49, 50, 51}
	o := []int{44, 45, 46, 47, 48}
	p := []int{43, 44, 45, 46, 47}
	fmt.Println(m, n, o, p)
	fmt.Println(a[1:6])
	fmt.Println(a[2:7])
	fmt.Println(a[3:8])
	fmt.Println(a[4:9])
	var x = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Print(x)
	var y = []int{56, 57, 58, 59, 60}
	x = append(x, y...)
}
