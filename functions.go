package main

import "fmt"

func main() {
    fmt.Println(square(3))
    fmt.Println(add(2, 4))
    fmt.Println(subtract(10, 3))
}

func square(number int) int{
  //add print line above return or else you get an error: fmt.Println("hello") 
    return number * number
}

func add(a float64, b float64) (sum float64) {
  return a + b
}

func subtract(a, b float64) (difference float64){
    difference = a - b
    return
}





// package main
// import "fmt"
//
// func main()  {
//   myFunction()
// }
//  func myFunction()  {
//    fmt.Println("Running myFunction")
//  }
