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
	
}