package algorithms

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func MergeSort(arr []*int) {
	copyArr := make([]*int, len(arr))
	copy(copyArr, arr)
	mergesort(copyArr, arr, 0, len(arr))
}

func mergesort(A, result []*int, start, end int) {
	if end-start < 2 {
		return
	}
	if end-start == 2 {
		if *result[start] > *result[start+1] {
			result[start], result[start+1] = result[start+1], result[start]
		}
		return
	}

	mid := (end + start) / 2
	mergesort(result, A, start, mid)
	mergesort(result, A, mid, end)

	// merge A left and right of A into result
	i, j := start, mid
	for idx := start; idx < end; idx++ {
		if j >= end || (i < mid && *A[i] < *A[j]) {
			result[idx] = A[i]
			i++
		} else {
			result[idx] = A[j]
			j++
		}
	}
}

/**
* >>>>>>>>>>>>> SORT FILE
 */
func MergeSortFile(filePath string) {

	// open file need sort
	fileOriginal, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileOriginal.Close()

	// create new file with content copy from file original
	fileCopy, err := os.Create("../mergesort-temp.bin")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileCopy.Close()

	// copy content
	err = copyFile(fileCopy, fileOriginal)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileCopy.Seek(0, 0)
	fileOriginal.Seek(0, 0)
	// read all content into memory
	ABuff, err := io.ReadAll(fileOriginal)

	if err != nil {
		fmt.Println(err)
		return
	}
	copyBuff, err := io.ReadAll(fileCopy)

	if err != nil {
		fmt.Println(err)
		return
	}

	// invoke mergesort
	mergeSortFile(copyBuff, ABuff, 0, len(ABuff))

	fileOriginal.Seek(0, 0)
	// write into file
	n, err := fileOriginal.Write(ABuff)
	if err != nil {
		fmt.Println("write error:", err)
		return
	}

	fmt.Println("byte has write:", n)
}

// Hàm thực hiện sort
func mergeSortFile(A, result []byte, start, end int) {
	if end-start < 8 {
		return
	}
	if end-start == 8 {
		var left, right int32
		buf := bytes.NewBuffer(result[start:])
		binary.Read(buf, binary.BigEndian, &left)
		binary.Read(buf, binary.BigEndian, &right)
		if left > right {
			binary.BigEndian.PutUint32(result[start:], uint32(right))
			binary.BigEndian.PutUint32(result[start+4:], uint32(left))
		}
		return
	}

	mid := (end + start) / 8 * 4
	mergeSortFile(result, A, start, mid)
	mergeSortFile(result, A, mid, end)

	// Quá trình merge
	n := 0
	for i, j, idx := start, mid, start; idx < end; idx += 4 {
		Ai := int32(binary.BigEndian.Uint32(A[i:]))
		var Aj int32 = 0

		// Chỉ lấy giá trị A[j] khi j < end
		// (Do trong java khi read thì sẽ  di chuyển con trỏ lên vị trí tiếp theo, điều này tránh đọc tiếp)
		// code dưới thì ko đi tiếp sau khi đọc, chỉ lấy đúng vị trí
		if j < end {
			Aj = int32(binary.BigEndian.Uint32(A[j:]))
		}
		if j >= end || (i < mid && Ai < Aj) {
			binary.BigEndian.PutUint32(result[start+n:], uint32(Ai))
			i += 4
			n += 4
		} else {
			binary.BigEndian.PutUint32(result[start+n:], uint32(Aj))
			j += 4
			n += 4
		}
	}
}

func DecodeFileBinary(path string, numOfElement int) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	bufReader := bufio.NewReader(file)
	buff := make([]byte, 64)
	arrRead := make([]int32, numOfElement)
	i := 0

	// read vào buffered 32 bytes
	// xong bắt đầu decode trong 32 bytes đó ra số int 4 byte
	for {
		n, err := bufReader.Read(buff)

		if err != nil {
			break
		}
		if err == io.EOF {
			fmt.Println("end of file")
			break
		}

		buffTransToInt := bytes.NewReader(buff[:n])
		// decode
		for {
			var v int32
			err = binary.Read(buffTransToInt, binary.BigEndian, &v)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("binary.Read failed:", err)
				break
			}
			arrRead[i] = v
			i++
		}
	}
	fmt.Println(arrRead)
}

// Copy file từ nguồn vào đích
func copyFile(dst, src *os.File) error {
	_, err := io.Copy(dst, src)
	return err
}
