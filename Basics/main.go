package main

import (
	"fmt"
	"math"
)

func main() {
	// Summary of 2 and 3, Package fmt
	var a int
	a = 2 + 3
	fmt.Println("Summary of 2 and 3 is", a)
	// Package math
	var b float64
	b = math.Sqrt(4)
	fmt.Println("Square of 4 is", b)
	var c float64
	c = math.Sin(math.Pi / 2)
	fmt.Println("Sin of Pi / 2 is", c)
	var d float64
	d = math.Exp(math.Log(10))
	fmt.Printf("Exponential of log10 is %.f\n", d) //10.000000000000002
	// Function
	fmt.Println(factorial(10))
}

func factorial(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}
