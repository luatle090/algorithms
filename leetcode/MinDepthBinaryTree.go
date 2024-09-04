package leetcode

import "github.com/algorithms"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
* 	Find its minimum depth.
*	The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
*
* 	Note: A leaf is a node with no children.
*
 */
type TreeNode[T int] algorithms.BinaryNode[T]

// Tạo cây suy biến bên phải
func InitTreeNodeTest[T int]() *TreeNode[T] {
	a := algorithms.InitBinaryNode[T](1)

	// ép kiểu từ *Binary sang *TreeNode
	root := (*TreeNode[T])(a)
	// root.Left = algorithms.InitBinaryNode[T](3)
	root.Right = algorithms.InitBinaryNode[T](20)
	root.Right.Right = algorithms.InitBinaryNode[T](4)
	root.Right.Right.Right = algorithms.InitBinaryNode[T](5)
	// root.Right.Left = algorithms.InitBinaryNode[T](15)
	return root
}

// Dùng queue
func MinDepth[T int](root *TreeNode[T]) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode[T], 0)
	queue = enqueue(queue, root)
	// left, right := 0, 0
	level := 0

	for len(queue) > 0 {
		level++
		size := len(queue)
		// duyệt hết các node con của parent
		for size != 0 {
			var node *TreeNode[T]
			node, queue = dequeue(queue)
			size--
			if node.Left == nil && node.Right == nil {
				return level
			}
			if node.Left != nil {
				// ép kiểu từ *Binary sang *TreeNode
				var convertedLeft *TreeNode[T] = (*TreeNode[T])(node.Left)
				queue = enqueue(queue, convertedLeft)
			}
			if node.Right != nil {
				queue = enqueue(queue, (*TreeNode[T])(node.Right))
			}
		}
	}
	return level
}

// dùng đệ quy
func Dfs[T int](root *TreeNode[T]) int {
	if root == nil {
		return 0
	}
	// ép kiểu từ *Binary sang *TreeNode
	var convertedLeft *TreeNode[T] = (*TreeNode[T])(root.Left)
	left := Dfs(convertedLeft)
	right := Dfs((*TreeNode[T])(root.Right))

	if root.Left == nil {
		return right + 1
	}
	if root.Right == nil {
		return left + 1
	}

	if left < right {
		return left + 1
	}
	return right + 1
}

func dequeue[T int](queue []*TreeNode[T]) (*TreeNode[T], []*TreeNode[T]) {
	element := queue[0]
	if len(queue) == 1 {
		return element, []*TreeNode[T]{}
	}
	return element, queue[1:]
}

func enqueue[T int](queue []*TreeNode[T], element *TreeNode[T]) []*TreeNode[T] {
	queue = append(queue, element)
	return queue
}
