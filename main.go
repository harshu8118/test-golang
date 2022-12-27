package main

import (
	"fmt"
	"go/calculator"
)

func main() {
	var a, b int
	fmt.Println("enter the two values of a,b")
	fmt.Scanf("%d,%d", &a, &b)
	fmt.Println("addition of two numbers a,b", calculator.Add(a, b))
	fmt.Println("subtraction of two numbers a,b", calculator.Subtract(a, b))
	fmt.Println("multiplication of two numbers a,b", calculator.Mul(a, b))

	c, d := calculator.Divison(a, b)
	fmt.Printf("quofitient of two numbers a,b = %d", c)
	fmt.Printf("reminder of two numbers a,b =%d", d)

}
