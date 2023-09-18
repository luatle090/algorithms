package leetcode

import "fmt"

// O(n) - time; O(n) - space
func CanConstruct(ransomNote string, magazine string) bool {
	s := make(map[rune]int)
	for _, val := range magazine {
		if _, ok := s[val]; !ok {
			s[val] = 0
		}
		s[val]++
	}

	n, i := len(ransomNote), 0
	for _, val := range ransomNote {
		if _, ok := s[val]; ok && s[val] > 0 {
			s[val]--
			i++
		}
	}
	if n == i {
		fmt.Print(true)
		return true
	}
	fmt.Print(false)
	return false
	// fmt.Print(result)
}

// optimization
func CanConstruct_2(ransomNote string, magazine string) bool {

	return false
}
