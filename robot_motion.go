package main

import (
  "bufio"
  "fmt"
  "strings"
  "os"
)

func move_robot(cmd string) string {
  cmd, _ := reader.ReadString('\n')
    // convert CRLF to LF
  cmd = strings.Replace(cmd, "\n", "", -1)
  msg := " "
   seq := make([]int,4)
  
    if strings.Compare("A", cmd) == 0 || strings.Compare("a", cmd) == 0 {
       msg = "\tRobot is turning left!" 
       for step := range seq {
         fmt.Println(msg)
          time.Sleep(1*time.Second)
       }
     }
    if strings.Compare("W", cmd) == 0 || strings.Compare("w", cmd) == 0 {
       msg = "\tRobot is moving forward!"
       for step := range seq {
           fmt.Println(msg)
           time.Sleep(1*time.Second)
        }
     }
    if strings.Compare("D", cmd) == 0 || strings.Compare("d", cmd) == 0 {
       msg = "\tRobot is turning right!"
       for step := range seq {
           fmt.Println(msg)
           time.Sleep(1*time.Second)
        }
     }
    if strings.Compare("X", cmd) == 0 || strings.Compare("x", cmd) == 0 {
      msg = "\tRobot is moving backward!"
      for step := range seq {
           fmt.Println(msg)
           time.Sleep(1*time.Second)
        }
     }
    if strings.Compare("S", cmd) == 0 || strings.Compare("s", cmd) == 0 {
       msg = "\tRobot has stopped its motors. Bye-bye!"
       fmt.Println(msg)
     }
    return msg
 }

func main() {

 reader := bufio.NewReader(os.Stdin)
 char, _, err := reader.ReadRune()

 if err != nil {
   fmt.Println(err)
 }
 // print out the unicode value i.e. A -> 65, a -> 97
 fmt.Println(char)
  command := " "
switch char {
 case 'A', 'a':
   fmt.Println("A key has been pressed!")
   command = "A"
   move_robot(command)
   break
 case 'W', 'w':
   fmt.Println("W key has been pressed!")
   command = "W"
   move_robot(command)
   break
 case 'D', 'd':
   fmt.Println("D key has been pressed!")
   command = "D"
   move_robot(command)
   break
 case 'X', 'x':
   fmt.Println("X key has been pressed!")
   command = "X"
   move_robot(command)
   break
 case 'S', 's':
   fmt.Println("S key has been pressed!")
   command = "S"
   move_robot(command)
   break
 default:
   fmt.Println("\tPlease press a valid key for the robot to move!")
 }
}
