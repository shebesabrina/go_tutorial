// go fmt filename.go will format/auto indent your code.

package main

import (
    "fmt"
    "os"
)

func main() {
    fileInfo, error := os.Stat("existent.txt")
    if error != nil {
        fmt.Println(error)
    } else {
        fmt.Println(fileInfo.Size())
    }
    fileInfo, error = os.Stat("nonexistent.txt")
    if error != nil {
        fmt.Println(error)
    } else {
        fmt.Println(fileInfo.Size())
    }
}


package main

import (
    "fmt"
    "log"
    "math"
)

func main() {
    squareRoot, err := squareRoot(9)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(squareRoot)
}

func squareRoot(x float64) (float64, error) {
    if x < 0 {
        return 0, fmt.Errorf("can't take square root of negative number")
    }
    return math.Sqrt(x), nil
}

// package main
//
// import "fmt"
//
// func main() {
//     fmt.Println(square(3))
//     fmt.Println(add(2, 4))
//     fmt.Println(subtract(10, 3))
// }
//
// func square(number int) int{
//   //add print line above return or else you get an error: fmt.Println("hello")
//     return number * number
// }
//
// func add(a float64, b float64) (sum float64) {
//   return a + b
// }
//
// func subtract(a, b float64) (difference float64){
//     difference = a - b
//     return
// }





// package main
// import "fmt"
//
// func main()  {
//   myFunction()
// }
//  func myFunction()  {
//    fmt.Println("Running myFunction")
//  }
