/**
- A method is just a function with a special receiver argument that is placed just before the
- function identifier and follows the func keyword.
- 
*/
package main

import "fmt"

type rectangle struct {
	width, height int
}

// Methods definition
func (rect *rectangle) area() int {
	return rect.width * rect.height
}

func (rect rectangle) perimeter() int {
	return 2 * (rect.width + rect.height)
}

func main() {
	rect := rectangle{height: 5, width: 12}
	fmt.Println("Area: ", rect.area())
	fmt.Println("Perimeter", rect.perimeter())
}