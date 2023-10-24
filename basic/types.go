/**
- a type of variable signifies the size, and operations that are legal for a value
- Go has following categorical types
- string
- int, int8, int16, int32, int64,
- uint, unit8, uint16, uint32, uint64, uintptr
- byte // alias for uint8
- rune // alias for int32
- float, float64,
- complex64, complex128
*/
package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T | Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T | Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T | Value: %v\n", z, z)
}
