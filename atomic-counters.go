package main

import (
  "fmt"
  "runtime"
  "sync/atomic"
  "time"
)

// The primary mechanism for managing state in Go is 
// communication over channels. We saw this for example 
// with worker pools. There are a few other options for 
// managing state though. Here we’ll look at using the 
// sync/atomic package for atomic counters accessed by 
// multiple goroutines.
func main() {
  // We’ll use an unsigned integer to represent our 
  // (always-positive) counter.
  var ops uint64 = 0

  // To simulate concurrent updates, we’ll start 50 goroutines 
  // that each increment the counter about once a millisecond. 
  for i:= 0; i < 50; i++ {
    go func() {
      for {
        // To atomically increment the counter we use AddUint64, 
        // giving it the memory address of our ops counter with 
        // the & syntax.
        atomic.AddUint64(&ops, 1)
        // Allow other goroutines to proceed.
        runtime.Gosched()
      }
    }()
  }

  // In order to safely use the counter while it’s still being 
  // updated by other goroutines, we extract a copy of the current 
  // value into a var via LoadUint64. As above we need to give 
  // this function the memory address &ops from which to fetch the 
  // value.
  opsIntermediary := atomic.LoadUint64(&ops)
  fmt.Println("opsIntermediary:", opsIntermediary)

  // Wait a second to allow some ops to accumulate.
  time.Sleep(time.Second)

  opsFinal := atomic.LoadUint64(&ops)
  fmt.Println("opsFinal:", opsFinal)
}
