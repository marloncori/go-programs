package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {

  reader := bufio.NewReader(os.Stdin)
  fmt.Println("\tSimple Chatbot - CONVERSATION")
  fmt.Println("---------------------")
  
  fmt.Println("\tChatbot has entered the room.")
  for {
    fmt.Print("-> ")
    text, _ := reader.ReadString('\n')
    // convert CRLF to LF
    text = strings.Replace(text, "\n", "", -1)

    if strings.Compare("Hi", text) == 0 || strings.Compare("Hello", text) == 0 || strings.Compare("Hi, Chatbot!", text) == 0 {
      fmt.Println("Hello, how are you?")
    }
    else {
      fmt.Println("I could not grasp it, sorry...")
    }
   
    fmt.Print("-> ")
    text2, _ := reader.ReadString('\n')
    // convert CRLF to LF
    text2 = strings.Replace(text2, "\n", "", -1)
    if strings.Compare("I'm fine.", text2) == 0 || strings.Compare("I'm fine, thanks.", text) == 0 || strings.Compare("Nice, thank you", text) == 0 {
      fmt.Println("It is good to here that. And what is your name?")
    }
    
    fmt.Print("-> ")
    text3, _ := reader.ReadString('\n')
    // convert CRLF to LF
    text3 = strings.Replace(text3, "\n", "", -1)
  
    switch text3 {
      case "Marlon":
       fmt.Println("\tIt was nice to talk to you, master Marlon!")
       break
      case "Aniko":
       fmt.Println("\tYour wife's name is so beautiful.")
       break
      case "Beniamin":
       fmt.Println("\tMay God bless your son ", text3, "A LOT!")
       break
      default:
       fmt.Println("\tWhat a strange name...")
       break
    }
  }
}
