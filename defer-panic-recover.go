/**
defer adalah sebuah function yang dipanggil setelah function lain
walaupun function sebelumnya error

panic adalah function yang dijalakan ketika error dan defer akan tetap dijalankan tetapi aplikasi akan berhenti

recover adalah function yang akan menangkap message dari panic
*/

package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
	message := recover()
	if message != nil {
		fmt.Println("Terjadi Error", message)
	}
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
}

func testPanic(error bool) {
	defer logging()
	if error {
		panic("Aplikasi Error")
	}
}

func main() {
	// runApplication(10)
	testPanic(false)
	fmt.Println("Singgi")
}
