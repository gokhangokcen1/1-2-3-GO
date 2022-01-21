package main

import ("fmt")

func main(){
  x , y , z := 3 , 10, 3
  fmt.Println(x < y) //true
  fmt.Println(x > y) //false
  fmt.Println(x != y)//true
  fmt.Println(x == z)//true
  fmt.Println((x < y) && (x == z))//true
  fmt.Println((x > y) || (y < z)) //false

  if 13 > 3{
    fmt.Println("13 bigger than 3")
  }

  if x == z{
    fmt.Println("x = z")
  }

  age := 12

  if age >=18{
    fmt.Println("You can enter.")
  } else{
    fmt.Println("You can't enter")
  }

  name := "okan"
  if (name == "gokhan"){
    fmt.Println("Welcome!")
  } else{
    fmt.Println("INTRODUCE YOURSELF")
  }

  luckyNumber1 := 3
  luckyNumber2 := 13

  if luckyNumber1 > luckyNumber2 {
    fmt.Println("luckyNumber1 is more than luckyNumber2")
  } else if luckyNumber1 < luckyNumber2 {
    fmt.Println("luckyNumber1 is less than luckyNumber2")
  } else {
    fmt.Println("luckyNumber1 and luckyNumber2 are equal")
  }

  myAge := 20
  broAge := 10
  sisAge := 15
  if myAge >= broAge {
    fmt.Println("I'm older than my bro")
    if myAge > sisAge {
      fmt.Println("I also older than my sister")
     }
  } else {
    fmt.Println("I'm the youngest child.")
  }
}