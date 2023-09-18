package leetcode

/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
Note that you must do this in-place without making a copy of the array.

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Input: nums = [0]
Output: [0]

Input: nums = [1,23,45,6,7,8,9,1]
Output: [1,23,45,6,7,8,9,1]
*/
func MoveZeros(nums []int) {
	// swap version optimal
	// dùng 2 con trỏ, con trỏ j chạy trước con trỏ i
	var i, j int = 0, 1
	for i < len(nums) && j < len(nums) {
		if nums[i] == 0 && nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			// swap xong thì tăng i và j lên 1
			i++
			j++
		} else if nums[j] == 0 && nums[i] == 0 {
			j++
		} else {
			i++
			j++
		}
	}
}

// swap mọi giá trị, kể cả khi ko cần swap
// với input thứ 3 thì hàm này cũng swap. Xem chi phí swap là nhỏ nên
// => chạy nhanh nhất
func MoveZeros_2(nums []int) {
	currentIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[currentIndex], nums[i] = nums[i], nums[currentIndex]
			currentIndex++
		}
	}
}

func MoveZeros_1(nums []int) {
	//  swap dump
	for i := 1; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == 0 && nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}
}
