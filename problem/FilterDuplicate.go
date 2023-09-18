package problem

import (
	"bufio"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func LikeMain() {
	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatalln(err)
	}

	buffReader := bufio.NewReader(file)

	// read 4 byte in chuck (1 chunk là 4 byte)
	buff := make([]byte, 4)

	for {

		_, err := buffReader.Read(buff)
		if err != nil {
			log.Fatalln(err)
		}

		if err == io.EOF {
			log.Println("end of file")
			break
		}

		// đọc tiếp cho đến khi gặp xuống hàng.
		// do buff của chunk ko chắc đã đọc hết 1 line hay chưa.
		nextUntilNewLine, err := buffReader.ReadBytes('\n')
		if err != io.EOF {
			buff = append(buff, nextUntilNewLine...)
		}

	}
}

// Tạo mẫu các dùng gây duplicate dữ liệu
func CreateSampleDupChar(lines int) {

	file, err := os.Create("text.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	// tạo n dòng từ buffer. Buffer mặc định là 4096
	bufWriter := bufio.NewWriter(file)
	for i := 0; i < lines; i++ {
		_, str := CreateRandomCharacter()
		bufWriter.WriteString(str + "\n")
	}

	// check how much is inside waiting to be written
	// fmt.Println(bufWriter.Buffered())

	// check available space left
	// fmt.Println(bufWriter.Available())

	err = bufWriter.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

func CreateRandomCharacter() (int, string) {
	count := 4
	chars := make([]byte, count)

	for i := 0; i < count; i++ {
		c := rand.Intn(122-97) + 97
		chars[i] = byte(c)
	}

	str := string(chars)
	// fmt.Println(str)
	return len(str), str
}

// Cho 1 mảng a chưa sắp xếp. Sắp xếp bằng quick sort sau đó, tìm trùng
// Với giá trị trùng sẽ nằm liền kề nhau
// Hàm trả về danh sách các giá trị ko trùng
func FindDup(a []int) []int {
	// a := []int{1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 5, 6, 7, 8, 8, 9}

	// sort them
	// sort.Ints(a)

	// store các giá trị đã loại trùng
	b := make([]int, 1)
	b[0] = a[0]
	for i := 1; i < len(a); i++ {
		if a[i-1] != a[i] {
			b = append(b, a[i])
		}
	}
	return b
}
