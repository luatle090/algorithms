package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{123, 456, 123, 123, 789, 456, 789, 789, 789, 789}
	timUngVien(a, len(a))
	timUngVien_Non_Efficient(a, len(a))
}

// mảng a chứa ID danh sách ứng viên được bầu,
// n là số cử tri đã bầu
// hàm trả ra kết quả ID người chiến thắng
// vd: [123, 456, 123, 123, 789, 456, 789, 123, 789, 123]
// ứng viên chiến thắng là 123.
func timUngVien(a []int, n int) int {

	// Big O là O(n) về mặt thời gian
	// Big O là O(n) về mặt ko gian chứa do cần thêm hashmap

	mapCount := make(map[int]int)

	var winner, maxCount int
	maxCount = 0

	for _, candidate := range a {
		if _, ok := mapCount[candidate]; !ok {
			mapCount[candidate] = 1
		} else {
			mapCount[candidate]++
		}

		// update vị trí chiến thắng của ứng viên qua mỗi lần lăp
		// get count from ID
		if mapCount[candidate] > maxCount {
			winner = candidate
			maxCount = mapCount[candidate]
		}
	}

	fmt.Print(mapCount)

	fmt.Println("ung vien chien thang ", winner)
	return winner
}

// Cách tiếp cận sort và tìm tuyến tính
func timUngVien_Non_Efficient(a []int, n int) int {

	// Vì dùng quicksort nên BigO là O(nLog(n))
	// BigO không gian chỉ là O(1)
	// thích hợp cho array có data lớn

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	// sort.Sort(sort.IntSlice(a))

	// fmt.Println(a)

	currCount := 1
	max := 0
	var winner int
	for i := 1; i < n; i++ {
		if a[i] == a[i-1] {
			currCount++
		} else {
			currCount = 1
		}

		if currCount > max {
			max = currCount
			winner = a[i]
		}
	}

	fmt.Println("Ung vien chien thang la ", winner)
	return winner
}
