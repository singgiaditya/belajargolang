package main

import "fmt"

func sayHelloTo(firstname string, lastname string) {
	fmt.Println("Hello", firstname, lastname)
}

func getHello(name string) string {
	if name == "" {
		return "Hello bro"
	} else {
		return "Hello " + name
	}
}

// multiple return values
func getFullName() (string, string) {
	return "Singgi", "Aditya"
}

// named return values
func getCompleteName() (firstname, middleName, lastName string) {
	firstname = "Singgi"
	middleName = "Aditya"
	lastName = "Ramadhan"

	return
}

//variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	sayHelloTo("Singgi", "Aditya")
	fmt.Println(getHello("Singgi"))
	fmt.Println(getHello(""))

	// firstname, lastname := getFullName()
	firstname, _ := getFullName()

	// fmt.Println(firstname, lastname)
	fmt.Println(firstname)

	namaDepan, namaTengah, namaAkhir := getCompleteName()

	fmt.Println(namaDepan, namaTengah, namaAkhir)

	numberSlice := []int{
		10, 90, 30, 50, 40,
	}

	total := sumAll(numberSlice...)
	fmt.Println(total)

	//variabel yang menyimpan function
	goodBye := getGoodBye

	fmt.Println(goodBye("Singgi"))

	//anonymous function (function yang langsung dimasukkan variabel)
	blacklist := func(name string) bool {
		return name == "root"
	}

	fmt.Println(blacklist("Singgi"))

	fmt.Println(factorialRecursive(5))
}
