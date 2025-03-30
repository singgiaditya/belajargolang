package main

import "fmt"

func random() interface{} {
	return "Ups"
}

func main() {
	var result interface{} = random()

	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// var resultInt int = result.(int) //panic
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("nilai string:", value)
	case int:
		fmt.Println("nilai int:", value)
	default:
		fmt.Println("Unknown")
	}

}
