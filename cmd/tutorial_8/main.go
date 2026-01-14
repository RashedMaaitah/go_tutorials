package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwh   float32
	mpkwh float32
}

// Generics with struct type constraints
type car[T gasEngine | electricEngine] struct {
	brand  string
	engine T
}

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float64
	Amount int
}

func main() {

	myCar := car[gasEngine]{brand: "Toyota", engine: gasEngine{gallons: 32, mpg: 15}}
	fmt.Printf("Car Brand: %s, Engine: %+v\n", myCar.brand, myCar.engine)

	contacts, cerr := loadJSON[contactInfo]("contacts.json")
	purchases, perr := loadJSON[purchaseInfo]("purchases.json")
	if cerr != nil {
		fmt.Println("Error loading contacts:", cerr)
		return
	}
	if perr != nil {
		fmt.Println("Error loading purchases:", perr)
		return
	}
	fmt.Println("Contacts:")
	for _, contact := range contacts {
		fmt.Printf("Name: %s, Email: %s\n", contact.Name, contact.Email)
	}
	fmt.Println("\nPurchases:")
	for _, purchase := range purchases {
		fmt.Printf("Name: %s, Price: %.2f, Amount: %d\n", purchase.Name, purchase.Price, purchase.Amount)
	}
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) ([]T, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read file %s: %w", filePath, err)
	}
	var result = []T{}
	err = json.Unmarshal(file, &result)
	if err != nil {
		panic(err)
	}
	return result, nil
}

// func main() {
// 	var intSlice = []int{1, 2, 3, 4, 5}
// 	fmt.Println(sumSlice(intSlice))

// 	var floatSlice = []float64{1.1, 2.2, 3.3, 4.4, 5.5}
// 	fmt.Println(sumSlice(floatSlice))

// 	var emptySlice = []string{}
// 	fmt.Println(isEmpty(emptySlice))
// }

// func sumSlice[T int | float32 | float64](s []T) T {
// 	var sum T
// 	for _, v := range s {
// 		sum += v
// 	}
// 	return sum
// }

// func isEmpty[T any](s []T) bool {
// 	return len(s) == 0
// }
