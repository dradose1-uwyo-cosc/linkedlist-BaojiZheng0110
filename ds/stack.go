package ds

type Stack struct {
	list LinkedList
}

// add new head node
func (s *Stack) Push(value string) {
	n := &Node{data: value}			//create a new node

	if s.list.Head == nil {			//if the head is nil, which mean the list is empty, n is the head and tail, list size = 1
		s.list.Head = n
		s.list.Tail = n
		s.list.Size = 1
		return
	}

	n.Next = s.list.Head		//the n pointer points to the head of the list
	s.list.Head = n				//the head of the list becomes to n
	s.list.Size++				//list size increase
}

// remove the head
func (s *Stack) Pop() (string, bool) {
	if s.list.Head == nil { 		//if the list is empty, can't remove anything
		return "", false
	}

	val := s.list.Head.data				//create a new node "val" to store the head value
	s.list.Head = s.list.Head.Next		//the next node of n becomes the head
	s.list.Size--

	if s.list.Size == 0 {			//if list is empty after removing, tail = nil
		s.list.Tail = nil
	}

	return val, true
}
