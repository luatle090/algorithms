package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/algorithms"
	"github.com/stretchr/testify/assert"
)

var data = []struct {
	nums     []int
	expected []int
	msgError string
}{
	{
		nums:     []int{4, 3, 2, 1, 15, 11, 8, 5},
		expected: []int{1, 2, 3, 4, 5, 8, 11, 15},
		msgError: "array hasn't been sorted",
	},
	{
		nums:     []int{77, 12, 15, 6, 8, 11, 4, 33, 2, -100, -1, 77},
		expected: []int{-100, -1, 1, 2, 4, 6, 8, 11, 12, 15, 33, 77},
		msgError: "array hasn't been sorted",
	},
	{
		nums:     []int{11, 1000, 222, 44, 5, -1, 4, 5, 2, 11, 2444, 222, 55, 11, 55},
		expected: []int{1},
		msgError: "array hasn't been sorted",
	},
}

func TestIsSorted(t *testing.T) {
	for _, d := range data {
		arr := make([]*int, len(d.expected))
		for i := range d.expected {
			arr[i] = &d.expected[i]
		}
		sorted := algorithms.IsSorted(arr, algorithms.Comparator)
		assert.Equal(t, true, sorted, d.msgError)
	}
}

func TestInsertionSort(t *testing.T) {
	assert := assert.New(t)

	for _, d := range data {
		arr := make([]*int, len(d.nums))
		for i := range d.nums {
			arr[i] = &d.nums[i]
		}
		algorithms.SortPointer(arr, algorithms.Comparator)
		sorted := algorithms.IsSorted(arr, algorithms.Comparator)
		afterSorted := make([]int, len(d.nums))
		for i := range arr {
			afterSorted[i] = *arr[i]
		}
		fmt.Println(afterSorted)
		assert.Equal(true, sorted, d.msgError)
	}
}

func TestHeapSort(t *testing.T) {
	assert := assert.New(t)

	for _, d := range data {
		arr := make([]*int, len(d.nums))
		for i := range d.nums {
			arr[i] = &d.nums[i]
		}
		algorithms.HeapSort(arr, algorithms.Comparator)
		sorted := algorithms.IsSorted(arr, algorithms.Comparator)
		afterSorted := make([]int, len(d.nums))
		for i := range arr {
			afterSorted[i] = *arr[i]
		}
		fmt.Println(afterSorted)
		assert.Equal(true, sorted, d.msgError)
	}
}

func TestHeapSortRandomList(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	v := rand.Perm(4096)
	arr := toPointerArray(v)
	algorithms.HeapSort(arr, algorithms.Comparator)
	sorted := algorithms.IsSorted(arr, algorithms.Comparator)
	assert.Equal(t, true, sorted, "not sorted")

	v2 := make([]int, 7000)
	for i := 0; i < len(v2); i++ {
		v2[i] = rand.Intn(200-100) - 100
	}
	arr = toPointerArray(v2)
	algorithms.HeapSort(arr, algorithms.Comparator)
	sorted = algorithms.IsSorted(arr, algorithms.Comparator)
	assert.Equal(t, true, sorted, "not sorted 2")
}

func toPointerArray(a []int) []*int {
	arr := make([]*int, len(a))
	for i := range a {
		arr[i] = &a[i]
	}
	return arr
}

func TestMergeSortFile(t *testing.T) {
	algorithms.MergeSortFile("../number.bin")
	algorithms.DecodeFileBinary("../number.bin", 1000)
}

func TestMergsort(t *testing.T) {
	assert := assert.New(t)

	for _, d := range data {
		arr := make([]*int, len(d.nums))
		for i := range d.nums {
			arr[i] = &d.nums[i]
		}
		algorithms.MergeSort(arr)
		sorted := algorithms.IsSorted(arr, algorithms.Comparator)
		afterSorted := make([]int, len(d.nums))
		for i := range arr {
			afterSorted[i] = *arr[i]
		}
		fmt.Println(afterSorted)
		assert.Equal(true, sorted, d.msgError)
	}
}
