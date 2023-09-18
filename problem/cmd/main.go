package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {
	//var x int32 = 64
	//fmt.Println(unsafe.Sizeof(x))
	numArr := 1000
	file, _ := os.Create("number.bin")
	defer file.Close()
	// buff is 32 bytes
	bufWriter := bufio.NewWriterSize(file, 32)

	rand.Seed(time.Now().UnixNano())
	// a := make([]int32, numArr)
	// for i := 0; i < numArr; i++ {
	// 	a[i] = rand.Int31n(1000+1000) - 1000
	// }
	// fmt.Println(a)
	a := []int32{4, 3, 2, 1, 15, 7, 8}
	for _, v := range a {
		err := binary.Write(bufWriter, binary.BigEndian, int32(v))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// flush data còn trong buff
	bufWriter.Flush()

	// di chuyển con trỏ file về đầu file
	_, err := file.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	bufReader := bufio.NewReader(file)
	buff := make([]byte, 32)
	arrRead := make([]int32, numArr)
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
}
