package main

import "fmt"

func main() {
	var withdraw int
	fmt.Print("Withdraw : ")
	fmt.Scanln(&withdraw)

	var balance float64
	fmt.Print("Balance : ")
	fmt.Scanln(&balance)

	if (withdraw%5 == 0) && (balance >= float64(withdraw)+0.5) {
		result := balance - float64(withdraw) - 0.5
		fmt.Printf("%v \t %v \t %v \n", withdraw, balance, result)

	} else {
		fmt.Println(balance)
	}

}
