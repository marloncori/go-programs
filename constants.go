package main

import (
  "fmt"
  "math"
)
const S string = "Jesus is my Lord!"

func main() {
  time.Sleep(1*time.Second)
  fmt.Printf("This is the constant %s:", S)
  
  const N = 50000
  const D = 3e20 / N
  
  fmt.Println("This is the constant D: ", D)
  time.Sleep(1*time.Second)
  
  fmt.Println("Let us convert it to int64:")
  time.Sleep(1*time.Second)
  fmt.Println(int64(D))
  
  fmt.Printf("And this is the sin for the N value (%d):", N)
  time.Sleep(1*time.Second)
  fmt.Println(math.Sin(N))
}
