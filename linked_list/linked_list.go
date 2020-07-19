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

func (list *LinkedList) Add(element int) error {
	if list.head == nil {
		list.head = &Node{
			next:  nil,
			prev:  nil,
			value: element,
		}
		list.tail = list.head
		return nil
	}

	new_node := &Node{
		prev:  list.tail,
		next:  nil,
		value: element,
	}

	list.tail.next = new_node
	list.tail = new_node

	return nil
}

func (list *LinkedList) Search(element int) error {
	return nil
}

func (list *LinkedList) RemoveLast() error {
	return nil
}
