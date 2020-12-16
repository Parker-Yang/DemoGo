package test

import (
	"testing"
)

func TestChain(t *testing.T) {
	node := &Node{
		data: 1,
		next: nil,
	}
	l := &List{}
	l.init()
	l.append(node)
}

type Object interface{}
type Node struct {
	data Object
	next *Node
}
type List struct {
	Size uint
	Head *Node
	Tail *Node
}

func (l *List) init() {
	l.Size = 0
	l.Head = nil
	l.Tail = nil
}
func (l *List) append(node *Node) {
	if node == nil {
		return
	}
	if l.Size == 0 {
		l.Size = 1
		l.Head = node
		l.Tail = node
		l.Head.next = l.Tail
	} else {
		OldTail := l.Tail
		OldTail.next = node
		l.Tail = node
	}
}

