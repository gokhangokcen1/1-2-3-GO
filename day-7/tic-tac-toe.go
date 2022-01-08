package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	var row,column int
	var letter string
	fmt.Println("Row:")
	fmt.Scanln(&row)
	fmt.Println("Row:")
	fmt.Scanln(&column)
	fmt.Println("O-X:")
	fmt.Scanln(&letter)
	fmt.Println(letter)
	
	board[row-1][column-1] = letter
	

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
