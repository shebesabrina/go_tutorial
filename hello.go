package main

import (
	"fmt"
	//format package
	"math"
)
// func main() {
// 	fmt.Println("hello, world\n")
// }
 var myNumber = 1.23
 func main(){
	 roundedUp := math.Ceil(myNumber)
	 roundedDown := math.Floor(myNumber)
	 fmt.Println(roundedUp, roundedDown)
 }
