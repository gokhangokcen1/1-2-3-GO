package main

import "fmt"

func main(){
  for a:=0; a < 5; a++ {
  fmt.Printf("a : %v\n", a)
  }

  fmt.Println()

  for b:=0; b <= 100; b+=10 {
  fmt.Printf("b : %v\n", b)
  }

  fmt.Println()

  for c:=0; c < 5; c++ {
  if c == 3 {
    continue
  }
  fmt.Printf("c : %v\n", c)
  }

  fmt.Println()
  
  for d:=0; d < 5; d++ {
  if d == 3 {
    break
  }
  fmt.Printf("d : %v\n", d)
  }
  
  fmt.Println("NESTED\n")

  adj := [2]string{"big", "tasty"}
  fruits := [3]string{"apple", "orange", "banana"}
  for i:=0; i < len(adj); i++ {
    for j:=0; j < len(fruits); j++ {
      fmt.Println(adj[i],fruits[j])
    }
  }

  fmt.Println()

  names := [3]string{"tolga", "okan", "gokhan"}
  for index, value := range names {
     fmt.Printf("%v\t%v\n", index+1, value)
  }

  fmt.Println()

  for _, value := range names{
    fmt.Printf("%v\n",value)
  }

  fmt.Println()

  for index, _ := range names{
    fmt.Printf("%v\t%v\n",index,index+1)
  }



}
