package test

import (
	"testing"

	"github.com/algorithms"
	"github.com/stretchr/testify/assert"
)

var data_2 = []struct {
	nums          []int
	expectedQueue []int
	expectStack   []int
	msgError      string
}{
	{
		nums:          []int{4, 3, 2, 1, 15, 11, 8, 5},
		expectedQueue: []int{4, 3, 2, 1, 15, 11, 8, 5},
		expectStack:   []int{5, 8, 11, 15, 1, 2, 3, 4},
		msgError:      "queue hasn't been",
	},
	{
		nums:          []int{77, 12, 15, 6, 8, 11, 4, 33, 2, -100, -1, 77},
		expectedQueue: []int{77, 12, 15, 6, 8, 11, 4, 33, 2, -100, -1, 77},
		expectStack:   []int{77, -1, -100, 2, 33, 4, 11, 8, 6, 15, 12, 77},
		msgError:      "queue hasn't been",
	},
	{
		nums:          []int{11, 1000, 222, 44, 5, -1, 4, 5, 2, 11, 2444, 222, 55, 11, 55},
		expectedQueue: []int{11, 1000, 222, 44, 5, -1, 4, 5, 2, 11, 2444, 222, 55, 11, 55},
		expectStack:   []int{55, 11, 55, 222, 2444, 11, 2, 5, 4, -1, 5, 44, 222, 1000, 11},
		msgError:      "queue hasn't been",
	},
}

func TestQueue(t *testing.T) {
	assert := assert.New(t)
	queue := make([]int, 0)
	for _, d := range data_2 {
		for _, n := range d.nums {
			queue = algorithms.EnQueue(queue, n)
		}

		// test enqueue
		assert.Equal(len(d.nums), len(queue), d.msgError)

		//test dequeue
		for _, n := range d.expectedQueue {
			value, queueReturn := algorithms.DeQueue(queue)
			assert.Equal(n, value, d.msgError)
			queue = queueReturn
		}

		// queue = queue[:0]
	}
}

func TestStack(t *testing.T) {
	assert := assert.New(t)
	stack := make([]int, 0)

	for _, d := range data_2 {
		for _, n := range d.nums {
			stack = algorithms.PushStack(stack, n)
		}

		// test push stack
		assert.Equal(len(d.nums), len(stack), d.msgError)

		// test pop stack
		for _, n := range d.expectStack {
			value, stackReturn := algorithms.PopStack(stack)
			assert.Equal(n, value, d.msgError)
			stack = stackReturn
		}
	}
}
