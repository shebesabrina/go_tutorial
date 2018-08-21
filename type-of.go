package main

import (
    "fmt"
    "reflect"
)

func main() {
    fmt.Println(reflect.TypeOf(1))                      // int
    fmt.Println(reflect.TypeOf(1.5))                    // float64
    fmt.Println(reflect.TypeOf("hello"))                // string
    fmt.Println(reflect.TypeOf(false))                  // bool
    // fmt.Println(reflect.TypeOf(net.IPv4(127, 0, 0, 1))) // net.IPv4
    // fmt.Println(reflect.TypeOf(time.Now()))             // time.Time
}
