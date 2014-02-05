package main

import (
  //"encoding/json"
  "fmt"
  "math"
  "reflect"
)

// Interfaces are named collections of method signatures.
// To learn more about gol's interfaces read this blog post:
// http://jordanorelli.tumblr.com/post/32665860244/how-to-use-interfaces-in-go

// Here’s a basic interface for geometric shapes.
type geometry interface {
  area() float64
  perim() float64
} 

// For our example we’ll implement this interface on 
// square and circle types.
type square struct {
  width, height float64
}

type circle struct {
  radius float64
}

// To implement an interface in Go, we just need 
// to implement all the methods in the interface. 
// Here we implement geometry on squares.
func (s square) area() float64 {
  return s.width * s.height
}
func (s square) perim() float64 {
  return 2*s.width + 2*s.height
}

// Here is the implementation for circles
func (c circle) area() float64 {
  return math.Pi*c.radius*c.radius
}
func (c circle) perim() float64 {
  return math.Pi*2*c.radius
}

// If a variable has an interface type, then 
// we can call methods that are in the named 
// interface. Here’s a generic measure function 
// taking advantage of this to work on any geometry.
func measure(g geometry) {
  fmt.Println(g)
  fmt.Println("area:", g.area())
  fmt.Println("perim:", g.perim())
}

func main() {
  s := square{width: 3, height: 4}
  c := circle{radius: 6}

  // The circle and square struct types both implement 
  // the geometry interface so we can use instances of 
  // these structs as arguments to measure.
  measure(s)
  measure(c)

  // since interfaces are types we can create slices of
  // types that implement the interface geometry:
  shapes := []geometry{s, c}
  for _, shape := range shapes {
    fmt.Println("type:", reflect.TypeOf(shape))
    fmt.Println(shape)
  }
}
