package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("Perulangan ke ", counter)
	// }

	// slice := []string{"singgi", "aditya", "ramadhan"}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }

	// 	for i, value := range slice {
	// 		fmt.Println("index", i, "=", value)
	// 	}
	// }

	// for _, value := range slice {
	// 	fmt.Println(value)
	// }

	person := make(map[string]string)
	person["name"] = "Singgi"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
