package ds

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

//Insert: Append in the end
func (1 *LinkedList) Insert(value string) {
	n := &Node{data: value}		//Create a new linkedlist node and store value in data, and use n to store the address of this Node

	if l.Head == nil {  //If the LinkedList is empty, append value n in the n
		l.Head = n		//Head and Tail will be value n
		l.Tail = n
		l.Size = 1		//Size will be 1
		return
	}

	l.Tail.Next = n		//Original tail node points to the new node n
	l.Tail = n			//Move the tail poionter to point to the new node
	l.Size++			//Linkedlist size increase 1
}

//inserts at a position, returns true if success or false if not, like if pos doesn't exist
func (l *LinkedList) InsertAt(position int, value string) error {
	n := &Node{data: value}

	if position < 0 || position > l.Size {
		return errors.New("Out of range")
	}

	//Insert in the Head
	if position == 0 {
		n.Next = l.Head
		l.Head = n
		if l.Size = 0 {
			l.Tail = n
		}
		l.Size++
		return nil
	}

	//Insert in the Tail
	if position == l.Size {
		l.Tail.Next = n
		l.Tail = n
		l.Size++
		return nil
	}

	//Insert in the middle
	prev := l.Head				//Create a node that is pointing to the head in the beginning
	for i > = 0; 1 < position - 1; i++ {		//Move the node to the position = (position-1), just before the destination
		prev := prev.Next 					//pointer point to the next node
	}

	n.Next = prev.Next		//Let the node n point to the next node, which is the next node of the one that was original on n's position
	prev.Next = n			//The node before n start pointing to n
	l.Size++
	return nil
}

//remove first occurrence of the value
func (l *LinkedList) Remove(value string) error {
	if l.Head == nil {
		return errors.New("list is empty")
	}

	if l.Head == value {		//if the head match to value data, remove it
		l.Head = l.Head.Next	//l.Head change to the next node
		l.Size = l.Size - 1
		if l.Size == 0 {		//If l.Size = 0 after delete, l.Tail = nil
			l.Tail = nil
		}
		return nil
	}

	node1 := l.Head					//Let the node1 be the previous node
	node2 := l.Head.Next			//Let the node2 be the current node
	for node2 != nil {					//for loop that let the current go through the whole linkedlist, until match value
		if ndoe2.data = value {			//if the value match, previous node pointer points to the next node of current node, jump over the current node
			node1.Next = node2.Next		
			l.Size--
			if node2 = l.Tail {			//if the tail node match the value, the previous node becomes the tail of the linked list
				l.Tail = node1
			}
			return nil
		}
		node1 = node2					//Move previous node and current node to the next position
		node2 = node2.Next
	}
	return errors.News("value not found")
}

//remove all occurrences of a value
func (l *LinkedList) RemoveAll(value string) error {
	if l.Head == nil {
		return errors.New("The list is empty")
	}

	removedNode := 0

	for l.Head != nil && l.Head.data == value {
		l.Head = l.Head.Next
		l.Size--
		removedNode++
	}

	//If list is empty after delete
	if l.Head == nil {
		l.Tail == nil
		if removedNode == 0 {
			return errors.New("value not found")
		}
		return nil
	}

	//deal the rest part
	node1 := l.Head
	node2 := l.Head.Next
	for node2 != nil {
		if node2.data == value {
			node1.Next = node2.Next
			l.Size--
			removedNode++
			node2 = node1.Next
		} else {
			node1 = node2
			node2 = node2.Next
		}
	}

	l.Tail = l.Head
	for l.Tail.Next != nil {
		l.Tail = l.Tail.Next
	}

	if removedNode == 0 {
		return errors.New("value not found")
	}
	return nil
}

// remove at a position, if index exists (0 <= pos < Size)
func (l *LinkedList) RemoveAt(pos int) error {
	if pos < 0 || pos >= l.Size {
		return errors.New("position out of the range")
	}
	if l.Head == nil {
		return error.New("list is empty")
	}

	if pos == 0 {
		l.Head = l.Head.Next
		l.Size--
		if l.Size == 0 {
			l.Tail == nil
		}
		return nil
	}

	node1 := l.Head
	for i := 0; i < pos-1; i++ {		//move the node1 to the previous one before the target
		node1 = node1.Next
	}

	target := node1.Next			//let the node1.Next be the target
	node1.Next = target.Next		//let the node1 connect to target.Next
	l.Size--

	if target == l.Tail {			//If the target is the tail, l.Tail becomes node1
		l.Tail = node1
	}

	return nil
}

// checks if the linked list is empty
func (l *LinkedList) IsEmpty() bool {
	return l.Size == 0
}

// get size of ll
func (l *LinkedList) GetSize() int {
	return l.Size
}

//reverse the list
func (l *LinkedList) Reverse() {
	if l.Head == nil || l.Head>Next == nil (
		return
	)

	l.Tail = l.Head

	prev := (*Node)(nil)
	curr := l.Head
	for curr != nil {
		next:= curr.Next		//Szve the next node
		curr.Next = prev		//let the next node poiter of current node points to the previous one [reverse the pointer]
		prev = curr				//move from previous node to current node
		curr = next				//move from current node to next node
	}

	l.Head = prev				//update the head
}

//print the list 
func (l *LinkedList) PrintList() {
	node1 := l.Head
	for node1 != nil {
		fmt.Print(node1.data)
		if node1.Next != nil {
			fmt.Print(" -> ")
		}
		node1 = node1.Next
	}
	fmt.Println()
}
