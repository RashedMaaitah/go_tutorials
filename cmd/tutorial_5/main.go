package main

import "fmt"

func main() {
	// var p *int32 = new(int32) // if we didn't initialize p, it would be nil
	// var i int32
	// fmt.Printf("The value p points to is: %v\n", *p)
	// fmt.Printf("The value of i is: %v\n", i)
	// // *p = 10
	// p = &i
	// *p = 1 // this will change i too
	// fmt.Printf("The value p points to is: %v\n", *p)
	// fmt.Printf("The value of i is: %v\n", i)

	// var slice = []int32{1, 2, 3}
	// var sliceCopy = slice
	// sliceCopy[2] = 4
	// fmt.Println(slice)
	// fmt.Println(sliceCopy)

	// var thing1 = [5]float64{1, 2, 3, 4, 5}
	// fmt.Printf("\nThe memory location of the thing1 variable is: %p", &thing1)
	// result := square(thing1)
	// fmt.Printf("\nThe result is: %v\n", result)
	// fmt.Printf("\nThe value of thing1 is: %v\n", thing1)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing1 variable is: %p", &thing1)
	result := square(&thing1)
	fmt.Printf("\nThe result is: %v\n", result)
	fmt.Printf("\nThe value of thing1 is: %v\n", thing1)
}
func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing2 variable is: %p", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2

}

// func square(thing2 [5]float64) [5]float64 {
// 	fmt.Printf("\nThe memory location of the thing1 variable is: %p", &thing2)
// 	for i := range thing2 {
// 		thing2[i] = thing2[i] * thing2[i]
// 	}
// 	return thing2
// }
