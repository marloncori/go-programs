package main

import (
  "fmt"
  "time"
)

func main() {
  nums := []int{24, 63, 98}
  sum := 0
  
  time.Sleep(1*time.Second)
  fmt.Println("These are the members of NUMS slice: {")
  for x := 0; x < len(nums); x++ {
    fmt.Println(nums[x], " ")
    time.Sleep(1*time.Second)
  }
  fmt.Println("}")
  
 time.Sleep(2*time.Second)
 fmt.Println("And this is sum of all elements:")
  for _, num := range nums {
      sum += num
      time.Sleep(1*time.Second)
  }
  
  fmt.Println("Now let us print the uneven elements:")
  time.Sleep(1*time.Second)
  for i, num := range nums {
        if num % 2 == 1 {
            fmt.Println("index:", i)
            time.Sleep(1*time.Second)
        }
    }
  
}
