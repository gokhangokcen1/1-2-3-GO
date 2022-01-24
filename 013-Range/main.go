package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for index, variable := range pow {
		fmt.Printf("2**%d = %d\n", index, variable)
	}
  
  cow := make([]int, 10)
	for i := range cow {
		cow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range cow {
		fmt.Printf("%d\n", value)
	}
}
