package etc

import (
	"fmt"
	"math"

	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

// Tìm phần tử thứ k lớn nhất trong mảng
// vd : [1,8,2,3,6,7,5,4] => tìm phần tử thứ 2 lớn nhất sẽ là 7
// tìm phần tử lớn thứ 3 lớn nhất sẽ là 6
func Find_k_elements[T any](arr []T, k int, comparator func(a, b T) int) T {
	// Sử dụng heapify O(n Log(n))
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify[T](arr, i, len(arr), comparator)
	}
	fmt.Println(arr)
	for i := len(arr) - 1; i >= len(arr)-k+1; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify[T](arr, 0, i, comparator)
	}
	fmt.Println(arr)
	return arr[0]
}

func heapify[T any](arr []T, idx, max int, comparator func(a, b T) int) {
	largest := idx
	left := 2*idx + 1
	right := 2*idx + 2

	if left < max && comparator(arr[left], arr[idx]) > 0 {
		largest = left
	}
	if right < max && comparator(arr[right], arr[largest]) > 0 {
		largest = right
	}
	if largest != idx {
		arr[idx], arr[largest] = arr[largest], arr[idx]
		heapify(arr, largest, max, comparator)
	}
}

// Khó hiểu!!!?
func Find_k_elements_Partition(arr []int, cmp func(a, b int) int, left, right, k int) int {
	// Sử dụng partition quicksort O(n log(n))
	pivotIndex := partition(arr, cmp, left, right, left)

	// if ...
	if pivotIndex-left == k-1 {
		return arr[pivotIndex]
	}

	if pivotIndex-left > k-1 {
		Find_k_elements_Partition(arr, cmp, left, pivotIndex-1, k)
	} else {
		Find_k_elements_Partition(arr, cmp, pivotIndex+1, right, k-pivotIndex+left-1)
	}
	return math.MaxInt
}

// partition lấy từ quicksort
func partition(arr []int, cmp func(a, b int) int, left, right, pivotIndex int) int {
	pivot := arr[pivotIndex]

	// mục đích để lưu vị trí pivot, vị trí lưu sẽ là cuối mảng
	// move pivot to the end of the array (swap)
	arr[right], arr[pivotIndex] = arr[pivotIndex], arr[right]

	// all values <= pivot are moved to front of array and pivot inserted
	// just after them
	store := left
	for idx := left; idx < right; idx++ {
		if cmp(arr[idx], pivot) <= 0 {
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

func Find_k_elements_PriorityQueue(arr []int, k int, priority func(a, b interface{}) int) int {
	// Sử dụng priority queue O(n Log(k))
	queue := priorityqueue.NewWith(priority)

	for _, v := range arr {
		el := Element{name: v, priority: v}
		queue.Enqueue(el)

		if queue.Size() > k {
			queue.Dequeue()
		}
	}
	v, _ := queue.Dequeue()
	return v.(Element).name
}

type Element struct {
	name     int
	priority int
}

func byPriorityAscending(a, b interface{}) int {
	priorityA := a.(Element).priority
	priorityB := b.(Element).priority
	return -utils.IntComparator(priorityB, priorityA) // "-" descending order
}

func byPriorityDescending(a, b interface{}) int {
	priorityA := a.(Element).priority
	priorityB := b.(Element).priority
	return -utils.IntComparator(priorityA, priorityB) // "-" descending order
}
