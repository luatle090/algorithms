package test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/algorithms"
	"github.com/stretchr/testify/assert"
)

var dataSearch = []struct {
	nums            []int
	find            []int
	expectedFind    []bool
	expectedOrdered []int
	msgError        string
}{
	{
		nums:            []int{4, 3, 2, 1, 15, 11, 8, 5},
		find:            []int{4, 15, 14, -1},
		expectedFind:    []bool{true, true, false, false},
		expectedOrdered: []int{1, 2, 3, 4, 5, 8, 11, 15},
		msgError:        "array hasn't been sorted",
	},
	{
		nums:            []int{77, 12, 15, 6, 8, 11, 4, 33, 2, -100, -1, 77},
		find:            []int{-2, 4, 5, -100, 77},
		expectedFind:    []bool{false, true, false, true, true},
		expectedOrdered: []int{-100, -1, 2, 4, 6, 8, 11, 12, 15, 33, 77, 77},
		msgError:        "array hasn't been sorted",
	},
	{
		nums:            []int{11, 1000, 222, 44, 5, -1, 4, 5, 2, 11, 2444, 222, 55, 11, 55},
		find:            []int{11, 233, -55, 100, -95},
		expectedFind:    []bool{true, false, false, false, false},
		expectedOrdered: []int{-1, 2, 4, 5, 5, 11, 11, 11, 44, 55, 55, 222, 222, 1000, 2444},
		msgError:        "array hasn't been sorted",
	},
}

func TestAVL(t *testing.T) {

	assert := assert.New(t)
	for _, d := range dataSearch {
		bt := algorithms.InitAVLTree[int]()
		for _, val := range d.nums {
			bt.AddNode(val)
		}

		bt.Iterator()
		fmt.Println()

		found := make([]bool, len(d.find))
		for i, val := range d.find {
			found[i] = bt.Contains(val)
		}

		assert.Equal(d.expectedFind, found, d.msgError)

		bt.RemoveNode(1)
		bt.Iterator()
	}
}

func TestBinarySearch(t *testing.T) {

	assert := assert.New(t)
	for _, d := range dataSearch {

		found := make([]bool, len(d.find))
		for i, val := range d.find {
			position := algorithms.BinarySearchStandard[int](d.expectedOrdered, val)
			if position != -1 {
				found[i] = true
			} else {
				found[i] = false
			}
		}

		assert.Equal(d.expectedFind, found, d.msgError)
	}
}

func TestFindMaxValue(t *testing.T) {
	for _, d := range dataSearch {
		sort.Sort(sort.Reverse(sort.IntSlice(d.expectedOrdered)))
		position := algorithms.BinarySearchFindMax[int](d.expectedOrdered, 40)
		if position != -1 {
			fmt.Println(d.expectedOrdered[position])
		} else {
			fmt.Println("not found")
		}
	}
}

func TestFindMaxInUnimodal(t *testing.T) {
	assert := assert.New(t)
	a := []int{1, 2, 3, 11, 10, 9}
	b := []int{1, 2, 5, 9, 7, 6, 2}
	find := algorithms.BinarySearchUniModal[int](a)
	assert.Equal(11, find)
	find = algorithms.BinarySearchUniModal[int](b)
	assert.Equal(9, find)
}
