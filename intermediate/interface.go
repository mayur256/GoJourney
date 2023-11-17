/**
- Interfaces in go can be thought of a set for method signatures. It is also a type.
- A value of interface type can hold any value that implements those methods.
- In Go, an interface is satisfied implicitly if a type implements all the methods of that interface.
*/
package main

import (
	"fmt"
	"math"
)

// interface definition
type geometry interface {
	area() float64
	perimeter() float64
}

// structure types
type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

// Implementing interface for rectangle type
func (rect rectangle) area() float64 {
	return rect.width * rect.height
}

func (rect rectangle) perimeter() float64 {
	return 2 * (rect.width + rect.height)
}

// Implementing interface for circle type
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return math.Pi * c.radius * 2
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perimeter())
}

func main() {
	rect := rectangle{ width: 5, height: 8 }
	circ := circle{ radius: 7 }

	measure(rect)
	measure(circ)
}
