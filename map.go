package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Singgi",
		"age":     "21",
		"address": "Malang",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["title"])
	fmt.Println(len(person))
}
