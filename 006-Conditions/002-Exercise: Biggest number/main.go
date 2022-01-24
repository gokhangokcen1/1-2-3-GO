package main

import "fmt"

func main(){
  x, y, z := 3 , 13, 39

  if (x > y) && (x > z){
    fmt.Printf("%v is the biggest number\n", x)
  } else if (y > x) && (y > z){
    fmt.Printf("%v is the biggest number\n", y)
  } else if (z > x) && (z > y){
    fmt.Printf("%v is the biggest number\n", z)
  }else {
    fmt.Println("Numbers are equal")
  }
}
