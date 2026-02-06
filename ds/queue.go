package ds

import "errors"

type Queue struct {
	list LinkedList
}

// add tail node
func (q *Queue) Push(value string) {
	q.list.Insert(value)
}

// remove the head
func (q *Queue) Pop() (string, error) {
	if q.list.Head == nil {
		return "", errors.New("queue is empty")
	}

	val := q.list.Head.data				//create new node "val" to store the value of list's head
	q.list.Head = q.list.Head.Next		//the next node of head becomes the new head of the list
	q.list.Size--						//list size - 1

	if q.list.Size == 0 {			//if list is empty, tail = nil
		q.list.Tail = nil
	}

	return val, nil
}
