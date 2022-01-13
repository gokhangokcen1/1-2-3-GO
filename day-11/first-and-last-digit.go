package main

import "fmt"

func main() {
	var T int
	fmt.Scanln(&T) //test case number

	var number int
	for i := 0; i < T; i++ {
		fmt.Scanln(&number) //input first number

		lastNumber := number % 10 //last number modulus 10   

		for number >= 10 {    //if first number bigger or equal than 10
			number = number / 10 //first number is first number division 10 now
		}
		result := lastNumber + number  
		fmt.Printf("Result : %v\n",result)
	}
}
