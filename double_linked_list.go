package main

import "fmt"

type Node struct {
	val  int
	next *Node
	prev *Node
}

type DoubleLinkedList struct {
	head *Node
}

func NewNode(val int) *Node {
	n := &Node{}
	n.val = val
	n.next = nil
	n.prev = nil
	return n
}

func (ll *DoubleLinkedList) AddAtBeg(val int) {
	n := NewNode(val)
	n.next = ll.head
	ll.head = n
}

func (ll *DoubleLinkedList) AddAtEnd(val int) {
	n := NewNode(val)
	if ll.head == nil {
		ll.head = n
		return
	}
	cur := ll.head
	prev := cur
	for prev = cur; cur != nil; cur = cur.next {
		prev = cur
	}

	prev.next = n
	n.prev = cur
}


func (ll *DoubleLinkedList) DelAtBeg() int{
	if ll.head == nil{
		return -1;
	}
	cur := ll.head
	ll.head = cur.next
	if ll.head != nil{
		ll.head.prev =  nil
	}
	return cur.val
}

func (ll *DoubleLinkedList) empty() bool{
	return ll.head == nil
}

func (ll * DoubleLinkedList) DeleteAtEnd() int{
	if ll.empty(){
		return -1
	}
	if ll.head.next == nil{
		return ll.DelAtBeg()
	}
	cur:= ll.head;
	for ; cur.next.next !=nil; cur = cur.next{
	}
	retval := cur.next.val
	cur.next = nil
	return retval

}

func (ll *DoubleLinkedList) Display() {
	for cur := ll.head; cur != nil; cur = cur.next {
		fmt.Print(cur.val, " ")
	}
	fmt.Print("\n")
}

func (ll *DoubleLinkedList) Count() int{
	var ctr int = 0
	for cur := ll.head; cur != nil; cur = cur.next {
		ctr += 1
	}
	return ctr
}

func (ll *DoubleLinkedList) Reverse() {
	var prev, next *Node
	cur := ll.head
	for cur != nil {
		next = cur.next
		cur.next = prev
		cur.prev = next
		prev =cur
		cur = next
	}
	ll.head = prev
}

func (ll *DoubleLinkedList) DisplayReverse() {
	ll.Reverse()
	ll.Display()
	ll.Reverse()
}

func main() {
	ll := DoubleLinkedList{}
	ll.AddAtBeg(10)
	ll.AddAtEnd(11)
	ll.AddAtBeg(14)
	ll.AddAtBeg(15)
	ll.AddAtBeg(34)
	ll.DisplayReverse()
	ll.Display()
	ll.Reverse()
	ll.Display()
	fmt.Println(ll.Count())
	ll.DeleteAtEnd()
	ll.Display()
	ll.DelAtBeg()
	ll.Display()
	fmt.Println(ll.Count())
}
