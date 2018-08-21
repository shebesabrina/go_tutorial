package main

import(
  "fmt"
  // "github.com/treehouse-projects/go-intro/welcome"
  "github.com/golang/example/stringutil"
)
 // func main(){
 //   fmt.Println(welcome.English)
 //   fmt.Println(welcome.Japanese)
 //   fmt.Println(welcome.klingon)
 // }
func main() {
  fmt.Println(stringutil.Reverse("hello"))
}
