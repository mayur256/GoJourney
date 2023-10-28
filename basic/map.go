package main

import (
	"fmt"
	"strings"
)

// returns the number of occurence of a word in a sentence
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	result := make(map[string]int)
	for i := 0; i < len(words); i++ {
		if count, ok := result[words[i]]; ok {
			result[words[i]] = count + 1
		} else {
			result[words[i]] = 1
		}
	}
	return result
}

func main() {
	// map literal
	mapLiteral := map[string]bool {"true": true, "false": false}
	fmt.Println(mapLiteral)

	// use the male builtin to create and initialise map
	myMap := make(map[string]int)
	
	// len prints the number of key=>value pair present in the map
	fmt.Println(len(myMap))

	// add element to map
	myMap["one"] = 1
	myMap["two"] = 2
	fmt.Println(myMap, len(myMap));

	// get element from map
	one := myMap["two"]
	fmt.Println(one)

	// delete an item from map
	delete(myMap, "two")
	fmt.Println(myMap)

	fmt.Println(WordCount("Mayur is a good boy"))
}
