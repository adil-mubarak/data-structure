package main

import "fmt"

func main() {
	Prices := make(map[string]float64)

	//inserting values
	Prices["laptop"] = 1200.50
	Prices["phone"] = 799.99
	Prices["tablet"] = 450.34

	//Retrive values
	fmt.Println("laptop price:", Prices["laptop"])
	fmt.Println("phone price:", Prices["phone"])
	fmt.Println("tablet price:", Prices["tablet"])

	//update price
	Prices["tablet"] = 999
	fmt.Println("updated tablet price:", Prices["tablet"])

	//delete price
	delete(Prices, "phone")
	fmt.Println("after delete phone price", Prices["phone"])

}
