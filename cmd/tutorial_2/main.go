package main

import "fmt"

func main() {
	fmt.Println("Tutorial 2")

	var intArr [3]int32
	intArr[1] = 1
	intArr[2] = 2
	fmt.Println("intArr[0]:", intArr[0])
	fmt.Println("intArr[1:3] this will print the first and second elements:", intArr[1:3])

	fmt.Println(&intArr[0]) // this will print the memory address of the first element

	intArr2 := [...]int32{10, 20, 30}
	fmt.Println("intArr2:", intArr2)

	// Slices
	var intSlice []int32 = []int32{1, 2, 3}
	fmt.Println("intSlice:", intSlice)
	fmt.Printf("The length,capacity of the array bfr appending %v, %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 4)
	fmt.Println("intSlice after append:", intSlice)
	fmt.Printf("The length,capacity of the array after appending %v, %v\n", len(intSlice), cap(intSlice))

	intSlice1 := []int32{5, 6, 7}
	intSlice2 := append(intSlice, intSlice1...) // using the spread operator
	fmt.Println("intSlice2 after appending intSlice1 to intSlice:", intSlice2)

	// use make to create a slice
	intSlice3 := make([]int32, 5, 10) // length 5, capacity 10
	fmt.Printf("intSlice3 created using make: %v\n", intSlice3)

	// Maps

	// This means we have a map with string keys and uint8 values
	var myMap map[string]uint8 = make(map[string]uint8)
	myMap["key1"] = 1
	myMap["key2"] = 2
	fmt.Println("myMap:", myMap)

	var myMap2 = map[string]uint8{
		"Adam":    30,
		"Bob":     25,
		"Charlie": 35,
	}
	fmt.Println("Adam age:", myMap2["Adam"])
	fmt.Println("Rashed age:", myMap2["Rashed"]) // this will print 0 as Rashed key does not exist

	age, exists := myMap2["Rashed"]
	if exists {
		fmt.Println("Rashed age:", age)
	} else {
		fmt.Println("Rashed key does not exist in the map")
	}

	// delete a value from map
	delete(myMap2, "Bob")
	bobAge, bobExists := myMap2["Bob"]
	if bobExists {
		fmt.Println("Bob age:", bobAge)
	} else {
		fmt.Println("Bob key does not exist in the map")
	}

	// Iterating
	for name := range myMap2 {
		fmt.Println("Name:", name, "Age:", myMap2[name])
	}
	for name, age := range myMap2 {
		fmt.Println("Name:", name, "Age:", age)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
