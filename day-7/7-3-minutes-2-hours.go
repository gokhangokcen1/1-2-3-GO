package main

import "fmt"

func minutes2Hours(){
    var minutes int
    fmt.Print("Minutes : ")
    fmt.Scanln(&minutes)
    fmt.Printf("%v minutes -> %v hour(s) and %v minute(s)\n",minutes,minutes/60,minutes%60)
    
}

func main(){
    minutes2Hours()
}
