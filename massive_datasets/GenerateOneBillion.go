package massivedatasets

import (
	"bufio"
	"encoding/binary"
	"io"
	"math/rand"
	"os"
)

// WriteIntTo write to buff one billion integer (one per line)
func WriteIntTo(writer io.Writer, n int) error {

	// write to default buffer 4096
	buff := bufio.NewWriterSize(writer, 4608)

	b := make([]byte, 8)
	for i := 1; i <= n; i++ {
		binary.BigEndian.PutUint64(b, uint64(i))
		_, err := buff.Write(b)
		if err != nil {
			return err
		}
		// ghi 1 byte la ky tu xuong hang
		buff.WriteString("\n")
		// fmt.Printf("write %d bytes to file\n", n)
	}

	buff.Flush()
	return nil
}

func SumOneMillion(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// sum of the first n million integers
func SumFirstOneMillion(reader io.Reader, n int) (int, error) {

	sum := 0
	i := 0

	// 1024/8 = 128 số integer; mà mỗi số nằm ở 1 dòng => thêm 128 bytes cho ký tự xuống dòng
	// 1024 + 128 = 1152 mảng cần chứa.
	// Để đọc 1 số integer cần slice dataBuff là 9, bao gồm 8 byte số và 1 byte ký tự xuống dòng
	// Tuy nhiên chỉ cần 8
	len := 4096
	lenOfNewLine := len / 8
	buff := bufio.NewReaderSize(reader, len+lenOfNewLine)
	dataBuff := make([]byte, len+lenOfNewLine)
	for {
		nn, err := buff.Read(dataBuff)
		byteRead := 0
		j := 0
		// fmt.Printf("read %d\n", nn)
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}

		for i < n && byteRead < nn {
			offset := j * 9 // Đọc 1 số interger cần slice là 9
			byteRead += 9
			val := binary.BigEndian.Uint64(dataBuff[offset : offset+9])
			// Debug bytes slice
			// if i >= 450 && i <= 470 {
			// 	fmt.Printf("i j offset: %d %d %d ", i, j, offset)
			// 	fmt.Println("val: ", int(val))
			// }
			// if int(val) < 0 || int(val) > 1_000_000 {
			// 	return 0, fmt.Errorf("error parse value %d: i j offset [%d %d %d]", int(val), i, j, offset)
			// }
			sum += int(val)
			i++
			j++
		}
	}
	return sum, nil
}

func SumRandomlyChosenOneMillion(readSeeker io.ReadSeeker, n int) (int, int, error) {

	i, sum, sumRandom := 0, 0, 0
	dataBuff := make([]byte, 8)
	// fmt.Printf("[")
	for i < n {

		chosen := rand.Intn(n)
		sumRandom += chosen + 1
		// fmt.Printf("%d ", chosen+1)
		offset := chosen * 9
		_, err := readSeeker.Seek(int64(offset), 0)
		if err != nil {
			return 0, 0, err
		}
		_, err = readSeeker.Read(dataBuff)
		if err != nil {
			return 0, 0, err
		}

		val := binary.BigEndian.Uint64(dataBuff)
		sum += int(val)

		i++
	}

	// fmt.Printf("]")

	return sum, sumRandom, nil
}

// Đọc trực tiếp từ file.
func SumRandomlyChosenOneMillion2(filename string, n int) (int, int, error) {

	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()
	i, sum, sumRandom := 0, 0, 0
	dataBuff := make([]byte, 8)
	// fmt.Printf("[")
	for i < n {

		chosen := rand.Intn(n)
		sumRandom += chosen + 1
		// fmt.Printf("%d ", chosen+1)
		offset := chosen * 9
		_, err := file.Seek(int64(offset), 0)
		if err != nil {
			return 0, 0, err
		}
		_, err = file.Read(dataBuff)
		if err != nil {
			return 0, 0, err
		}

		val := binary.BigEndian.Uint64(dataBuff)
		sum += int(val)

		i++
	}

	// fmt.Printf("]")

	return sum, sumRandom, nil
}
