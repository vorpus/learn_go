package main
import (
  "fmt"
  "math/rand"
  "sync"
  "sync/atomic"
  "time"
)

func main() {

  var state = make(map[int]int)

  var mutex = &sync.Mutex{}

  var readOps uint64 = 0
  var writeOps uint64 = 0

  for r := 0; r < 100; r++ {
    go func() {
      total := 0
      for {
        key := rand.Intn(5)
        mutex.Lock()
        total += state[key]
        mutex.Unlock()
        atomic.AddUint64(&readOps, 1)

        time.Sleep(time.Millisecond)
      }
    }()
  }

  for w := 0; w < 10; w++ {
    go func() {
      for {
        key := rand.Intn(5)
        val := rand.Intn(100)
        mutex.Lock()
        state[key] =
      }
    }
  }
}
