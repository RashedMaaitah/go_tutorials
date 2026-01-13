package main

import (
	"fmt"
)

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
	//owner if we do this we can access the fields of owner directly from gasEngine
	// so it will be gasEngine.name instead of gasEngine.ownerInfo.name
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	// e.gallons = 100 // This won't affect the original struct, because e is a copy of the original gasEngine
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

type owner struct {
	name string
}

func canMakeIt(e engine, miles uint8) {
	if e.milesLeft() >= miles {
		fmt.Println("You can make it!")
	} else {
		fmt.Println("You can't make it!")
	}
}

func main() {
	fmt.Println("Tutorial 4")

	// var myEngine gasEngine
	// fmt.Println("My engine:", myEngine) // zero values for struct fields

	myEngine := gasEngine{mpg: 25, gallons: 15, ownerInfo: owner{name: "John"}} // we can also ommit the fields names if we provide values in order
	// we can also set the value directly like this:
	// myEngine.mpg = 25
	// myEngine.gallons = 15
	fmt.Printf("My engine: %+v\n", myEngine)

	// Annonymous struct
	myStruct := struct {
		field1 string
		field2 int
	}{field1: "Hello", field2: 42}
	fmt.Printf("My annonymous struct: %+v\n", myStruct)

	miles := myEngine.milesLeft()
	fmt.Printf("Miles left: %v\n", miles)
	fmt.Printf("Gallons left: %v\n", myEngine.gallons)

	/**
	 * Interfaces
	 * An interface is a set of method signatures
	 * A type implements an interface by implementing its methods
	 * There is no explicit declaration of intent to implement an interface
	 * If a type implements all the methods of an interface, it implements the interface
	 * This is called structural typing
	 * Interfaces are used to define behavior
	 * They allow us to write functions that can work with different types that implement the same interface
	**/
	/**
	* a method signature is the method name and the parameters and return types
	* e.g. milesLeft() uint8 is a method signature
	* what if the params had differnt types?
	* e.g. milesLeft(distance uint8) bool
	* then the method signature would be different
	* so the parameters types and return types are part of the method signature
	**/
}
