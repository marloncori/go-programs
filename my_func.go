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
  elem1 := 0
  elem2, elem3 := 0, 0
  
  if choice == "1" {
    elem1 _ = strconv.Atoi(os.Args[2])
    elem2, _ = strconv.Atoi(os.Args[3])
    
    result := plus(elem1, elem2)
    fmt.Printf("The result of adding %d to %d equals:\n", elem1, elem2)
    
    time.Sleep(1*time.Second)
    fmt.Println("\n", os.Args[2], " + ", os.Args[3], " = ", result)
    
  } else if choice == "2" {
    elem1 _ = strconv.Atoi(os.Args[2])
    elem2, _ = strconv.Atoi(os.Args[3])
    elem3 _ = strconv.Atoi(os.Args[4])
    
    res := addThree(elem1, elem2, elem3)
    fmt.Printf("The result of adding %d and %d to %d equals:\n", elem1, elem2, elem3)
    
    time.Sleep(1*time.Second)
    fmt.Println("\n", os.Args[2], " + ", os.Args[3], " + ", os.Args[4], " = ", res)
    
  } else {
    elem1 _ = strconv.Atoi(os.Args[2])
    elem2, _ = strconv.Atoi(os.Args[3])
    
    result1, result2 := mult_divide(elem1, elem2)
    fmt.Printf("The result of multiplying, dividing %d and %d equals:\n", elem1, elem2)
    
    time.Sleep(1*time.Second)
    fmt.Println("\n", os.Args[2], " * ", os.Args[3], " = ", result1)
    
    if elem1 > elem2 {
      time.Sleep(1*time.Second)
      fmt.Println("\n", os.Args[2], " / ", os.Args[3], " = ", result2)
    } else {
      time.Sleep(1*time.Second)
      fmt.Println("\n", os.Args[3], " / ", os.Args[2], " = ", result2)
    }
  }
}
