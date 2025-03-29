package main

import "fmt"

func main() {

	type NoKtp string
	type Married bool

	var noKtpEko NoKtp = "1231358484651320"
	var marriedStatus = false
	fmt.Println(noKtpEko)
	fmt.Println(marriedStatus)
}
