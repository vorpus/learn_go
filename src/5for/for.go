package main

import (
  "fmt"
)

func main() {
  i := 1

  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }

  fmt.Println()

  for j := 7; j <= 9; j++ {
    fmt.Println(j)
  }

  for k := 1; k > 0; k++ {
    fmt.Println(k)
    if (k == 5) {
      break
    }
  }

  for {
    fmt.Println("loop")
    break
  }

  for n := 0; n <= 5; n++ {
    if n%2 == 0 {
      continue
    }
    fmt.Println(n)
  }
}
