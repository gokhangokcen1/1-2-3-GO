package main

import "fmt"

func main() {

	var T int
	fmt.Scan(&T)

	for T > 0 {
		var number int
		fmt.Scan(&number) //769
		var addition int
		for number > 0 {
			addition = addition + number%10 //Modulus ex: 769 -> 9
			number = number / 10            //Division ex: 847 -> 84
		}
		fmt.Println(addition)
		T = T - 1 //or t--
	}

}
