package algorithms

/// ------ Hàm mục đích dùng cho leetcode - Test stack và Queue trong package test ------

/// --------  Queue ---------
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

/// -------- Stack ----------
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

// Returns the top element without removing it
func PeekStack(stack []int) int {
	if len(stack) == 0 {
		return 0
	}
	return stack[len(stack)-1]
}

// need review: not test yet. Returns a slice from bottom stack (first index) to top stack (last index)
func PopStackToSlice(stack []int) []int {
	return stack[:]
}

/// ----------------------
