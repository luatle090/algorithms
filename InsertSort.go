package algorithms

func SortPointer[T any](points []*T, compare func(a, b T) int) {
	for pos := 1; pos < len(points); pos++ {
		i := pos - 1
		value := points[pos]
		for i >= 0 && compare(*points[i], *value) > 0 {
			points[i+1] = points[i]
			i--
		}
		points[i+1] = value
	}
}

// func sortPointerVer2(points []*int, compare func(a, b int) int) {
// 	for pos := 1; pos < len(points); pos++ {
// 		i := pos - 1
// 		value := points[pos]
// 		for i >= 0 && compare(*points[i], *value) > 0 {
// 			points[i+1] = points[i]
// 			i--
// 		}

// 		points[i+1] = value
// 	}
// }

func Comparator(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}
