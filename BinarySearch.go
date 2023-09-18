package algorithms

// Hàm tìm kiếm binary tiêu chuẩn với mảng phải được sắp xếp để tìm kiếm.
// Hàm trả về vị trí, ko tìm thấy trả về -1
func BinarySearchStandard[T Ordered](a []T, t T) int {
	if len(a) == 0 {
		return -1
	}
	from, to := 0, len(a)-1
	for from <= to {
		mid := (from + to) / 2
		if t < a[mid] {
			to = mid - 1
		} else if t > a[mid] {
			from = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// Hàm tìm kiếm nhị phân, tìm phần tử max gần hoặc nhỏ hơn chặn trên trong mảng.
// Mảng a là mảng thứ tự giảm dần
//
// Return chỉ số nếu thỏa; ngược lại, return -1 nếu ko thỏa.
func BinarySearchFindMax[T Ordered](a []T, max T) int {
	if len(a) == 0 {
		return -1
	}
	from, to := 0, len(a)-1
	// quét hết mảng, ko break trong mảng
	for from < to {
		mid := (from + to) / 2
		if a[mid] > max { // giá trị cần tìm ko xuất hiện từ from đến mid-1 => tìm từ mid+1 đến to
			from = mid + 1
		} else { // giá trị cần tìm ko xuất hiện từ mid+1 đến to => tìm từ from đến mid
			to = mid
		}
	}
	// tìm cho đến khi còn 1 phần tử thì xét phần tử này với max
	if a[from] <= max {
		return from
	}
	return 0
}

// Hàm tìm kiếm nhị phân, mảng a chia thành 2 phần: phần đầu là dãy tăng dần, phần sau là dãy giảm dần
// a[i] < a[i+1] với 0 <= i < m
// a[i] > a[i+1] với m <= i < n
// => a[m] là phần tử lớn nhất trong mảng
// vd: [1,4,8,9,7,6,2] là mảng unimodal, [1,3,5,4,1,2] ko thỏa điều kiện
// Yêu cầu: Tìm max element trong mảng unimodal này
// Nhận xét: đi tìm cực đại của mảng nếu, mảng này vẽ ra đồ thị thì có biểu diễn là hình chuông
// Nhận xét: Đọc trang 359 sách Kỹ thuật lập trình của trường
func BinarySearchUniModal[T Ordered](a []T) T {
	from, to := 0, len(a)-1
	for from < to {
		mid := (from + to) / 2
		if a[mid] > a[mid+1] {
			to = mid
		} else {
			from = mid + 1
		}
	}
	return a[from]
}
