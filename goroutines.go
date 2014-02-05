package main

import (
  "fmt"
  "time"
)

// A goroutine is a lightweight thread of execution.

func f(from string) {
  for i := 0; i < 3; i++ {
    fmt.Println(from, ":", i)
    time.Sleep(time.Second)
  }
}

func main() {
  // Suppose we have a function call f(s). Here’s how 
  // we’d call that in the usual way, running it synchronously.
  f("synchonous")

  // To invoke this function in a goroutine, use go f(s). 
  // This new goroutine will execute concurrently with the 
  // calling one.
  go f("asynchronous0")

  // You can also start a goroutine for an anonymous function call.
  go func(msg string) {
    fmt.Println(msg)
  }("anon asynchronous")

  go f("asynchronous1")

  // Our three goroutines are running asynchronously in 
  // separate goroutines now, so execution falls through 
  // to here. This Scanln code requires we press a key 
  // before the program exits.
  var input string
  fmt.Scanln(&input)
  fmt.Println("done")
}
