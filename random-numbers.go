package main

import (
  "fmt"
  "math/rand"
)

// Go’s math/rand package provides pseudorandom number generation.
func main() {
  // For example, rand.Intn returns a random int n in [0, 100)
  fmt.Print(rand.Intn(100), ",")
  fmt.Print(rand.Intn(100))
  fmt.Println()

  // rand.Float64 returns a float64 f in [0.0, 1.0)
  fmt.Println(rand.Float64())

  // This can be used to generate random floats in other ranges, 
  // for example 5.0 <= f' < 10.0
  fmt.Print((rand.Float64()*5)+5, ",")
  fmt.Print((rand.Float64()*5)+5)
  fmt.Println()

  // To make the pseudorandom generator deterministic, 
  // give it a well-known seed.
  s1 := rand.NewSource(42)
  r1 := rand.New(s1)

  // Call the resulting rand.Source just like the functions 
  // on the rand package.
  fmt.Print(r1.Intn(100), ",")
  fmt.Print(r1.Intn(100))
  fmt.Println()

  // If you seed a source with the same number, it produces 
  // the same sequence of random numbers.
  s2 := rand.NewSource(42)
  r2 := rand.New(s2)
  fmt.Print(r2.Intn(100), ",")
  fmt.Print(r2.Intn(100))
  fmt.Println()
}
