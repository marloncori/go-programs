pakcage main

import (
    "fmt"
    "time"
 )
 
 func main() {
     fmt.Prinln()
      x := [5]int{5, 4, 3, 2, 1}
      s := x[1:3]
      fmt.Println("Slice s = ", s)
      time.Sleep(2*time.Second)
      
      fmt.Println("Capacity of S slice: ", cap(s))
      time.Sleep(2*time.Second)
      
      fmt.Println("Length of S slice: ", len(s))
      time.Sleep(2*time.Second)
      
      fmt.Println("And the X slice equals: x={")
         for i := range x {
              fmt.Print(i, " ")
              time.Sleep(1*time.Second)
         }
         fmt.Println("}")
      
      //resize slice s, to add to it the other x slice elements
      new_s := s[:cap(s)]
      time.Sleep(2*time.Second)
      fmt.Println("\nThis is the new slice s = {")
       for j := range new_s {
           fmt.Print(j, " ")
        }
        fmt.Println("}")
 }
