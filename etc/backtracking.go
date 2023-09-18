package etc

import (
	"fmt"
)

var (
	loiGiai        = make([]int, 8)
	cotTrong       = make([]bool, 8)
	cheoXuoiTrong  = make([]bool, 8*2-1)
	cheoNguocTrong = make([]bool, 8*2-1)
	kichThuoc      = 8
)

// In Tám hậu đệ quy
func BacktrackingQueen(j int) {
	for i := 0; i < kichThuoc; i++ {
		if cotTrong[i] && cheoXuoiTrong[j-i+7] && cheoNguocTrong[j+i] {
			loiGiai[j] = i
			if j == kichThuoc-1 {
				fmt.Println(loiGiai)
			} else {
				cotTrong[i] = false
				cheoXuoiTrong[j-i+7] = false
				cheoNguocTrong[j+i] = false
				BacktrackingQueen(j + 1)
				cotTrong[i] = true
				cheoXuoiTrong[j-i+7] = true
				cheoNguocTrong[j+i] = true
			}

		}
	}
}

func KhoiTao() {
	for i := range cotTrong {
		cotTrong[i] = true
	}

	for i := range cheoNguocTrong {
		cheoNguocTrong[i] = true
		cheoXuoiTrong[i] = true
	}
}

var (
	arr    = make([]int, 100)  // mảng lời giải
	choose = make([]bool, 100) // mảng đánh dấu
	t      = make([]int, 100)  // tổng phần tử trong mảng arr
)

// Liệt kê dãy nhị phân độ dài n
// i là biến chạy, n là giới hạn độ dài của dãy
func LietKeNhiPhan(i, n int) {
	// dãy nhị phân có giá trị 0 hoặc 1
	for j := 0; j <= 1; j++ {
		arr[i] = j    // thử đặt giá trị 0 hoặc 1 vào mảng lời giải
		if i == n-1 { // nếu i trong mảng đã bằng n thì dừng
			printLietKe(n)
		} else {
			// thử node a[i] khác trong các node
			LietKeNhiPhan(i+1, n)
		}
	}
}

// Liệt kê chỉnh hợp không lặp chập k
// Dãy n phần tử chọn ra k phần tử, với k phần tử này sẽ khác nhau
// Các phần tử chọn trong dãy là {1,2,3, ..., n}
func LietKeChinhHopKoLap(i, n, k int) {
	for j := 0; j < n; j++ {
		// Chọn các giá trị mà chưa chọn ở node cha
		// Khi đã chọn rồi thì ko cần tìm vì đang tìm giá trị ko lặp
		if choose[j] {
			arr[i] = j + 1
			if i == k {
				printLietKe(k)
			} else {
				choose[j] = false
				LietKeChinhHopKoLap(i+1, n, k)
				choose[j] = true // các node anh em là giá trị j khác nên cần bỏ đánh dấu
			}
		}
	}
}

// xét n nhỏ hơn 30 vì gọi đệ quy.
// Liệt kê các cách để sum giá trị lại. Hoán vị của nhau là 1 cách
// vd: n = 6, ta có 11 cách liệt kê; 1+5; 6; 2+4; 3+3; ....
// 2+4 và 4+2 là 1 cách nên chỉ cần liệt kê 1 trong 2 cách trên
func PhanTichSo(i, n int) {
	for j := arr[i-1]; j <= (n-t[i-1])/2; j++ {
		arr[i] = j
		t[i] = t[i-1] + j
		PhanTichSo(i+1, n)
	}
	arr[i] = n - t[i-1]
	printLietKePhanTichSo(i)
}

func KhoiTaoPhanTichSo() {
	arr[0] = 1
	t[0] = 0
}

func printLietKePhanTichSo(n int) {
	for i := 1; i <= n-1; i++ {
		fmt.Print(arr[i], "+")
	}
	fmt.Print(arr[n])
	fmt.Println()
}

// In phần tử liệt kê trong mảng arr
func printLietKe(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(arr[i])
	}
	fmt.Println()
}
