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

var (
	Red   = true
	Black = false
)

type NodeRedBlack[T Ordered] struct {
	color               bool
	key                 T
	left, right, parent *NodeRedBlack[T]
}

type RedBlackTree[T Ordered] struct {
	numOfNode               int
	root, blackNodeSentinel *NodeRedBlack[T]
}

func CreateRedBlackTree[T Ordered]() *RedBlackTree[T] {
	tree := new(RedBlackTree[T])
	tree.blackNodeSentinel = &NodeRedBlack[T]{
		color:  Black,
		left:   nil,
		right:  nil,
		parent: nil,
	}

	tree.root = tree.blackNodeSentinel
	return tree
}

// Khởi tạo node đỏ đen mới, color by default is red
func initNodeRedBack[T Ordered](val T, blackNodeSentinel, parrent *NodeRedBlack[T]) *NodeRedBlack[T] {
	return &NodeRedBlack[T]{
		key:    val,
		color:  Red,
		left:   blackNodeSentinel,
		right:  blackNodeSentinel,
		parent: parrent,
	}
}

// Get the number of nodes have been added into tree
func (rb *RedBlackTree[T]) GetSize() int {
	return rb.numOfNode
}

// Searching by key
func (rb *RedBlackTree[T]) SearchNode(val T) *NodeRedBlack[T] {
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
func (rb *RedBlackTree[T]) InorderWalk() []*NodeRedBlack[T] {
	store := make([]*NodeRedBlack[T], 0)
	if rb.root != rb.blackNodeSentinel {
		store = rb.inorderHelper(rb.root, store)
	}
	return store
}

func (rb *RedBlackTree[T]) inorderHelper(root *NodeRedBlack[T], store []*NodeRedBlack[T]) []*NodeRedBlack[T] {
	if root.left != rb.blackNodeSentinel {
		store = rb.inorderHelper(root.left, store)
	}
	store = append(store, root)
	if root.right != rb.blackNodeSentinel {
		store = rb.inorderHelper(root.right, store)
	}
	return store
}

func (rb *RedBlackTree[T]) GetRoot() *NodeRedBlack[T] {
	return rb.root
}

/// +++++++ func for Node
func (node NodeRedBlack[T]) GetColor() bool {
	return node.color
}

func (node NodeRedBlack[T]) GetKey() T {
	return node.key
}

func (node NodeRedBlack[T]) GetParent() *NodeRedBlack[T] {
	return node.parent
}

func (node NodeRedBlack[T]) GetLeft() *NodeRedBlack[T] {
	return node.left
}

func (node NodeRedBlack[T]) GetRight() *NodeRedBlack[T] {
	return node.right
}

/// -------

// Nhận vào value, thực hiện insert cho cây nhị phân. Lưu ý: func insertFix
func (rb *RedBlackTree[T]) Add(val T) {
	// nếu cây empty thì node mới là root
	// if rb.root == nil {
	// 	rb.numOfNode++
	// 	rb.initRootRedBlack(val)
	// 	return
	// }

	y := rb.blackNodeSentinel
	node := rb.root

	// tìm vị trí thích hợp để add node vào cây
	for node != rb.blackNodeSentinel {
		y = node // do duyệt tới node lá (node nil) vì vậy cần node y giữ lại vị trí cuối cùng
		if val <= node.key {
			node = node.left
		} else {
			node = node.right
		}
	}

	// tạo node mới với parent là y
	newNode := initNodeRedBack(val, rb.blackNodeSentinel, y)
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
func (rb *RedBlackTree[T]) insertFix(node *NodeRedBlack[T]) {
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
func (rb *RedBlackTree[T]) rotateLeft(node *NodeRedBlack[T]) {
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
}

// Phép xoay tương tự xoay trái
// Sau phép xoay, x thành parent của y và cây con phải của x thành cây trái của y
func (t *RedBlackTree[T]) rotateRight(node *NodeRedBlack[T]) {
	x := node.left
	node.left = x.right
	if x.right != t.blackNodeSentinel {
		x.right.parent = node
	}
	x.parent = node.parent
	if node.parent == t.blackNodeSentinel {
		t.root = x
	} else if node == node.parent.left {
		node.parent.left = x
	} else {
		node.parent.right = x
	}
	x.right = node
	node.parent = x
}

func (rb *RedBlackTree[T]) transplant(u, v *NodeRedBlack[T]) {
	if u.parent == rb.blackNodeSentinel {
		rb.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

func (rb *RedBlackTree[T]) Delete(val T) error {
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
func (rb *RedBlackTree[T]) TreeMinimun(node *NodeRedBlack[T]) *NodeRedBlack[T] {
	for node.left != rb.blackNodeSentinel {
		node = node.left
	}
	return node
}

// Hàm phụ trợ cho delete
func (rb *RedBlackTree[T]) deleteFixUp(nodeX *NodeRedBlack[T]) {
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
