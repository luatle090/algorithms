package massivedatasets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var filename = "one-billion-test.bin"
var n = 200

func TestCreate(t *testing.T) {
	assert := assert.New(t)
	err := Create(filename, n)
	assert.NoError(err, "error create file")
}

func TestSumOneMillion(t *testing.T) {
	assert := assert.New(t)
	sum, err := SumFirstOneMillion(filename, n)
	assert.NotZero(sum, "sum not zero")
	assert.NoError(err, "error create file")
	fmt.Println("sum", sum)
}

func TestSumRandomlyOneMillion(t *testing.T) {
	assert := assert.New(t)
	sum, sumRandom, err := SumRandomlyChosenOneMillion(filename, n)
	assert.NoError(err, "error create file")
	assert.NotZero(sum, "sum not zero")
	assert.Equal(sumRandom, sum, "sum is invalid value")
	fmt.Println("sum", sum)
}
