package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Malang", "Jatim", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Pasuruan"

	fmt.Println(address1)
	fmt.Println(address2)

	// address2 = &Address{"Bandung", "Jabar", "Indonesia"}
	// fmt.Println(address2)

	*address2 = Address{"Bandung", "Jabar", "Indonesia"}
	fmt.Println(address1)

}
