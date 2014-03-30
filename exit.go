package main

import (
  "fmt"
  "os"
)
// Use os.Exit to immediately exit with a given status.
func main() {
  // defers will not be run when using os.Exit, so this 
  // fmt.Println will never be called.
  defer fmt.Println("I am never called!")

  // Exit with status 3
  os.Exit(3)
}

// Note that unlike e.g. C, Go does not use an integer return 
// value from main to indicate exit status. If you’d like to 
// exit with a non-zero status you should use os.Exit.

// If you run exit.go using go run, the exit will be picked up 
// by go and printed.

// By building and executing a binary you can see the status in 
// the terminal.
// $ go build exit.go
// ./exit