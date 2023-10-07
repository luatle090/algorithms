package algorithms

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

// // convert string to interger, can convert negative interger
// func string2Integer(s string) int {
// 	result := 0
// 	i := 0
// 	if s[0] == '-' {
// 		i = 1
// 	}
// 	for ; i < len(s); i++ {
// 		digit := int(s[i] - '0')
// 		result = result*10 + digit
// 	}
// 	if s[0] == '-' {
// 		result = -result
// 	}
// 	return result
// }

// convert interger to string
func Int2String(x int) string {
	negative := false
	if x < 0 {
		negative = true
	}
	var result strings.Builder
	// do { ... } while(x != 0)
	for {
		// ascii 0 là 48 hệ 10 rồi cộng với số int hệ 10 là ra, rồi chuyển sang byte
		result.WriteByte(byte('0' + math.Abs(float64(x%10))))
		x /= 10
		if x == 0 {
			break
		}
	}
	if negative {
		result.WriteByte('-')
	}
	return reverse(result.String())
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

// Decoding RLE: "3e2f5u" -> "eeeffuuuuu"
func Decoding(str string) string {
	count := 0
	rns := make([]rune, 0)
	for _, c := range str {
		if unicode.IsDigit(c) {
			count = count*10 + int(c-'0')
		} else {
			for count > 0 {
				rns = append(rns, c)
				count--
			}
		}
	}
	return string(rns)
}

// encoding RLE: "eeeffuuuuu" -> "3e2f5u"
func Encoding(str string) string {
	count := 1
	var rns strings.Builder
	for i := 1; i <= len(str); i++ {
		if i == len(str) || str[i] != str[i-1] {
			x := strconv.FormatInt(int64(count), 32)
			rns.WriteString(x) // skip checking error
			rns.WriteByte(str[i-1])
			count = 1 // reset count to 1
		} else {
			count++
		}
	}
	return rns.String()
}
