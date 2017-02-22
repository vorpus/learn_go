package LLnode

type LLnode struct {
  val  int
  next *LLnode
}

func (n *LLnode) Value() int {
  return n.val
}
