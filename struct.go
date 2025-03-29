/**
- struct adalah sebuah template data yang digunakan untuk menggabungkan nol / lebih tipe data
seperti class karena golang buka oop programming
*/

package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

//struct method
func (customer Customer) sayHello() {
	fmt.Println("Hello, My Name is", customer.Name)
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hi", name)
}

func main() {
	var singgi Customer
	singgi.Name = "Singgi Aditya"
	singgi.Address = "Indonesia"
	singgi.Age = 21

	fairuz := Customer{
		Name:    "Achmad Fairuz",
		Address: "Indonesia",
		Age:     30,
	}

	qolbi := Customer{"Qolbi", "Indonesia", 25}

	fmt.Println(singgi)
	fmt.Println(fairuz)
	fmt.Println(qolbi)

	//pemanggil struct method
	singgi.sayHello()
	singgi.sayHi("kamu")
}
