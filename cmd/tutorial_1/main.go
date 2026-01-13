package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello, World!")

	var intNum int
	fmt.Println("Default value of int:", intNum)

	var floatNum float64 = 12345678.9
	fmt.Println("Default value of float64:", floatNum)

	var floatNum32 float32 = 10.1
	fmt.Println("Default value of float32:", floatNum)

	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println("Result of adding float32 and int32:", result)

	var myString1 string = "Hello, Go!"
	fmt.Println("myString1:", myString1)
	var myString2 string = `This is a raw string literal.
It can span multiple lines.`
	fmt.Println("myString2", myString2)

	// len funtions counts the number of bytes in a string
	// for ASCII characters, number of bytes is equal to number of characters
	// but for non-ASCII characters, number of bytes can be more than number of characters
	fmt.Println(len("test1")) // this will print 5
	fmt.Println(len("你好"))    // this will print 6

	fmt.Println(utf8.RuneCountInString("test1")) // this will print 5
	fmt.Println(utf8.RuneCountInString("你 好"))   // this will print 3

	var myRune rune = 'a'
	fmt.Println("myRune:", myRune) // prints the Unicode code point of 'a', which is 97

	var myBool bool = true
	fmt.Println("myBool:", myBool)

	myVar := "text"
	fmt.Println("myVar:", myVar)

	var1, var2 := 1, 2
	fmt.Println("var1:", var1, "var2:", var2)

	const myConst string = "I am a const"
	// myConst = "New Value" // This will cause a compile-time error
	fmt.Println("myConst:", myConst)

	numerator := 10
	denominator := 0
	divisionResult, remainder, err := intDivision(numerator, denominator)

	// We don't need to write a break statement in each case
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the division is %v with no remainder\n", divisionResult)
	default:
		fmt.Printf("The result of the division is %v with a remainder of %v\n", divisionResult, remainder)
	}

	// conditional switch with multiple expressions in each case
	switch remainder {
	case 0:
		fmt.Printf("The result of the division is %v with no remainder\n", divisionResult)
	case 1, 2, 3:
		fmt.Printf("The result of the division is %v with a small remainder of %v\n", divisionResult, remainder)
	default:
		fmt.Printf("The result of the division is %v with a large remainder of %v\n", divisionResult, remainder)
	}
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}

	return numerator / denominator, numerator % denominator, err
}
