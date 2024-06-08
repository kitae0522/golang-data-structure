package list

import (
	"errors"
	"fmt"
)

/*

TODO
1. PushBack
2. PushFront
3. InsertBefore
4. InsertAfter
5. GetCount
6. DeleteNode
7. IncludeNode
8. GetRoot
9. GetTail

*/

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type List[T any] struct {
	root *Node[T]
	tail *Node[T]
	len  int
}

func NewNode[T any](value T, next *Node[T], prev *Node[T]) *Node[T] {
	return &Node[T]{value, next, prev}
}

func (n *Node[T]) GetValue() T {
	return n.value
}

func (l *List[T]) GetCount() int {
	return l.len
}

func (l *List[T]) GetRoot() *Node[T] {
	return l.root
}

func (l *List[T]) GetTail() *Node[T] {
	return l.tail
}

func (l *List[T]) PushBack(value T) {
	newNode := NewNode(value, nil, nil)
	if l.root == nil {
		l.root = newNode
		l.tail = newNode
		l.len = 1
		return
	}
	newNode.prev = l.tail
	l.tail.next = newNode
	l.tail = newNode
	l.len++
}

func (l *List[T]) PushFront(value T) {
	newNode := NewNode(value, nil, nil)
	if l.root == nil {
		l.root = newNode
		l.tail = newNode
		l.len = 1
		return
	}
	newNode.next = l.root
	l.root.prev = newNode
	l.root = newNode
	l.len++
}

func (l *List[T]) InsertBefore(value T, targetNode *Node[T]) error {
	newNode := NewNode(value, nil, nil)
	if _, ok := l.IncludeNode(targetNode); !ok {
		return errors.New("Node Not Found")
	}

	// targetNode is root
	if targetNode.prev == nil {
		l.PushFront(value)
		return nil
	}

	prevNodeCpy := targetNode.prev
	targetNode.prev = newNode

	newNode.next = targetNode
	newNode.prev = prevNodeCpy

	prevNodeCpy.next = newNode
	l.len++
	return nil
}

func (l *List[T]) InsertAfter(value T, targetNode *Node[T]) error {
	newNode := NewNode(value, nil, nil)
	if _, ok := l.IncludeNode(targetNode); !ok {
		return errors.New("Node Not Found")
	}

	// targetNode is tail
	if targetNode.next == nil {
		l.PushBack(value)
		return nil
	}

	nextNodeCpy := targetNode.next
	targetNode.next = newNode

	newNode.next = nextNodeCpy
	newNode.prev = targetNode

	nextNodeCpy.prev = newNode
	l.len++
	return nil
}

func (l *List[T]) IncludeNode(targetNode *Node[T]) (*Node[T], bool) {
	for e := l.root; e != nil; e = e.next {
		if e == targetNode {
			return e, true
		}
	}
	return nil, false
}

func (l *List[T]) GetNodeByIdx(idx int) *Node[T] {
	curNode := l.root
	for curIdx := 0; curIdx < idx; curIdx++ {
		curNode = curNode.next
	}
	return curNode
}

func (l *List[T]) DeleteNode(targetNode *Node[T]) error {
	if _, ok := l.IncludeNode(targetNode); !ok {
		return errors.New("Node Not Found")
	}

	// targetNode is root or tail
	if targetNode.prev == nil {
		l.root = targetNode.next
		l.len--
		return nil
	}
	if targetNode.next == nil {
		l.tail = targetNode.prev
		l.len--
		return nil
	}

	prevNodeCpy := targetNode.prev
	nextNodeCpy := targetNode.next

	prevNodeCpy.next = nextNodeCpy
	nextNodeCpy.prev = prevNodeCpy

	l.len--
	return nil
}

/*
func main() {
	newList := new(List[int])
	newList.PushBack(1)
	newList.PushBack(2)
	newList.PushFront(0)
	newList.InsertBefore(-1, newList.GetNodeByIdx(0))
	newList.InsertAfter(3, newList.GetNodeByIdx(3))

	for idx, e := 0, newList.root; e != nil; e = e.next {
		fmt.Printf("Idx[%d]: %v\n", idx, e.GetValue())
		idx++
	}
	fmt.Printf("Count: %d\n", newList.GetCount())

	fmt.Println("=================================")

	// after delete node
	newList.DeleteNode(newList.GetNodeByIdx(0))
	for idx, e := 0, newList.root; e != nil; e = e.next {
		fmt.Printf("Idx[%d]: %v\n", idx, e.GetValue())
		idx++
	}
	fmt.Printf("Count: %d\n", newList.GetCount())
}
*/
