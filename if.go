package main

import "fmt"

func main() {
	var name = "Aditya"

	if name == "Singgi" {
		fmt.Println("Hello Singgi")
	} else if name == "Aditya" {
		fmt.Println("Hello Aditya")
	} else {
		fmt.Println("Hi, Kamu Siapaa?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
