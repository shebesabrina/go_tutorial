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
