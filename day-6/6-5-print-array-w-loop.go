package main

import "fmt"

func main(){
  var cities [5]string
  cities[0] = "Ankara" 
  cities[1] = "New York"
  cities[2] = "London"
  cities[3] = "Tokyo"
  cities[4] = "Bursa"

  for index,city := range cities{
    fmt.Printf("%v\t%v\n",index+1,city)
  }

  var names [5]string
  names[0] = "gokhan" 
  names[1] = "cennet"
  names[2] = "mert"
  names[3] = "salih"
  names[4] = "guliz"

  var ages [5]int
  ages[0] = 20 
  ages[1] = 18
  ages[2] = 20
  ages[3] = 21
  ages[4] = 19

  for i:=0;i<5;i++{
    fmt.Printf("%v, %v yaşındadır.\n",names[i],ages[i])
  }
}