package main

import "fmt"

func main(){
  number := 40
  var guess int
  tryCount :=1
  fmt.Print("Enter your guess (1-100) :  ")
  fmt.Scanln(&guess)
  fmt.Println(guess)

  for  guess != number{
    if guess > number {
    fmt.Printf("Number is smaller than %v\n Try again : ",guess)
    fmt.Scanln(&guess)
    tryCount++
  }else if number > guess {
    fmt.Printf("Number is greater than %v\n Try again : ",guess)
    fmt.Scanln(&guess)
    tryCount++
  }
  }
  success := ""
  if tryCount>=0 && tryCount<=3 {
    success = "\nEXCELLENT!"
  }else if tryCount<=6{
    success = "\nGood"
  }else {
    success = "\nnot bad"
  }

  fmt.Printf("%v \nMY NUMBER WAS : %v\nYou tried %v times\n",success,number, tryCount)

}
