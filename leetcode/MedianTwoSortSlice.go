package leetcode

import "fmt"

/**
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
The overall run time complexity should be O(log (m+n)).

Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
**/

/*
test case leetcode. 1 cặp slice gồm 2 dòng
[1,3]
[2] // expected 2.0
[1,2]
[3,4] // expected 2.5
[1,2]
[2,3] // expected 2.0
[1,3,5,7]
[2,4]  // 3.5
[1,3,4,7]
[1,3,5,7] // 3.5
[5,8,10]
[2,9]    // 8.0
[1,2,4]
[]    // 2.0
[]
[1] // 1.0
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// nếu là mảng chẵn, thì lấy 2 phần tử giữa cộng lại chia 2
	// nếu là mảng lẽ, lấy phần tử ở giữa

	// NAIVE
	naiveSolution(nums1, nums2)
	// naiveSolution2(nums1, nums2)
	// twoPointer(nums1, nums2)
	return 0
}

// solution for naive
// compare giữa 2 array. Giá trị của array nào nhỏ nhất thì đưa vào mergeSlice. Cứ làm tiếp tục cho đến khi hết 1 trong 2 slice
// Đối với slice còn lại chưa hết thì cần phải chạy thêm 1 bước để đưa toàn bộ phần tử còn lại vào mergeSlice
// Sau khi xong thì check xem mảng chẵn hay lẻ để tính median
// run time complexity O(n).
func naiveSolution(nums1 []int, nums2 []int) float64 {
	lenOfCombie := len(nums1) + len(nums2)
	lenNum1 := len(nums1)
	lenNum2 := len(nums2)
	mergeSlice := make([]int, lenOfCombie)
	i, j, k := 0, 0, 0
	for i = 0; j < lenNum1 && k < lenNum2; i++ {
		if nums1[j] <= nums2[k] {
			mergeSlice[i] = nums1[j]
			j++
		} else if nums2[k] <= nums1[j] {
			mergeSlice[i] = nums2[k]
			k++
		}
	}

	fmt.Printf("i j k: %d, %d, %d\n", i, j, k)

	addAll := func(nums []int, start int) {
		fmt.Printf("i: %d\n", i)
		for ; start < len(nums); start++ {
			mergeSlice[i] = nums[start]
			i++
		}
		// fmt.Println("in add all", mergeSlice)
	}

	if j < lenNum1 {
		addAll(nums1, j)
	} else if k < lenNum2 {
		addAll(nums2, k)
	}

	// slice is even
	fmt.Println("mergeSlice", mergeSlice, "len mergeSlice", len(mergeSlice))
	if len(mergeSlice)%2 == 0 {
		mid := len(mergeSlice) / 2
		return float64(mergeSlice[mid]+mergeSlice[mid-1]) / 2.0
	} else {
		// slide is odd
		mid := len(mergeSlice) / 2
		return float64(mergeSlice[mid])
	}
}

// Cải thiện từ naive 1
// Thay vì merge toàn bộ 2 mảng vào 1 thì ta chỉ merge 1 nửa, đúng bằng ((n+m)/2) + 1
// Lưu ý mảng merge là chẵn hay lẻ thì khác nhau.
// O(m+n), O(m+n)
func naiveSolution2(nums1 []int, nums2 []int) float64 {
	// NAIVE
	lenOfCombie := len(nums1) + len(nums2)
	lenNum1 := len(nums1)
	lenNum2 := len(nums2)
	halfLenCombie := lenOfCombie / 2
	mergeSlice := make([]int, halfLenCombie+1)
	j, k := 0, 0
	for i := 0; i <= halfLenCombie; i++ {
		if j < lenNum1 && k < lenNum2 {
			if nums1[j] <= nums2[k] {
				mergeSlice[i] = nums1[j]
				j++
			} else {
				mergeSlice[i] = nums2[k]
				k++
			}
		} else if j < lenNum1 {
			mergeSlice[i] = nums1[j]
			j++
		} else {
			mergeSlice[i] = nums2[k]
			k++
		}
	}

	fmt.Printf("j k: %d, %d\n", j, k)

	// slice is even
	fmt.Println("mergeSlice", mergeSlice, ",len mergeSlice", len(mergeSlice), ",len total of two array", lenOfCombie)
	if lenOfCombie%2 == 0 {
		return float64(mergeSlice[halfLenCombie]+mergeSlice[halfLenCombie-1]) / 2.0
	} else {
		// slide is odd
		return float64(mergeSlice[halfLenCombie])
	}
}

// two pointer là di chuyển index của 2 slice,
// cần lưu lại value và pre-value của mỗi lần di chuyển
// O(m+n), space O(1)
func twoPointer(nums1 []int, nums2 []int) float64 {
	i, j, val, preVal := 0, 0, 0, 0
	lenOfCombie := len(nums1) + len(nums2)
	halfLenOfCombie := lenOfCombie / 2
	for count := 0; count <= halfLenOfCombie; count++ {
		preVal = val
		// cần đảm bảo ko bị out of bound của slice
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] <= nums2[j] {
				val = nums1[i]
				i++
			} else {
				val = nums2[j]
				j++
			}
		} else if i < len(nums1) {
			val = nums1[i]
			i++
		} else {
			val = nums2[j]
			j++
		}
	}
	// slice is even
	if lenOfCombie%2 == 0 {
		return float64(val+preVal) / 2.0
	} else {
		// slide is odd
		return float64(val)
	}
}

// binary search returns an index of slice
