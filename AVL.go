package algorithms

import (
	"fmt"
	"math"
)

type Ordered interface {
	~int | ~float32 | ~string
}

type BinaryNode[T Ordered] struct {
	height      int
	val         T
	Left, Right *BinaryNode[T]
}

type AVLTree[T Ordered] struct {
	Root *BinaryNode[T]
}

var HeightNodeParrent = 1

func initBinaryNode[T Ordered](val T) *BinaryNode[T] {
	return &BinaryNode[T]{val: val, Left: nil, Right: nil, height: 0}
}

func InitAVLTree[T Ordered]() *AVLTree[T] {
	return new(AVLTree[T])
}

func (bt *AVLTree[T]) Iterator() {
	if bt.Root == nil {
		fmt.Printf("no have values")
		return
	}
	bt.Root.inorder()
}

// func (node *BinaryNode[T]) inorder(list []T) []T {
// 	if node.Left != nil {
// 		list = node.Left.inorder(list)
// 	}
// 	list = append(list, node.val)
// 	if node.Right != nil {
// 		list = node.Right.inorder(list)
// 	}
// 	return list
// }

func (node *BinaryNode[T]) inorder() {
	if node.Left != nil {
		node.Left.inorder()
	}
	fmt.Printf("%v ", node.val)
	if node.Right != nil {
		node.Right.inorder()
	}
}

func (bt *AVLTree[T]) AddNode(val T) {
	if bt.Root == nil {
		bt.Root = initBinaryNode(val)
	} else {
		bt.Root = bt.Root.add(val)
	}
}

func (bt *AVLTree[T]) RemoveNode(val T) error {
	if bt.Root == nil {
		return fmt.Errorf("error")
	}

	bt.Root = bt.Root.remove(val)
	return nil
}

// Thêm giá trị mới vào BST bắt nguồn từ nút đó và trả về một BinaryNode đối tượng,
// đối tượng này sẽ trở thành gốc mới của BST đó.
// Điều này xảy ra do thao tác xoay sẽ di chuyển một nút mới làm gốc của BST
func (node *BinaryNode[T]) add(val T) *BinaryNode[T] {
	// gán nút newRoot là node nếu ko có phép xoay cân bằng thì newRoot là node cũ
	newRoot := node

	if val <= node.val {
		// add vào cây con trái
		node.Left = node.addToSubTree(node.Left, val)
		// quá trình xoay sẽ thay đổi nút gốc
		if node.heightDiff() == 2 {
			// lệch left-left
			if val <= node.Left.val {
				newRoot = node.rotateRight()
			} else {
				// lệch left-right
				newRoot = node.rotateLeftRight()
			}
		}
	} else {
		// add vào cây con phải
		node.Right = node.addToSubTree(node.Right, val)
		// quá trình xoay sẽ thay đổi nút gốc
		if node.heightDiff() == -2 {
			// lệch right-right
			if val > node.Right.val {
				newRoot = node.rotateRight()
			} else {
				// lệch right-left
				newRoot = node.rotateRightLeft()
			}
		}
	}

	// tính lại độ cao của new root
	newRoot.computeHeight()
	return newRoot
}

// Hàm helper cho quá trình gọi đệ quy để add value vào BST
// Điều kiện dừng khi node parent là nil.
func (node *BinaryNode[T]) addToSubTree(parent *BinaryNode[T], val T) *BinaryNode[T] {
	if parent == nil {
		return initBinaryNode(val)
	}
	parent = parent.add(val)
	return parent
}

// tính độ cao chênh lệch của cây con trái và cây con phải
func (node *BinaryNode[T]) heightDiff() int {
	heightLeft := 0
	heightRight := 0
	if node.Left != nil {
		heightLeft = node.Left.height + HeightNodeParrent
	}
	if node.Right != nil {
		heightRight = node.Right.height + HeightNodeParrent
	}
	return heightLeft - heightRight
}

// tính độ cao của node
// nếu node leaf thì height sẽ là 0
// nếu node cha thì height sẽ là 1
// chiều cao của node không tồn tại là –1
func (node *BinaryNode[T]) computeHeight() {
	height := -1
	if node.Left != nil {
		height = int(math.Max(float64(height), float64(node.Left.height)))
	}
	if node.Right != nil {
		height = int(math.Max(float64(height), float64(node.Right.height)))
	}
	node.height = height + 1
}

func (node *BinaryNode[T]) rotateRight() *BinaryNode[T] {
	// root := node.Left
	// node.Left.Right, node.Left = node, node.Left.Right
	// node.computeHeight()

	root := node.Left
	grandson := root.Right
	root.Right = node
	node.Left = grandson
	node.computeHeight()
	return root
}

func (node *BinaryNode[T]) rotateLeft() *BinaryNode[T] {
	root := node.Right
	grandson := root.Left
	root.Left = node
	node.Right = grandson
	node.computeHeight()
	return root
}

func (node *BinaryNode[T]) rotateLeftRight() *BinaryNode[T] {
	child := node.Left.rotateLeft()
	node.Left = child
	newRoot := node.rotateRight()
	// node.computeHeight()
	// child.computeHeight()
	return newRoot
}

func (node *BinaryNode[T]) rotateRightLeft() *BinaryNode[T] {
	child := node.Right.rotateRight()
	node.Right = child
	newRoot := node.rotateLeft()
	// node.computeHeight()
	// child.computeHeight()
	return newRoot
}

func (node *BinaryNode[T]) removeFromParent(parent *BinaryNode[T], val T) *BinaryNode[T] {
	if parent != nil {
		return parent.remove(val)
	}
	return nil
}

func (node *BinaryNode[T]) remove(val T) *BinaryNode[T] {
	newRoot := node
	if val == node.val {
		if node.Left == nil {
			return node.Right
		}
		child := node.Left
		for child.Right != nil {
			child = child.Right
		}

		childKey := child.val
		node.Left = node.removeFromParent(node.Left, childKey)
		node.val = childKey

		if node.heightDiff() == -2 {
			if node.Right.heightDiff() <= 0 {
				newRoot = node.rotateLeft()
			} else {
				newRoot = node.rotateRightLeft()
			}
		}
	} else if val < node.val {
		node.Left = node.removeFromParent(node.Left, val)
		if node.heightDiff() == -2 {
			if node.Right.heightDiff() <= 0 {
				newRoot = node.rotateLeft()
			} else {
				newRoot = node.rotateRightLeft()
			}
		}
	} else {
		node.Right = node.removeFromParent(node.Right, val)
		if node.heightDiff() == 2 {
			if node.Left.heightDiff() >= 0 {
				newRoot = node.rotateRight()
			} else {
				newRoot = node.rotateLeftRight()
			}
		}
	}
	newRoot.computeHeight()
	return newRoot
}

// kiểm tra giá trị v có contains in binary tree
func (bt *AVLTree[T]) Contains(value T) bool {
	return bt.Root.contains(value)
}

// performing kiểm tra giá trị v có contains in binary tree
func (node *BinaryNode[T]) contains(value T) bool {
	n := node
	for n != nil {
		if value < n.val {
			n = n.Left
		} else if value > n.val {
			n = n.Right
		} else {
			return true
		}
	}
	return false
}
