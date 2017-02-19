package main
import (
  "fmt"
  "reflect"
)

func main() {
  s1 := 4.0
  fmt.Println(s1, "is a", reflect.TypeOf(s1))

  for key, value := range map[string]string{"Bob":"yes", "Bill":"yes", "Joe":"no"} {
        // for each pair in the map, print key and value
        fmt.Printf("key=%s, value=%s\n", key, value)
    }

}
