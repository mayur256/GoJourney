/**
- Go supports constants on its primitive values like string, number and bool
- Like in most languages constants variables can't be redeclared or reinitialised
*/
package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
