package main
import (
  "fmt"
  "reflect"
)

func main() {
  // s1 := []int{1,2,3,4}
  // s1 := make([]int, 4)
  var s1 [][]float64
  fmt.Println(s1)
  fmt.Println(reflect.TypeOf(s1))
}
