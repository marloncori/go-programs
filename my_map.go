package main

import (
  "fmt"
  "time"
)

func main() {
  myMap := make(map[string]int)
  
  myMap["key_1"] = 7
  myMap["key_2"] = 13
  myMap["key_3"] = 21
  
  fmt.Println("this is the myMap contents: ", myMap)
  time.Sleep(1*time.Second)
  
  fruits := map[string]int{
    "apple":5,
    "pear":6,
    "mango":7
    "banana":8
    }
  
  fmt.Println("This is the contents of the \"FRUITS\" map:")
  for k, v := range fruits {
    fmt.Printf("%s : %d%, k, v)
    time.Sleep(1*time.Second)
  }
  fruits = delete(fruits, "banana")
  fmt.Println("This is the new map: fruits = ", fruits)
  time.Sleep(1*time.Second)
  
  counter := 1
  fmt.Println("And these are the keys: ")
  for key := range fruits {
    fmt.Printf("key %d: %s", counter, key)
    counter += 1
    time.Sleep(1*time.Second)
   }
}
