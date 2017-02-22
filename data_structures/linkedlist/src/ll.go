package LinkedList
import (
  "fmt"
  "linkedlist/LLnode"
)

type LinkedList struct {
  first *LLnode
}

func NewLinkedList() LinkedList {
  return LinkedList{nil}
}

func (ll *LinkedList) Add(val int) {
  newNode := LLnode{val, nil}

  if ll.first == nil {
    ll.first = &newNode
  } else {
    curNode := ll.Last()
    curNode.next = &newNode
  }
}

func (ll *LinkedList) Find(val int) *LLnode {
  curNode := ll.first
  for curNode != nil {
    if curNode.val == val {
      return curNode
    } else {
      curNode = curNode.next
    }
  }
  return nil
}

func (ll *LinkedList) Delete(val int) {
  fmt.Printf("Attempting to delete node %d...\n", val)
  var prevNode *LLnode
  curNode := ll.first

  for curNode.val != val && curNode.next != nil {

    prevNode = curNode
    curNode = curNode.next
  }
  if curNode.val == val {
    fmt.Printf("Node %d deleted.\n", val)
    if prevNode == nil {
      ll.first = curNode.next
    } else {
      prevNode.next = curNode.next
    }
  } else {
    fmt.Printf("Node not found.\n")
  }
}

func (ll *LinkedList) First() *LLnode {
  return ll.first
}

func (ll *LinkedList) Last() *LLnode {
  curNode := ll.first
  for curNode.next != nil {
    curNode = curNode.next
  }
  return curNode
}

func (ll *LinkedList) PrintAll() {
  curNode := ll.first
  fmt.Printf("List: %d ", curNode.val)
  for curNode.next != nil {
    curNode = curNode.next
    fmt.Printf("-> %d ", curNode.val)
  }
  fmt.Printf("\n")
}
