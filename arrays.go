package main

import "fmt"

func main() {
  months := [3]string{"Apr", "May", "Jun"}
  salesByMonth := [3]float64{1710.26, 2245.97, 3032.40}
  // fmt.Println(months[0], salesByMonth[0])
  // fmt.Println(months[1], salesByMonth[1])
  // fmt.Println(months[2], salesByMonth[2])

  for i, month := range months {
  fmt.Println(month, salesByMonth[i])
}
//
// for _, month := range months { this prints just the months. Need to hash out the salesByMonth
//   fmt.Println(month)
// }


  // for i := 0; i < len(months); i++ {
  //   fmt.Println(months[i], salesByMonth[i])
  // }

    // var months [3]string
    // months[0] = "Apr"
    // months[1] = "May"
    // months[2] = "Jun"
    // var salesByMonth [3]float64
    // salesByMonth[0] = 1710.26
    // salesByMonth[1] = 2245.97
    // salesByMonth[2] = 3032.40
    // fmt.Println(months[0], salesByMonth[0])
    // fmt.Println(months[1], salesByMonth[1])
    // fmt.Println(months[2], salesByMonth[2])
}
