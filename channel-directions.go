package main

import "fmt"

// When using channels as function parameters, you can 
// specify if a channel is meant to only send or receive 
// values. This specificity increases the type-safety of 
// the program.

// XXX IMPORTANT: A channel that does NOT specify its 
// direction is known as bi-directional. A bi-directional 
// channel can be passed to a function that takes send-only 
// or receive-only channels, but the reverse is not true.
// (http://www.golang-book.com/10)

// This ping function only accepts a channel for sending 
// values. It would be a compile-time error to try to 
// receive on this channel.
func ping(pings chan<- string, msg string) {
  pings <- msg
}

// The pong function accepts one channel for receives (pings) 
// and a second for sends (pongs).
func pong(pings <-chan string, pongs chan<- string) {
  msg := <-pings
  pongs <- msg
}

func main() {
  pings := make(chan string, 1)
  pongs := make(chan string, 1)
  ping(pings, "passed message")
  pong(pings, pongs)
  fmt.Println(<-pongs)
}
