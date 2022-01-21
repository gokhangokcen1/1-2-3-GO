package main

import("fmt")

func main(){
  //BOOLEAN
  var flag bool = true
  var luckyNumber int = 3
  var pi float32 = 3.14  
  var name string = "gokhan"

  fmt.Println("Boolean : ",flag)
  fmt.Println("Integer : ",luckyNumber)
  fmt.Println("Float : ",pi)
  fmt.Println("String : ",name)

  var flag1 bool = true  //true
  var flag2 = true //true
  var flag3 bool //false
  flag4 := true //true

  fmt.Println(flag1)
  fmt.Println(flag2)
  fmt.Println(flag3)
  fmt.Println(flag4) 
  flag3 = true
  fmt.Println("New flag3 : ",flag3) //now true

  //INTEGER
  fmt.Print("\n")
  var positive int = 500
  var negative int = -4500
  fmt.Printf("Type: %T, value: %v\n", positive, positive)
  fmt.Printf("Type: %T, value: %v\n", negative, negative)

  var number1 uint = 500
  var number2 uint = 4500
  fmt.Printf("Type: %T, value: %v\n", number1, number1)
  fmt.Printf("Type: %T, value: %v\n", number2, number2)
  /*
  five types of SIGNED integers: 
  int int8 int16 int32 int64

  five types of UNSIGNED integers:
  uint uint8 uint16 uint32 uint64
  */

  //FLOAT ->   float32 float64
  fmt.Print("\n","FLOAT 32","\n")
  var strangeNumber1 float32 = 123.78
  var strangeNumber2 float32 = 3.4e+38
  fmt.Printf("Type: %T, value: %v\n", strangeNumber1, strangeNumber1)
  fmt.Printf("Type: %T, value: %v\n", strangeNumber2, strangeNumber2)
  fmt.Print("\n","FLOAT 64","\n")
  var biggerStrangeNumber float64 = 1.7e+308
  fmt.Printf("Type: %T, value: %v\n", biggerStrangeNumber, biggerStrangeNumber)

  //STRING
  fmt.Print("\n")  
  var txt1 string = "Hello!"
  var txt2 string
  txt3 := "World"

  fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
  fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
  fmt.Printf("Type: %T, value: %v\n", txt3, txt3)
}

