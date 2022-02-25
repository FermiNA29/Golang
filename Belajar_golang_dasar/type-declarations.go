package main

import "fmt"

func main() {
	type noKtp string

	var noKtpFermi noKtp = "12345"
	fmt.Println(noKtpFermi)
}
