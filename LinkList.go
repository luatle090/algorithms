package algorithms

type LinkedList[T comparable] struct {
	first, last *Node[T]
	Length      int
}

type LinkedListDouble[T comparable] struct {
	header *NodeDouble[T]
	Length int
}

type NodeDouble[T comparable] struct {
	value     T
	next, pre *NodeDouble[T]
}

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

// func InitLinkedLink() LinkedList{

// }

func (link *LinkedList[T]) AddLastNode(v T) *Node[T] {
	node := &Node[T]{value: v, next: nil}

	if link.IsEmpty() {
		link.first = node
	} else {
		link.last.next = node
	}
	link.last = node
	link.Length++
	return node
}

// Viết lại, delete tail node nhưng link.last vẫn trỏ vào last cũ
func (link *LinkedList[T]) DeleteNode(v T) bool {

	if link.IsEmpty() {
		return false
	}

	if link.first.value == v {
		link.first = link.first.next
		link.Length--
		return true
	}

	found := false

	current := link.first

	preCurrent := current
	for i := 0; i < link.Length; i++ {
		if current.value == v {
			preCurrent.next = current.next
			link.Length--
			found = true
			break
		}
		preCurrent = current
		current = current.next
	}

	if link.Length == 0 {
		link.first, link.last = nil, nil
	}

	// if current.value == v {
	// 	link.first, link.last = nil, nil
	// 	found = true
	// } else {
	// 	for current.next.value != v {
	// 		if current.next.value == v {
	// 			current.next = current.next.next
	// 			found = true
	// 		}
	// 		current = current.next
	// 	}
	// }
	// if found {
	// 	link.Length--
	// }
	return found
}

func (link *LinkedList[T]) IsEmpty() bool {
	return link.Length == 0
}

func (link *LinkedList[T]) Contains(t T) bool {
	if link.IsEmpty() {
		return false
	}
	current := link.first
	for i := 0; i < link.Length; i++ {
		if current.value == t {
			return true
		}
		current = current.next
	}
	return false
}

// khởi tạo dummy header
func InitLinkedLinkDouble[T comparable](v T) *LinkedListDouble[T] {
	node := &NodeDouble[T]{next: nil, pre: nil}
	link := &LinkedListDouble[T]{
		header: node,
	}
	link.header.next, link.header.pre = link.header, link.header
	return link
}

// add theo last node
func (link *LinkedListDouble[T]) AddNode(v T) {
	node := &NodeDouble[T]{value: v}
	// last node cần trỏ vào dummy header
	node.next = link.header

	// pre của last node mới trỏ vào con trỏ last node cũ
	// tức pre node mới trỏ vào tail
	node.pre = link.header.pre

	// next của last node cũ trỏ vào node mới tương đương với 2 dòng code dưới
	// oldTail := link.header.pre
	// oldTail.next = node
	link.header.pre.next = node

	// pre của header trỏ vào node mới
	link.header.pre = node
}
