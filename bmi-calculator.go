
package main
import "fmt"

func bodyMassIndex() {
    var height float64
    var weight float64
    fmt.Print("Height (example : 185): ")
    fmt.Scanln(&height)
    
    decimalHeight := height/100 // 185 -> 1,85
    
    fmt.Print("\nWeight  : ")
    fmt.Scanln(&weight)
    
    result1 := (decimalHeight*decimalHeight) // (weight/height*height)
    result2 := weight / result1
    
    fmt.Printf("\nYour height : %v\n",decimalHeight)
    fmt.Printf("Your weight : %v\n", weight)
    
    fmt.Printf("BMI : %v\n",result2)
    
    if result2 < 18.5 {
        fmt.Println("You are Underweight")
    } else if (result2>=18.5) && (result2<=24.9) {
        fmt.Println("You are Normal Weight")
    } else if (result2>=25.0) && (result2<=29.9) {
        fmt.Println("You are Pre-obesity")
    } else if (result2>=30.0) && (result2<34.9) {
        fmt.Println("You are Obesity class I")
    } else if (result2>=35.0) && (result2<=39.9) {
        fmt.Println("You are Obesity class II")
    } else if (result2>=40.0)  {
        fmt.Println("You are Obesity class III")
    } 
    
    
}

func main() {
     	bodyMassIndex()
}
