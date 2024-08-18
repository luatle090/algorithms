package massivedatasets

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var filename = "one-billion-test.bin"
var n = 1_000

func TestWriteToFile(t *testing.T) {
	assert := assert.New(t)
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	err = WriteIntTo(file, n)
	assert.NoError(err, "error create file")
}

func TestWriteToBuffer(t *testing.T) {
	assert := assert.New(t)
	buffer := new(bytes.Buffer)
	err := WriteIntTo(buffer, n)
	assert.NoError(err, "error write to buffer")
	assert.NotEmpty(buffer, "buffer is empty")
	assert.Equal(n*9, buffer.Len(), "len of buffer not equal")
}

func TestSumOneMillion(t *testing.T) {
	assert := assert.New(t)
	buff := new(bytes.Buffer)
	WriteIntTo(buff, n)
	sum, err := SumFirstOneMillion(buff, n)
	assert.NotZero(sum, "sum not zero")
	assert.NoError(err, "error create file")
	sumExpected := SumOneMillion(n)
	assert.Equal(sumExpected, sum, "sum sequence not equal")
	fmt.Println("sum", sum)
}

func TestSumRandomlyOneMillion(t *testing.T) {
	assert := assert.New(t)
	buff := new(bytes.Buffer)
	WriteIntTo(buff, n)
	reader := bytes.NewReader(buff.Bytes())
	sum, sumRandom, err := SumRandomlyChosenOneMillion(reader, n)
	assert.NoError(err, "error create file")
	assert.NotZero(sum, "sum not zero")
	assert.Equal(sumRandom, sum, "sum is invalid value")
	fmt.Println("sum", sum)
}
