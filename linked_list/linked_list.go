package linked_list

type Node struct {
	next  *Node
	prev  *Node
	value int
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (list *LinkedList) Add(element int) {
	if list.head == nil {
		list.head = &Node{
			next:  nil,
			prev:  nil,
			value: element,
		}
		list.tail = list.head
	} else {
		new_node := &Node{
			prev:  list.tail,
			next:  nil,
			value: element,
		}

		list.tail.next = new_node
		list.tail = new_node
	}
}

func (list *LinkedList) Search(element int) bool {
	if list.head == nil {
		return false
	}

	nodeHead := list.head
	nodeTail := list.tail
	for nodeHead.prev != nodeTail {
		if nodeHead.value == element || nodeTail.value == element {
			return true
		}
		nodeHead = nodeHead.next
		nodeTail = nodeTail.prev
	}
	return false
}

func (list *LinkedList) RemoveLast() error {
	if list.head == nil {
		return nil
	} else if list.head == list.tail {
		list.head = nil
		list.tail = nil
	} else {
		list.tail.prev.next = nil
		list.tail = list.tail.prev
	}
	return nil
}

func (list *LinkedList) RemoveFirst() error {
	if list.head == nil {
		return nil
	} else if list.head == list.tail {
		list.head = nil
		list.tail = nil
	} else {
		list.head.next.prev = nil
		list.head = list.head.next
	}
	return nil
}
