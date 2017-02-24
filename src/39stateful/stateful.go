package main
import (
  "fmt"
  "math/rand"
  "sync/atomic"
  "time"
)

type readOp struct {
  key int
  resp chan int
}

type writeOp struct {
  key int
  val int
  resp chan bool
}

func main() {
  var readOps uint64 = 0
  var writeOps uint64 = 0

  reads := make(chan *readOp)
  writes := make(chan *writeOp)

  go func() {
    var state = make(map[int]int)
    for {
      select {
      case read := <-reads:
        read.resp <- state[read.key]
      case write := <-writes:
        state[write.key] = write.val
        write.resp <- true
      }
    }
  }()
}
