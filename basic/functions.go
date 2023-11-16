/**
- Like in all other languages a function in Go is a buidling block and paramount in development
- However in Go a function is treated as a first-class citizen. This behaviour is also seen in JS
- Go also has support of closure for their functions
*/
package main

import (
	"fmt"
	"math"
)

// go allows a function to have multiple returns
func multipleResult() (int, string) {
	return 1, "Hello"
}

// a variadic function is a function that can accept variable number of arguments
func variadicSum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func fibonacci() func() int {
	prev := 0
	next := 1

	return func() int {
		result := prev + next;
		prev = next
		next = result
		return result
	}
}

func main() {
	// function as a value
	hypotenuse := func(a, b float64) float64 {
		return math.Sqrt(a*a + b*b)
	}
	fmt.Println(hypotenuse(5, 12))

	fmt.Println(multipleResult())

	// variadic function calls
	fmt.Println(variadicSum(1,2))
	fmt.Println(variadicSum(1,2,3))

	// print out fibonnaci series with help of closure
	f := fibonacci()
	fmt.Println("Fibonacci series with 10 items")
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
