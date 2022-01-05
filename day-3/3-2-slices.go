package main

import("fmt")
/*
slices can grow and shrink, unlike arrays
The length of a slice is the number of elements it contains. 
The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice. 

*/
func main(){
  numbers := []int{}
  fmt.Println(len(numbers))
  fmt.Println(cap(numbers))
  fmt.Println(numbers)

  names := []string{"gokhan","guliz","fatih"}
  fmt.Println(len(names))
  fmt.Println(cap(names))
  fmt.Println(names)

  var anArray = [5]int{1,2,3,4,5}
  aSlice := anArray[2:4]

  fmt.Printf("aSlice = %v\n", aSlice)
  fmt.Printf("length = %d\n", len(aSlice))
  fmt.Printf("capacity = %d\n", cap(aSlice))


  myslice1 := make([]int, 5, 10)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))
  myslice1[4] = 7
  fmt.Println(myslice1)

  // with omitted capacity
  myslice2 := make([]int, 5) 
  fmt.Printf("myslice2 = %v\n", myslice2)
  fmt.Printf("length = %d\n", len(myslice2))
  fmt.Printf("capacity = %d\n", cap(myslice2))
  myslice2[4] = 1
  fmt.Println(myslice2)


  fmt.Println("\nCHANGE, APPENDS")
  pages := []int{1,2,3,4,5}
  fmt.Printf("pages = %v\n", pages)
  fmt.Printf("length = %d\n", len(pages))
  fmt.Printf("capacity = %d\n", cap(pages))

  pages = append(pages,6,7,8)
  fmt.Printf("pages = %v\n", pages)
  fmt.Printf("length = %d\n", len(pages))
  fmt.Printf("capacity = %d\n", cap(pages))

  fmt.Println("\nAppend One Slice To Another Slice")
  animals1 := []string{"cat","dog","mouse"}
  animals2 := []string{"lion","wolf","bird"}
  animals3 := append(animals1,animals2...)

  fmt.Printf("animals3 : %v\n",animals3)
  fmt.Printf("length : %d\n",len(animals3))
  fmt.Printf("capacity : %d\n",cap(animals3))

  



}