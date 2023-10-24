package main

import "fmt"

func main() {
	// string
	var hello string = "Hello World!"
	fmt.Println(hello);

	// numbera
	var number int = 5
	fmt.Println(number);

	var floatNumber float32 = 3.14
	fmt.Println(floatNumber);

	// boolean
	var truthy bool = true
	var falsy bool // non-initialised value will have zero value for their corresponding types
	fmt.Println(truthy, falsy);

	// Variable assignment shorthand
	g := 9.8
	fmt.Println("Gravity due to acceleration: ", g);
}
