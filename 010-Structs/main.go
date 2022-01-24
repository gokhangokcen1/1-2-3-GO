package main

import "fmt"

type Students struct{
  name string
  surname string
  studentNumber int
  classNo int
}

type Vertex struct {
	X int
	Y int
}

func main(){
  
v := Vertex{1, 2}
v.X = 4
fmt.Println(v.X)  
  
p := &v
p.X = 1e9
fmt.Println(v)
  
var student1 Students
var student2 Students
 

student1.name = "Gökhan"
student1.surname = "Gökcen"
student1.studentNumber = 660
student1.classNo = 210

student2.name = "Ramazan"
student2.surname = "Yildirim"
student2.studentNumber = 571
student2.classNo = 202

fmt.Printf("Name : %v\n", student1.name)
fmt.Printf("Surname : %v\n", student1.surname)
fmt.Printf("Student Number : %v\n", student1.studentNumber)
fmt.Printf("Class No : %v\n", student1.classNo)
fmt.Println()
fmt.Printf("Name : %v\n", student2.name)
fmt.Printf("Surname : %v\n", student2.surname)
fmt.Printf("Student Number : %v\n", student2.studentNumber)
fmt.Printf("Class No : %v\n", student2.classNo)
fmt.Println()

printStudent(student1)

printStudent(student2)
}

func printStudent(student Students) {
  fmt.Printf("Name : %v\n", student.name)
  fmt.Printf("Surname : %v\n", student.surname)
  fmt.Printf("Student Number : %v\n", student.studentNumber)
  fmt.Printf("Class No : %v\n", student.classNo)
  fmt.Println()
} 
