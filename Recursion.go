package algorithms

import "fmt"

// f(n) = 1 + 2 + 3 + 4 + ... + n
func CapSoCong(n int) int {
	if n <= 1 {
		return 1
	}
	return CapSoCong(n-1) + n
}

// f(n) = 1 x 2 x 3 x 4 x 5 x ... x n
func GiaiThua(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return GiaiThua(n-1) * n
}

func TailRecursion(x int) {
	if x > 0 {
		fmt.Print(x, " ")

		// gọi đệ quy ở vị trí cuối hàm
		TailRecursion(x - 1)
	}
}

func HeadRecursion(x int) {
	if x > 0 {
		// Gọi đệ quy ở vị trí đầu hàm sau khi đã pass kiểm tra input
		HeadRecursion(x - 1)
		fmt.Print(x, " ")
	}
}

type NhanVien struct {
	Id          int
	LineManager int
	ListNV      []NhanVien
}

// Cây phân cấp tìm line manager của x
func CayPhanCap(id int, lookupNV map[int]NhanVien) {
	if lookupNV[id].LineManager == 0 {
		fmt.Print(lookupNV[id].Id, " --- ")
		return
	}
	fmt.Print(lookupNV[id].Id, " - ", lookupNV[id].LineManager, "  ---  ")
	CayPhanCap(lookupNV[id].LineManager, lookupNV)
}

// Cây phân cấp tìm các đệ của x quản lý trực tiếp và ko trực tiếp
func CayPhanCap_2(id int, listEmployee []NhanVien) {
	for _, emp := range listEmployee {
		CayPhanCap_2(emp.Id, emp.ListNV)
		fmt.Print(emp.Id, " - ", emp.LineManager, "  ---  ")
	}
}

// Cây phân cấp tìm các đệ của x quản lý trực tiếp và ko trực tiếp
// Hai hàm này có cách viết khá giống nhau nhưng cách duyệt sẽ khác nhau.
// Khác biệt giữa CapPhanCap_2 và hàm này là hàm này sẽ duyệt qua toàn bộ mảng
// Đối với CapPhanCap_2 thay vì duyệt qua toàn bộ mảng trong tập ban đầu
// thì CapPhanCap_2 chỉ cần duyệt qua danh sách nhân viên mà nó đang quản lý
func CayPhanCap_3(id int, listEmployee []NhanVien) {
	for _, emp := range listEmployee {
		if emp.LineManager == id {
			CayPhanCap_3(emp.Id, listEmployee)
			fmt.Print(emp.Id, " - ", emp.LineManager, "  ---  ")
		}
	}
}

// Xây dựng quản lý từ cấp quản lý cao nhất, như composite pattern
// trả ra danh sách nhân viên đã thêm vì là truyền pass-by-value nên cần assign lại
func AddNhanVienToLineManager(emp NhanVien, lookupLineManager map[int][]NhanVien) []NhanVien {
	if _, ok := lookupLineManager[emp.Id]; ok {
		emp.ListNV = lookupLineManager[emp.Id]
		for i := range emp.ListNV {
			emp.ListNV[i].ListNV = AddNhanVienToLineManager(emp.ListNV[i], lookupLineManager)
		}
	}
	return emp.ListNV
}

// 731: Tính S(n) = 1^2 + 2^2 + 3^2 + … + (n-1)^2 + n^2
func Sn_731(n int) int {
	if n == 1 {
		return 1
	}

	return Sn_731(n-1) + n*n
}

// 732: S(n) = 1/2 + 1/4 + …  + 1/2n
func Sn_732(n float32) float32 {
	if n == 1 {
		return float32(0.5)
	}
	return Sn_732(n-1) + 1/(2*n)
}

// 736 : Tính S(n) = 1/2 + 2/3 + 3/4 + … + n/(n+1)
func Sn_736(n float32) float32 {
	if n == 1 {
		return float32(0.5)
	}
	return Sn_736(n-1) + n/(n+1)
}

func LuyThua(x, n int) int {
	if n == 0 {
		return 1
	}
	return LuyThua(x, n-1) * x
}

// 741 : Tính S(x,n) = x + x^2 + x^3 + … + x^n
func Sn_741(x, n int) int {
	if n == 1 {
		return x
	}
	s := x
	for i := 2; i <= n; i++ {
		s *= x
	}
	return Sn_741(x, n-1) + s
}

// hàm Sn_741 dùng đệ quy gián tiếp (gọi 1 hàm đệ quy khác)
func Sn_741_2(x, n int) int {
	if n == 1 {
		return x
	}
	return Sn_741(x, n-1) + LuyThua(x, n)
}

// Có 3 cột A, B, C và cột A hiện có n đĩa
// di chuyển n đĩa từ cột A sang cột C
// sao cho: Mỗi lần chỉ di chuyển 1 đĩa
// 			Đĩa lớn hơn (tức có chỉ số n lớn hơn) phải nằm bên dưới
//			Có thể sử dụng các cột A,B,C làm cột trung gian
// vd: 2 đĩa thì đĩa 1 sang B; đĩa 2 sang C; đĩa 1 sang C => xong
// Kết luận: có tổng cộng 2^n - 1 bước đi => O(2^n)
func TowerHaNoi(n int, from, to, aux string) {
	if n == 0 {
		return
	}
	TowerHaNoi(n-1, from, aux, to)

	fmt.Printf("Disk %d moved from %s to %s\n", n, from, to)
	TowerHaNoi(n-1, aux, to, from)
}
