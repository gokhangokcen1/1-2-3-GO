package main

import "fmt"

func main() {

	var T int
	fmt.Scan(&T)

	for T > 0 {
		var N int
		fmt.Scan(&N) 
		var addition int
		for N > 0 {
			addition = addition + N%10 //Modulus ex: 769 -> 9
			N = N / 10            //Division ex: 847 -> 84
		}
		fmt.Println(addition)
		T = T - 1 //or t--
	}

}
