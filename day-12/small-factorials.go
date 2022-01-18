package main

import (
	"fmt"
	"math/big"
)

func main() {
	var testcase, number, index int64
	var ans big.Int
	
	fmt.Scanln(&testcase)
	
	for index = 0; index < testcase; index++ {
	    
		fmt.Scan(&number)
		
		fmt.Println(ans.MulRange(1, number))
	}

}
