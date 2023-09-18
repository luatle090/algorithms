package algorithms

import "fmt"

// invoke the recursion
func MaxElement(vals []int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no max element in emty array")
	}

	return Element(vals, 0, len(vals))
}

// compute element in sub vals[left, right)
func Element(vals []int, left, right int) (int, error) {
	if right-left == 1 {
		return vals[left], nil
	}

	mid := (left + right) / 2
	max1, _ := Element(vals, left, mid)
	max2, _ := Element(vals, mid, right)
	if max1 > max2 {
		return max1, nil
	}
	return max2, nil
}
