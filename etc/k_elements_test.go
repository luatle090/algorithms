package etc

import (
	"fmt"
	"testing"
)

func comparatorMax(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

func comparatorMin(a, b int) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

func TestK_Elemets(t *testing.T) {
	arr := []int{11, 2, 6, 4, 5, 3, 7, 16, 9, 10, 1, 13, 12, 14, 15, 8}
	x := Find_k_elements[int](arr, 5, func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		} else {
			return 0
		}
	})
	fmt.Print(x)
	arr = []int{11, 2, 6, 4, 5, 3, 7, 16, 9, 10, 1, 13, 12, 14, 15, 8}
	x = Find_k_elements[int](arr, 5, func(a, b int) int {
		if a < b {
			return 1
		} else if a > b {
			return -1
		} else {
			return 0
		}
	})

	fmt.Println(x)
}

func TestK_Elemets_Partition(t *testing.T) {
	k := 3
	arr := []int{11, 2, 6, 5, 20, 16, 7, 3, 9, 10, 1, 13, 12, 14, 15, 8}
	Find_k_elements_Partition(arr, comparatorMax, 0, len(arr)-1, k)
	fmt.Println(arr)
	fmt.Print(arr[len(arr)-k])
	// arr = []int{11, 2, 6, 4, 5, 3, 7, 16, 9, 10, 1, 13, 12, 14, 15, 8}

	// fmt.Println(x)
}

func TestK_Elemets_Priority_Queue(t *testing.T) {
	arr := []int{11, 2, 6, 5, 4, 16, 7, 3, 9, 10, 1, 13, 12, 14, 15, 8}
	x := Find_k_elements_PriorityQueue(arr, 3, byPriorityAscending)
	// print k largest
	fmt.Println(x)

	x = Find_k_elements_PriorityQueue(arr, 4, byPriorityDescending)
	// print k smallest
	fmt.Println(x)
}
