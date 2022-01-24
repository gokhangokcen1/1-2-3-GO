package main

import("fmt")

func main(){
  var names = [2]string{"gokhan","guliz"}
  var books = [...]string{"hamlet", "binbogalar efsanesi","kasagi"}
  
  fmt.Println(names)
  fmt.Println(books)
  fmt.Println(names[0])
  
  
  numbers := [5]int{12,34,45,56,67}
  
  luckyNumbers := [...]int{3,13}

  fmt.Println(numbers)
  
  numbers[2] = 789
  
  fmt.Println(numbers)
  fmt.Println(luckyNumbers)
  fmt.Println(luckyNumbers[1])


  myMarksBiology := [5]int{}
  myMarksPhysics := [5]int{90,80,10}
  myMarksCalculus := [5]int{20,40,50,60,80}
  
  fmt.Println(myMarksBiology)
  fmt.Println(myMarksPhysics)
  fmt.Println(myMarksCalculus)

  ages := [5]int{1:20,4:30}
  fmt.Println(ages)

  fmt.Println(len(numbers))
  fmt.Println(len(books))
}
