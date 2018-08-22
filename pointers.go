package main

import "fmt"

func main() {
    myNumber := 2.6
    halve(&myNumber)
    fmt.Println("myNumber in 'main':", myNumber)
    // var aValue float64 = 1.23
    // var aPointer *float64 = &aValue
    // fmt.Println("aPointer:", aPointer) // prints something like "aPointer: 0x1040a128"
    // fmt.Println("*aPointer:", *aPointer) // prints "*aPointer: 1.23"
}

func halve(number *float64)  {
  *number = *number / 2
  fmt.Println("*number in 'halve':", *number)
}
// the * before a pointer gets the value of that pointer
// Typing & before a variable name gives you the address of that variable's value. & can be read aloud as "address of".
// func main()  {
//   car:= Car {
//     Doors:        4,
//     Transmission: "automatic",
//     Odometer:      3600,
//   }
//   describe(&car)
// }
//
// func describe(c *Car)  {
//   fmt.Println("A", c.Doors, "door")
//   fmt.Println(c.Transmission, "car")
//   fmt.Println("with", c.Odometer, "miles")
// }
