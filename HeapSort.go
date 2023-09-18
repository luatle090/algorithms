package algorithms

/*
	Heap sort đầu tiên sẽ xây dựng cây thỏa tính chất như trong sách, cây này nhằm mục đích
	tìm giá trị lớn nhất đẩy về vị trí đầu mảng.
	Sau đó sẽ swap vị trí đầu mảng với vị trí cuối mảng rồi xây dựng cây heap tiếp, làm tiếp tục
	cho đến khi tiến tới index là 1.

	Heap sort là sắp xếp ko ổn định khi có mảng trước đã được sắp xếp rồi
	nhưng sort thêm theo tiêu chí khác thì heap sort ko ổn định được tiêu chí đã sắp trước đó


	Heap sort sẽ chia mảng A với n elements thành 2 mảng riêng biệt.
		Từ [0, m), kích thước m là heap với phần tử lớn nhất (hoặc nhỏ nhất) sẽ đầu mảng
		Từ [m, n), là mảng đã được sắp xếp theo tiêu chí

*/

func HeapSort[T any](arr []*T, cmp func(a, b T) int) {
	buildHeap(arr, cmp)
	for i := len(arr) - 1; i >= 1; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i, cmp)
	}
}

func buildHeap[T any](arr []*T, cmp func(a, b T) int) {
	for i := (len(arr)/2 - 1); i >= 0; i-- {
		heapify(arr, i, len(arr), cmp)
	}
}

func heapify[T any](arr []*T, idx, max int, comparator func(a, b T) int) {
	largest := idx
	left := 2*idx + 1
	right := 2*idx + 2

	if left < max && comparator(*arr[left], *arr[idx]) < 0 {
		largest = left
	}
	if right < max && comparator(*arr[right], *arr[largest]) < 0 {
		largest = right
	}
	if largest != idx {
		arr[idx], arr[largest] = arr[largest], arr[idx]
		heapify(arr, largest, max, comparator)
	}
}

// Mảng đc sắp xếp phải thỏa tính chất mọi phần tử a[i] < a[j]
// Nếu có phần tử trùng thì a[i] = a[j] và ko có vị trí k (i < k < j) (nhiều hơn 2 phần tử trùng)
// tạo thành a[i] != a[k]
func IsSorted[T any](arr []*T, cmp func(a, b T) int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if cmp(*arr[i], *arr[i+1]) > 0 {
			return false
		}
	}
	return true
}
