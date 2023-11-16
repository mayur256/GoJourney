/**
- The primary distinction between an array and a slice is that an array is of fixed size, while
- a slice can dynamically grow/shrink throughout it's lifetime
- A slice can be thought of as a view/reference to an array
 */
package main

import (
	"fmt"
	"slices"
)

func main() {
	// slice literal
	fmt.Println([]bool{ true, false, true })

	// using make buitlin to construct a slice
	s := make([]string, 3)
	fmt.Println(s, len(s))

	// construct slice from an array
	var primes = [5]int {2, 3, 5, 7, 11} // fixed length array
	primesSubset := primes[1:4]
	fmt.Println(primes, primesSubset)

	// operations on slice
	primesSubset = append(primesSubset, 11) // append items to the slice
	fmt.Println("After append::", primesSubset)

	s1 := []string {"a", "b", "c"}
	s2 := []string {"a", "b", "c"}
	fmt.Println(slices.Equal(s1, s2))
}
