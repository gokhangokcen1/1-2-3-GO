package main

import "fmt"

func helloWorld(){
  fmt.Println("hello world!")
}

func hey(name string){
  fmt.Printf("Hello %v\n",name)
}

func information(name string,age int){
  fmt.Printf("%v is %v years old.\n",name,age)
}

func addition(x int, y int) int{
  return x + y
}

func multp(x int, y int) (result int){
  result = x * y
  return  
}


func sub(x int, y int) (result int){
  result = x - y
  return result
}

func yearlyMoney(name string, money int) (text string, result int){
  text = name + " earned $"
  result = money * 12
  return
}

func counter(x int) int {
  if x == 11 {
    return 0
  }
  fmt.Println(x)
  return counter(x + 1)
}

func factorial_recursion(x float64) (y float64) {
  if x > 0 {
    y = x * factorial_recursion(x-1)
  } else {
    y = 1
  }
  return
}


func main(){
  helloWorld()
  hey("gokhan")
  information("Tolga", 32)
  fmt.Println(addition(3,13))
  fmt.Println(multp(3,5))
  fmt.Println(sub(17,3))
  
  total := addition(3,13)
  fmt.Println(total)

  fmt.Println(yearlyMoney("gökhan", 30000))

  _, b := yearlyMoney( " earned $", 30000)
  fmt.Println(b) 

  a, _ := yearlyMoney( "gökhan", 30000)
  fmt.Println(a) 

  counter(1)

  fmt.Println(factorial_recursion(4))
}
