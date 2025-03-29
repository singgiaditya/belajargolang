package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Singgi"
	names[1] = "Aditya"
	names[2] = "Ramadhan"

	fmt.Println(names[0])

	var values = [3]int{
		90,
		80,
		80,
	}

	fmt.Println(values)
	fmt.Println(len(values))
}
