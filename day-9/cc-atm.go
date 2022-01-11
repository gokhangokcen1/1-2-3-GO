package main

import "fmt"

func main() {
	var withdraw int
	fmt.Scan(&withdraw)

	var balance float64
	fmt.Scan(&balance)

	if (withdraw%5 == 0) && (balance >= float64(withdraw)+0.5) {
		fmt.Println(balance - float64(withdraw) - 0.5)
		

	} else {
		fmt.Println(balance)
	}

}
