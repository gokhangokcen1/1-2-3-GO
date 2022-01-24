package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {    //nill : len & cap = 0
		fmt.Println("nil!")
	}
}
