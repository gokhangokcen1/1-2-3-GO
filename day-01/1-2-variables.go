package main
/*

    int- stores integers (whole numbers), such as 123 or -123
    float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
    string - stores text, such as "Hello World". String values are surrounded by double quotes
    bool- stores values with two states: true or false

*/
import (
  "fmt"
)

var okanAge int
var tolgaAge int = 31
var gulizAge int = 20 

func ages(){
  okanAge = 27
  fmt.Println(okanAge,tolgaAge,gulizAge)

}

func main() {
  var name string
  name = "gokhan"
  fmt.Println(name)

  surname := "gokcen"
  fmt.Println(surname)

  var age int
  age = 20
  fmt.Println(age)

  var luckyNumber int
  luckyNumber = 3
  fmt.Println(luckyNumber)

  var x string //""
  var y int // 0
  var z bool // false
  fmt.Println(x,y,z)


  /*
  var
  Can be used inside and outside of functions
  Variable declaration and value assignment can be done separately
  
  := 
  Can only be used inside functions
  Variable declaration and value assignment cannot be done separately (must be done in the same line)*/
  ages()
  
  var a,b,c,d int = 1,2,3,4
  fmt.Println(a,b,c,d)


  var first,thirtyTwo,nine string = "adana", "isparta", "aydin"
  fmt.Println(first,thirtyTwo,nine)

  var g,o = 17,"hey"
  var t,f = 37, "mate!"
  fmt.Println(g)
  fmt.Println(o)
  fmt.Println(t)
  fmt.Println(f)

}
