package main

import (
  "fmt"
  "os"
  "time"
)

func plus(x int, y int) int {
  sum := x + y
  return sum
}

func addThree(a int, b int, c int) int {
  new_sum := a + b + c
  return new_sum
}

func mult_divide(m int, n int) (int, int) {
   mul := m * n
    if m > n {
     div := m / n
    } else
     div := n / m
  return mul, div
}

func main() {
  
  choice := os.Args[1]
  if choice == "1" {
    result := plus(int(os.Args[2]), int(os.Args[3]))
    fmt.Printf("The result of adding %s to %s equals:\n", os.Args[2], os.Args[3])
    
    time.Sleep(1*time.Second)
    fmt.Println("\n", os.Args[2], " + ", os.Args[3], " = ", result)
    
  } else if choice == "2" {
    res := addThree(int(os.Args[2]), int(os.Args[3]), int(os.Args[4]))
    fmt.Printf("The result of adding %s and %s to %s equals:\n", os.Args[2], os.Args[3], os.Args[4])
    
    time.Sleep(1*time.Second)
    fmt.Println("\n", os.Args[2], " + ", os.Args[3], " + ", os.Args[4], " = ", res)
    
  } else {
    result1, result2 := mult_divide(int(os.Args[2]), int(os.Args[3]))
    fmt.Printf("The result of multiplying, dividing %s and %s equals:\n", os.Args[2], os.Args[3])
    
    time.Sleep(1*time.Second)
    fmt.Println("\n", os.Args[2], " * ", os.Args[3], " = ", result1)
    
    if os.Args[2] > os.Args[3] {
      time.Sleep(1*time.Second)
      fmt.Println("\n", os.Args[2], " / ", os.Args[3], " = ", result2)
    } else {
      time.Sleep(1*time.Second)
      fmt.Println("\n", os.Args[3], " / ", os.Args[2], " = ", result2)
    }
  }
}
