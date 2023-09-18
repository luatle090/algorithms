package algorithms

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSymmetric(root *TreeNode) {

}

func newNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

// Hàm trả về slice, duyệt theo inorder
func InorderTravesal(root *TreeNode) []int {
	arr := make([]int, 0)
	if root != nil {
		arr = inorderHelper(root, arr)
	}
	return arr
}

func inorderHelper(root *TreeNode, arr []int) []int {
	if root.Left != nil {
		arr = inorderHelper(root.Left, arr)
	}
	arr = append(arr, root.Val)

	if root.Right != nil {
		arr = inorderHelper(root.Right, arr)
	}
	return arr
}
