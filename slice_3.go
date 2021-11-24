package main
import (
  "fmt"
  "time"
)

func main() {
  y := make([]int, 4)
  fmt.Println(y)
  //this will be a slice of zeros
  time.Sleep(1*time.Second)
  fmt.Printf("Type of y: %T", y)
  
  y := append(y, 11)
  time.Sleep(1*time.Second)
  fmt.Println("The new Y slice is = {")
  for x := range y {
    fmt.Printf("%d", y[x], " ")
  }
  fmt.Println("}")
  
  var evens []int = []int{2, 6, 10, 14, 18}
  fmt.Println("The new slice called \"evens\" is: {")
  for i := 0; i < len(evens); i++ {
    fmt.Print(evens[i], " ")
    time.Sleep(1*time.Second)
   }
  fmt.Println("}")
  
  time.Sleep(1*time.Second)
  fmt.Println("Now let us print sorted members of evens slice:")
  for i, elem := range evens {
    for j, elem2 := range evens {
      if elem == elem2 && i > j {
        time.Sleep(1*time.Second)
        fmt.Printf("Sorted element of EVENS: %d", elem)
        time.Sleep(1*time.Second)
      }
    }
  }
}
