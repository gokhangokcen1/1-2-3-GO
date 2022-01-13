package main

import "fmt"
/*
 A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns. 
*/


func helloWorld(){
  defer fmt.Println("world") 
  fmt.Println("gokhan")
  fmt.Println("web application security")
  fmt.Println("hello")

}

func counting(){

  fmt.Println()

  fmt.Println("Counting...")
  for i:=10;i>0;i--{
    defer fmt.Println(i)
  }
  fmt.Println("Counting done!")

}
func main(){
  helloWorld()
  counting()

}
