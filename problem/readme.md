
## Problem:
=====================

### Loại bỏ giá trị trùng trong tập tin
---------------------

FilterDuplicate.go 

Cho 1 tập tin có chứa n dòng, mỗi dòng là 1 giá trị. Hãy lọc các giá trị sao cho output chỉ chứa các giá trị phân biệt.

vd: cho 1 chuỗi gồm có 4 ký tự (gọi là k). Với 4 ký tự này được random ngẫu nhiên và có thể trùng lặp như (aaaa, aaab), các ký tự này được lấy trong bảng chữ cái tiếng Anh. 

Vậy theo phép tính thì có 24^k = 24^4 = 331.776 các ký tự có thể được tạo

1 ký tự trong java và golang sẽ lưu là 1 byte (ko theo chuẩn UTF-8). k ký tự là k byte

n * k = dung lượng tập tin
