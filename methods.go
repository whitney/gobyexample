package main

import "fmt"

// Go supports methods defined on struct types.

type rect struct {
  width, height int
}

// Methods can be defined for either pointer or 
// value receiver types: 

// This area method has a receiver type of *rect (pointer).
func (r *rect) area() int {
  return r.width * r.height
}

// This perim method has a receiver type of rect (value).
func (r rect) perim() int {
  return 2*r.width + 2*r.height
}

// You may want to use a pointer receiver type to avoid 
// copying on method calls or to allow the method 
// to mutate the receiving struct.
func (r *rect) scale(scalar int) {
  r.width = r.width * scalar
  r.height = r.height * scalar
}

func main() {
  r := rect{width: 10, height: 5}
  
  // Go automatically handles conversion between 
  // values and pointers for method calls.
  fmt.Println("area:", r.area())
  fmt.Println("perim:", r.perim())

  rp := &r
  fmt.Println("ptr area:", rp.area())
  fmt.Println("ptr perim:", rp.perim())

  fmt.Println("rp pre-scale:", *rp)
  rp.scale(2)
  fmt.Println("rp post-scale:", *rp)
}
