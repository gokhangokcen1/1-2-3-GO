package main

import ("fmt")

const age = 20
var number int = 42

const(
luckyNumber int = 3
pi = 3.14
name = "gokhan"
)

func main(){
  number = 41
  fmt.Println(number)
  //if i tried to change age as 21. it will be error.
  fmt.Println(age)

  fmt.Println(luckyNumber)
  fmt.Println(pi)
  fmt.Println(name)


}
