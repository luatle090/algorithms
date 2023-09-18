package algorithms

import (
	"math/rand"
	"time"
)

var MinSize = 0

/*
* Trong thời gian tuyến tính, nhóm mảng con arr[left, right] xung quanh pivot pivot=arr[pivotIndex]
* bằng cách lưu pivot vào vị trí thích hợp, trong mảng con (có vị trí được trả về bởi hàm này)
* và đảm bảo tất cả arr[left, store) <= pivot và tất cả arr[store + 1, right] > pivot.
*
* In linear time, group the subarray ar[left, right] around a pivot
* element pivot=ar[pivotIndex] by storing pivot into its proper
* location, store, within the subarray (whose location is returned
* by this function) and ensuring all ar[left,store) <= pivot and
* all ar[store+1,right] > pivot.
 */
func Partition(arr []*int, cmp func(a, b int) int, left, right, pivotIndex int) int {
	pivot := arr[pivotIndex]

	// mục đích để lưu vị trí pivot, vị trí lưu sẽ là cuối mảng
	// move pivot to the end of the array (swap)
	arr[right], arr[pivotIndex] = arr[pivotIndex], arr[right]

	// all values <= pivot are moved to front of array and pivot inserted
	// just after them
	store := left
	for idx := left; idx < right; idx++ {
		if cmp(*arr[idx], *pivot) <= 0 {
			arr[idx], arr[store] = arr[store], arr[idx]
			store++
		}
	}

	// fmt.Printf("after swap %v", arr)

	// Đặt pivot vào vị trí giữa để đảm bảo
	// Put pivot in the central to ensuring all arr[left, store) <= pivot
	// and all arr[store+1, right] > pivot
	arr[right], arr[store] = arr[store], arr[right]
	return store // return center
}

// return index of pivot has selected
func SelectPivotFirst(arr []*int, left, right int) int {
	return left
}

func SelectPivotLast(arr []*int, left, right int) int {
	return right
}

func SelectPivotRandom(arr []*int, left, right int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(right-left) + left
}

func SelectPivotMedian(arr []*int, left, right int) int {
	mid := (right - left + 1) / 2

	for {
		idx := selectMedianIndex(arr, Comparator, left, right)

		pivotIndex := Partition(arr, Comparator, left, right, idx)
		if left+mid-1 == pivotIndex {
			return pivotIndex
		}
		if left+mid-1 < pivotIndex {
			right = pivotIndex - 1
		} else {
			mid = mid - pivotIndex - left + 1
			left = pivotIndex + 1
		}
	}
}

// Select median of three. Select median element of vals[left], vals[mid], and vals[right] to use.
func selectMedianIndex(arr []*int, cmp func(a, b int) int, left, right int) int {
	choices := make([]*int, 3)
	choices[0] = arr[left]
	mid := (left + right + 1) / 2
	choices[1] = arr[mid]
	choices[2] = arr[right]

	if cmp(*choices[0], *choices[1]) < 0 {
		if cmp(*choices[1], *choices[2]) <= 0 {
			return mid
		} else if cmp(*choices[0], *choices[1]) < 0 {
			return right
		} else {
			return left
		}
	}
	if cmp(*choices[0], *choices[2]) < 0 {
		return left
	} else if cmp(*choices[1], *choices[2]) < 0 {
		return right
	} else {
		return mid
	}
}

func QuickSort(arr []*int, cmp func(a, b int) int,
	selectPivot func(arr []*int, left, right int) int) {

	do_qsort(arr, cmp, 0, len(arr)-1, selectPivot)
}

func do_qsort(arr []*int,
	cmp func(a, b int) int,
	left, right int,
	selectPivotIndex func(arr []*int, left, right int) int) {

	if right <= left {
		return
	}

	// partition
	pivotIndex := selectPivotIndex(arr, left, right)
	pivotIndex = Partition(arr, cmp, left, right, pivotIndex)

	// call insertion sort
	if pivotIndex-1-left <= MinSize {
		insertion(arr, cmp, left, pivotIndex-1)
	} else {
		do_qsort(arr, cmp, left, pivotIndex-1, selectPivotIndex)
	}

	if right-pivotIndex-1 <= MinSize {
		insertion(arr, cmp, left, pivotIndex-1)
	} else {
		do_qsort(arr, cmp, pivotIndex+1, right, selectPivotIndex)
	}
}

func insertion(arr []*int, cmp func(a, b int) int, low, high int) {
	for pos := low + 1; pos <= high; pos++ {
		j := pos - 1
		value := arr[pos]
		for j >= 0 && cmp(*arr[pos], *value) > 0 {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = value
	}
}
