package algorithms

import (
	"fmt"

	"golang.org/x/exp/slices"
)

// Thể hiện ví dụ Dynamic trong chương 3 (cuối chương)
// i dòng, j cột. i biểu diễn cho s1 ( số ký tự chuỗi ban đầu), j biểu diễn cho s2 (số ký tự chuỗi đích).
// s1 là chuỗi ban đầu GCTAC, s2 chuỗi đích CTCA.
// Mục tiêu từ chuỗi s1 chuyển thành chuỗi s2.

// Lập ma trận: m[0][4]
// m[0][4] tức i = 0 nghĩa là chuỗi ban đầu sẽ là 0 (empty string), j = 4 là chuỗi đích, chuỗi đích có 4 ký tự.
// Sau khi thao tác sau thì m[0][4] = 4 nghĩa là thao tác này tốn chi phí là 4

// Tính khoảng tối thiểu
func MinEditDistance(s1, s2 string) [][]int {
	len1 := len(s1)
	len2 := len(s2)

	matrix := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		matrix[i] = make([]int, len2+1)
	}

	// set up chi phí ban đầu cho cột 0 và dòng 0
	// bỏ qua chi phí của matrix[0][0] vì chi phí là 0!
	for i := 1; i <= len1; i++ {
		matrix[i][0] = i
	}
	for j := 1; j <= len2; j++ {
		matrix[0][j] = j
	}

	// tính chi phí tốt nhất (chi phí thấp)
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			cost := 1
			if s1[i-1] == s2[j-1] {
				cost = 0
			}
			replaceCost := matrix[i-1][j-1] + cost
			removeCost := matrix[i-1][j] + 1
			insertCost := matrix[i][j-1] + 1
			matrix[i][j] = Min(replaceCost, removeCost, insertCost)
		}
	}
	return matrix
}

func Min(a, b, c int) int {
	result := a
	if b < result {
		result = b
	}
	if c < result {
		result = c
	}
	return result
}

var (
	REPLACE = 0
	REMOVE  = 1
	INSERT  = 2
)

// Xem quá trình hoạt động của hàm minEdit
func MinEditDistanceOper(s1, s2 string) ([][]int, []string) {
	len1, len2 := len(s1), len(s2)

	matrix, prev := make([][]int, len1+1), make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		matrix[i] = make([]int, len2+1)
		prev[i] = make([]int, len2+1)
	}

	// init prev với giá trị -1
	for i := range prev {
		for j := range prev[i] {
			prev[i][j] = -1
		}
	}

	// set up chi phí ban đầu cho cột 0 và dòng 0
	// bỏ qua chi phí của matrix[0][0] vì chi phí là 0!
	for i := 1; i <= len1; i++ {
		matrix[i][0] = i
	}
	for j := 1; j <= len2; j++ {
		matrix[0][j] = j
	}

	// tính chi phí tốt nhất (chi phí thấp)
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			cost := 1
			if s1[i-1] == s2[j-1] {
				cost = 0
			}
			replaceCost := matrix[i-1][j-1] + cost
			removeCost := matrix[i-1][j] + 1
			insertCost := matrix[i][j-1] + 1

			costs := []int{replaceCost, removeCost, insertCost}
			matrix[i][j] = Min(replaceCost, removeCost, insertCost)
			prev[i][j] = slices.Index(costs, matrix[i][j])
		}
	}

	var ops []string

	i, j := len1, len2
	for i != 0 || j != 0 {
		if prev[i][j] == REMOVE || j == 0 {
			str := fmt.Sprintf("remove %d-th char %c of %s", i, s1[i-1], s1)
			ops = append(ops, str)
			i--
		} else if prev[i][j] == INSERT || i == 0 {
			ops = append(ops, fmt.Sprintf("insert %d-th char %c of %s", j, s2[j-1], s2))
			j--
		} else {
			if matrix[i-1][j-1] < matrix[i][j] {
				str := fmt.Sprintf("replace %d-th char of %s (%c) with %c", i, s1, s1[i-1], s2[j-1])
				ops = append(ops, str)
			}
			i--
			j--
		}
	}
	return matrix, ops
}
