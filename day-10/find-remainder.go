package main

import "fmt"

func main(){
var T,A,B int 
fmt.Scanln(&T) 

for i:=0;i<T;i++{
    fmt.Scanln(&A,&B)
    fmt.Println(A%B)
}
}
