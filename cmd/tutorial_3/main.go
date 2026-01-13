package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Turorial 3")

	/**
	 * Strings are a sequence of bytes in Golang
	 * They are immutable
	 * They can contain any data, including binary data
	 * They are UTF-8 encoded by default
	 * So a keynote to take is that string are a sequence of bytes, not characters
	 * Using a for range loop will do extra work to decode the UTF-8 bytes into runes (Unicode code points)
	 * meaning that for example é is actually two bytes in UTF-8, but one rune
	 * That is why when we take the len of a string with special characters, it may be longer than expected
	 * because its counting bytes, not characters
	**/
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Println("First byte of myString:", indexed)
	fmt.Printf("%v, %T", indexed, indexed) // the %T will print the type of indexed
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("Length of 'myString' is %v\n", len(myString))

	// Usign Go built in strings
	var strSlice = []string{"Hello", "World", "from", "Go"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i] + " ")
	}
	finalString := strBuilder.String()
	fmt.Println("Final string:", finalString)
}
