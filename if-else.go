package main

import "fmt"

func main() {
    if false {
        fmt.Println("if") // not printed
    } else if true {
        fmt.Println("else if") // printed
    } else {
        fmt.Println("else") // not printed
    }
}


func PlayHand(score int) string {
    var action string
    if score < 17 {
        action = "hit me"
    } else if score > 21 {
        action = "bust"
    } else {
        action = "stand"
    }

    fmt.Println(action)
    return action
}

package golf


//ultimate flow control
func ShotsDescription(par int, shots int) string {
  shotsOverPar := shots - par
    if shotsOverPar == 1 {  // If the shotsOverPar variable is 1, return "bogey".
      return "bogey"
    } else if shotsOverPar == 0 { // If it's 0, return "par".
      return "par"
    } else if shotsOverPar == -1 { // For -1 return "birdie".
      return "birdie"
    } else if shotsOverPar == -2 { // For -2 return "eagle".
      return "eagle"
    } else { // For all other cases, return "no description".
      return "no description"
  }

}
