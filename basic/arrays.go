package main

import "fmt"

func main() {
	var friends[3] string
	friends[0] = "John"
	friends[1] = "Alice"
	friends[2] = "Mark"

	fmt.Println(friends, len(friends))

	// looping for array
	for i:= 0; i < len(friends); i++ {
		fmt.Println(friends[i])
	}
}
