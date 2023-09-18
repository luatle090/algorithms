package test

import (
	"fmt"
	"testing"

	"github.com/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert := assert.New(t)
	data := []struct {
		a, b, c  int
		expected int
		errMsg   string
	}{
		{1, 2, 3, 1, "1 = 1"},
		{3, 2, 1, 1, "1 = 1"},
		{4, 1, 6, 1, "1 = 1"},
		{5, 5, 5, 5, "5 = 5"},
		{-1, 5, 0, -1, "-1 = -1"},
		{10, 11, 9, 9, "min is 9"},
		{1, 1, 2, 1, "min is 1"},
		{2, 3, 2, 2, "min is 2"},
		{4, 2, 2, 2, "min is 2"},
		{10, 9, 10, 9, "min is 9"},
	}
	for _, d := range data {
		result := algorithms.Min(d.a, d.b, d.c)
		assert.Equal(d.expected, result, d.errMsg)
	}
	//assert.NotEqual(456, 456, "they should not be equal")
}

func TestDynamic(t *testing.T) {
	fmt.Println(algorithms.MinEditDistance("GCTAC", "CTCA"))
}

func TestDynamicOper(t *testing.T) {
	_, ops := algorithms.MinEditDistanceOper("GCTAC", "CTCA")
	for _, str := range ops {
		fmt.Println(str)
	}
}
