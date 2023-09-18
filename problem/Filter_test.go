package problem_test

import (
	"testing"

	"github.com/algorithms/problem"
	"github.com/stretchr/testify/assert"
)

func TestRandomCharacter(t *testing.T) {
	_, strLen := problem.CreateRandomCharacter()
	assert.Equal(t, 4, strLen, "len string not equals")
}

func TestLines(t *testing.T) {
	problem.CreateSampleDupChar(1_000_000)

}

func TestLikeMain(t *testing.T) {
	problem.LikeMain()
}

func TestFindDup(t *testing.T) {
	assert := assert.New(t)
	suitTest := []struct {
		arr      []int
		expected []int
	}{
		{
			arr:      []int{1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 5, 6, 7, 8, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			arr:      []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 5, 6, 7, 8, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			arr:      []int{2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 6, 7, 8, 8, 9, 9},
			expected: []int{2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, data := range suitTest {
		b := problem.FindDup(data.arr)
		assert.Equal(data.expected, b)
	}
}
