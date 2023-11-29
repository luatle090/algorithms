package algorithms

import "fmt"

//// author: Luat Le
//// implementation follow Introduction to Algorithms - 3rd
//// Cây đỏ-đen có 5 thuộc tính:
//// 1. Mọi node là đỏ hoặc đen
//// 2. Root luôn là đen
//// 3. Mọi node lá (nil) là node đen
//// 4. Nếu node là đỏ, thì cả 2 node con phải là đen
//// 5. Với mọi node, đường đi đơn giản từ node tới lá chứa cùng số lượng node đen
//// Node lá sẽ được làm lính canh luôn là node đen, giá trị nil, mọi node gần node lá sẽ trỏ về node này
////
//// Các thuộc tính cài đặt của node gồm: color, key, left, right, parent

type NodeRedBlackOrderStatistic[T Ordered] struct {
	color               bool
	key                 T
	size                int // order-statistic tree
	left, right, parent *NodeRedBlackOrderStatistic[T]
}

type RedBlackOrderStatisticTree[T Ordered] struct {
	numOfNode               int
	root, blackNodeSentinel *NodeRedBlackOrderStatistic[T]
}

func CreateOrderStatistic[T Ordered]() *RedBlackOrderStatisticTree[T] {
	tree := new(RedBlackOrderStatisticTree[T])
	tree.blackNodeSentinel = &NodeRedBlackOrderStatistic[T]{
		color:  Black,
		left:   nil,
		right:  nil,
		parent: nil,
		size:   0,
	}

	tree.root = tree.blackNodeSentinel
	return tree
}

// Khởi tạo node đỏ đen mới, color by default is red
func initNodeOrderStatistic[T Ordered](val T, blackNodeSentinel, parrent *NodeRedBlackOrderStatistic[T]) *NodeRedBlackOrderStatistic[T] {
	return &NodeRedBlackOrderStatistic[T]{
		key:    val,
		color:  Red,
		size:   1,
		left:   blackNodeSentinel,
		right:  blackNodeSentinel,
		parent: parrent,
	}
}

// Get the number of nodes have been added into tree
func (rb *RedBlackOrderStatisticTree[T]) GetSize() int {
	return rb.numOfNode
}

// Search by key. Return a node sentinel if not found value
func (rb *RedBlackOrderStatisticTree[T]) SearchNode(val T) *NodeRedBlackOrderStatistic[T] {
	node := rb.root
	for node != rb.blackNodeSentinel && node.key != val {
		if val <= node.key {
			node = node.left
		} else {
			node = node.right
		}
	}
	return node
}

// Walking entire nodes in the tree. Return an array nodes in the tree
func (rb *RedBlackOrderStatisticTree[T]) InorderWalk() []*NodeRedBlackOrderStatistic[T] {
	store := make([]*NodeRedBlackOrderStatistic[T], 0)
	if rb.root != rb.blackNodeSentinel {
		store = rb.inorderHelper(rb.root, store)
	}
	return store
}

func (rb *RedBlackOrderStatisticTree[T]) inorderHelper(root *NodeRedBlackOrderStatistic[T], store []*NodeRedBlackOrderStatistic[T]) []*NodeRedBlackOrderStatistic[T] {
	if root.left != rb.blackNodeSentinel {
		store = rb.inorderHelper(root.left, store)
	}
	store = append(store, root)
	if root.right != rb.blackNodeSentinel {
		store = rb.inorderHelper(root.right, store)
	}
	return store
}

func (rb *RedBlackOrderStatisticTree[T]) GetRoot() *NodeRedBlackOrderStatistic[T] {
	return rb.root
}

/// +++++++ func for Node
func (node NodeRedBlackOrderStatistic[T]) GetColor() bool {
	return node.color
}

func (node NodeRedBlackOrderStatistic[T]) GetKey() T {
	return node.key
}

func (node NodeRedBlackOrderStatistic[T]) GetParent() *NodeRedBlackOrderStatistic[T] {
	return node.parent
}

func (node NodeRedBlackOrderStatistic[T]) GetLeft() *NodeRedBlackOrderStatistic[T] {
	return node.left
}

func (node NodeRedBlackOrderStatistic[T]) GetRight() *NodeRedBlackOrderStatistic[T] {
	return node.right
}

/// -------

// Nhận vào value, thực hiện insert cho cây nhị phân. Lưu ý: func insertFix
func (rb *RedBlackOrderStatisticTree[T]) Add(val T) {
	y := rb.blackNodeSentinel
	node := rb.root

	// tìm vị trí thích hợp để add node vào cây
	for node != rb.blackNodeSentinel {
		y = node // do duyệt tới node lá (node nil) vì vậy cần node y giữ lại vị trí cuối cùng
		y.size++
		if val <= node.key {
			node = node.left
		} else {
			node = node.right
		}
	}

	// tạo node mới với parent là y
	newNode := initNodeOrderStatistic(val, rb.blackNodeSentinel, y)
	if y == rb.blackNodeSentinel {
		rb.root = newNode
	} else if val <= y.key {
		y.left = newNode
	} else {
		y.right = newNode
	}
	rb.insertFix(newNode)
	rb.numOfNode++
}

// Thực hiện đổi màu, cân bằng cây
func (rb *RedBlackOrderStatisticTree[T]) insertFix(node *NodeRedBlackOrderStatistic[T]) {
	for node.parent.color == Red {
		if node.parent == node.parent.parent.left {
			y := node.parent.parent.right
			if y.color == Red {
				node.parent.color = Black
				y.color = Black
				node.parent.parent.color = Red
				node = node.parent.parent
			} else {
				if node == node.parent.right {
					node = node.parent
					rb.rotateLeft(node)
				}
				node.parent.color = Black
				node.parent.parent.color = Red
				rb.rotateRight(node.parent.parent)
			}
		} else {
			// else này giống như if bên trên nhưng là cho right child
			y := node.parent.parent.left
			if y.color == Red {
				node.parent.color = Black
				y.color = Black
				node.parent.parent.color = Red
				node = node.parent.parent
			} else {
				if node == node.parent.left {
					node = node.parent
					rb.rotateRight(node)
				}
				node.parent.color = Black
				node.parent.parent.color = Red
				rb.rotateLeft(node.parent.parent)
			}
		}
	}
	// the root always is black color
	rb.root.color = Black
}

// Phép xoay giúp duy trì thuộc tính của cây đỏ-đen. Khi xoay tại node xoay x,
// giả định cây con phải của y không được là node nil, x là node bất kỳ trong cây
// có x.right khác nil.
// Sau phép xoay: y sẽ là cha của x, cây con trái của y sẽ là cây con phải của x
//
// Thời gian phép xoay là O(1)
func (rb *RedBlackOrderStatisticTree[T]) rotateLeft(node *NodeRedBlackOrderStatistic[T]) {
	y := node.right
	// Đặt cây con trái của y thành cây con phải của node x
	node.right = y.left
	// set lại parent cho cây con trái của y
	// sau phép gán parent của cây con trái y sẽ là node x
	if y.left != rb.blackNodeSentinel {
		y.left.parent = node
	}
	// đổi parent. link parent của x vào y
	y.parent = node.parent

	// đổi root hoặc subtree
	if node.parent == rb.blackNodeSentinel {
		rb.root = y
	} else if node == node.parent.left {
		node.parent.left = y
	} else {
		node.parent.right = y
	}

	// xoay y thành parent của x
	y.left = node
	node.parent = y

	rb.asb(node, y)
}

// Phép xoay tương tự xoay trái
// Sau phép xoay, x thành parent của y và cây con phải của x thành cây trái của y
func (rb *RedBlackOrderStatisticTree[T]) rotateRight(node *NodeRedBlackOrderStatistic[T]) {
	x := node.left
	node.left = x.right
	if x.right != rb.blackNodeSentinel {
		x.right.parent = node
	}
	x.parent = node.parent
	if node.parent == rb.blackNodeSentinel {
		rb.root = x
	} else if node == node.parent.left {
		node.parent.left = x
	} else {
		node.parent.right = x
	}
	x.right = node
	node.parent = x

	rb.asb(node, x)
}

// Node from assigment to node to, then node from will sum size left and right child
func (t *RedBlackOrderStatisticTree[T]) asb(nodeFrom, nodeTo *NodeRedBlackOrderStatistic[T]) {
	nodeTo.size = nodeFrom.size
	nodeFrom.size = nodeFrom.left.size + nodeFrom.right.size + 1
}

func (rb *RedBlackOrderStatisticTree[T]) transplant(u, v *NodeRedBlackOrderStatistic[T]) {
	if u.parent == rb.blackNodeSentinel {
		rb.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

func (rb *RedBlackOrderStatisticTree[T]) Delete(val T) error {
	z := rb.SearchNode(val)
	if z == rb.blackNodeSentinel {
		return fmt.Errorf("a value not have in the tree")
	}
	y := z
	originalColorOfNodeY := y.color
	x := z
	if z.left == rb.blackNodeSentinel {
		x = z.right
		rb.transplant(z, z.right)
	} else if z.right == rb.blackNodeSentinel {
		x = z.left
		rb.transplant(z, z.left)
	} else {
		y = rb.TreeMinimun(z.right)
		originalColorOfNodeY = y.color
		x = y.right
		if y.parent == z {
			x.parent = y
		} else {
			rb.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		rb.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}

	if originalColorOfNodeY == Black {
		rb.deleteFixUp(x)
	}
	return nil
}

// Get node has key is a minimun
func (rb *RedBlackOrderStatisticTree[T]) TreeMinimun(node *NodeRedBlackOrderStatistic[T]) *NodeRedBlackOrderStatistic[T] {
	for node.left != rb.blackNodeSentinel {
		node = node.left
	}
	return node
}

// Hàm phụ trợ cho delete
func (rb *RedBlackOrderStatisticTree[T]) deleteFixUp(nodeX *NodeRedBlackOrderStatistic[T]) {
	for nodeX != rb.root && nodeX.color == Black {
		if nodeX == nodeX.parent.left {
			w := nodeX.parent.right
			if w.color == Red {
				w.color = Black
				nodeX.parent.color = Red
				rb.rotateLeft(nodeX.parent)
				w = nodeX.parent.right
			}
			if w.left.color == Black && w.right.color == Black {
				w.color = Red
				nodeX = nodeX.parent
			} else {
				if w.right.color == Black {
					w.left.color = Black
					w.color = Red
					rb.rotateRight(w)
					w = nodeX.parent.right
				}
				w.color = nodeX.parent.color
				nodeX.parent.color = Black
				w.right.color = Black
				rb.rotateLeft(nodeX.parent)
				nodeX = rb.root
			}
		} else {
			w := nodeX.parent.right
			if w.color == Red {
				w.color = Black
				nodeX.parent.color = Red
				rb.rotateRight(nodeX.parent)
				w = nodeX.parent.right
			}
			if w.right.color == Black && w.left.color == Black {
				w.color = Red
				nodeX = nodeX.parent
			} else {
				if w.left.color == Black {
					w.left.color = Black
					w.color = Red
					rb.rotateLeft(w)
					w = nodeX.parent.left
				}
				w.color = nodeX.parent.color
				nodeX.parent.color = Black
				w.left.color = Black
				rb.rotateRight(nodeX.parent)
				nodeX = rb.root
			}
		}
	}
	nodeX.color = Black
}

// +++++++++ tìm phần tử nhỏ thứ i
// ví dụ: tìm phần tử nhỏ thứ 3 trong cây

func OsSelect[T Ordered](nodeX *NodeRedBlackOrderStatistic[T], i int) *NodeRedBlackOrderStatistic[T] {
	r := nodeX.left.size + 1
	if i == r {
		return nodeX
	} else if i < r {
		return OsSelect[T](nodeX.left, i)
	}
	return OsSelect[T](nodeX.right, i-r)
}

// Find the rank of node in the tree
func (rb *RedBlackOrderStatisticTree[T]) OsRank(nodeX *NodeRedBlackOrderStatistic[T]) int {
	r := nodeX.left.size + 1
	y := nodeX
	for y != rb.blackNodeSentinel {
		if y == y.parent.right {
			r = r + y.parent.left.size + 1
		}
		y = y.parent
	}
	return r
}

/// ++++++++++
