package algorithms

/// ------ Hàm mục đích dùng cho leetcode - Test stack và Queue trong package test ------

func DeQueue(queue []int) (int, []int) {
	element := queue[0]
	if len(queue) == 1 {
		return element, []int{}
	}
	return element, queue[1:]
}

func EnQueue(queue []int, element int) []int {
	queue = append(queue, element)
	return queue
}

func PushStack(stack []int, element int) []int {
	stack = append(stack, element)
	return stack
}

func PopStack(stack []int) (int, []int) {
	if len(stack) == 0 {
		return 0, []int{}
	}
	lastIndex := len(stack) - 1
	return stack[lastIndex], stack[:lastIndex]
}

/// ----------------------
