package main

import "fmt"

func main() {
  
  nums := []int{2, 3, 4}
  sum := 0
  for _, num := range nums {
    sum += num
  }
  fmt.Println("sum:", sum)

  for i, num := range nums {
    if num == 3 {
      fmt.Println("index:", i)
    }
  }

  kvs := map[string]string{"a": "apple", "b": "banana"}
  for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k, v)
  }

  for i, runeValue := range "go" {
    //fmt.Println(i, runeValue)
    //fmt.Printf("%d: %#U\n", i, runeValue)
    fmt.Printf("%d: %q\n", i, runeValue)
  }

  // see this blog post on string internals:
  // http://blog.golang.org/strings
  const nihongo = "日本語"
  for index, runeValue := range nihongo {
    fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
  }
}
