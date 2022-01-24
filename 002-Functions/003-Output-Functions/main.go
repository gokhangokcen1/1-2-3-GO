package main

import("fmt")
/*
Print() : default format no space
Printf() : given formatting
Println() : with \n new line
*/
/*
Printf:
%v value  %#v go-syntax format       
%T type   %%  just % sign
 */
func main(){
  var name, surname = "gokhan","gokcen" 
  fmt.Print(name,surname) //gokhangokcen
  fmt.Print("\n")
  fmt.Print(name, "\n")
  fmt.Print(surname, "\n")

  var luckyNumber,age = 3,20
  fmt.Print(luckyNumber,age,"\n") //3 20 seperate

  fmt.Println(name,surname) //gokhan gokcen\n
  fmt.Printf("name has value: %v and type: %T\n",name,name)
  fmt.Printf("name has value: %v and type: %T\n",luckyNumber,luckyNumber)

  fmt.Print("\n")
  var number = 3.14
  var newYear = "Happy New Year"

  fmt.Printf("%v\n", number)
  fmt.Printf("%#v\n", number)
  fmt.Printf("%v%%\n", number)
  fmt.Printf("%T\n", number)

  fmt.Printf("%v\n", newYear)
  fmt.Printf("%#v\n", newYear)
  fmt.Printf("%T\n", newYear) 

  /*
  Integer Formatting Verbs
  %b 	Base 2
  %d 	Base 10
  %+d 	Base 10 and always show sign
  %o 	Base 8
  %O 	Base 8, with leading 0o
  %x 	Base 16, lowercase
  %X 	Base 16, uppercase
  %#x 	Base 16, with leading 0x
  %4d 	Pad with spaces (width 4, right justified)
  %-4d 	Pad with spaces (width 4, left justified)
  %04d 	Pad with zeroes (width 4
  */
  fmt.Print("\n")
  var secondLuckyNumber = 13
  
  fmt.Printf("%b\n", secondLuckyNumber)
  fmt.Printf("%d\n", secondLuckyNumber)
  fmt.Printf("%+d\n", secondLuckyNumber)
  fmt.Printf("%o\n", secondLuckyNumber)
  fmt.Printf("%O\n", secondLuckyNumber)
  fmt.Printf("%x\n", secondLuckyNumber)
  fmt.Printf("%X\n", secondLuckyNumber)
  fmt.Printf("%#x\n", secondLuckyNumber)
  fmt.Printf("%4d\n", secondLuckyNumber)
  fmt.Printf("%-4d\n", secondLuckyNumber)
  fmt.Printf("%04d\n", secondLuckyNumber)
}
