package test

import (
	"fmt"
	"testing"

	"github.com/algorithms"
	"github.com/stretchr/testify/assert"
)

var dataRedBlack = []struct {
	nums        []int
	expectKey   []int
	expectColor []bool
	root        int
	msgError    []string
}{
	{
		nums:        []int{41, 38, 31, 12, 19, 8},
		expectKey:   []int{8, 12, 19, 31, 38, 41},
		expectColor: []bool{algorithms.Red, algorithms.Black, algorithms.Red, algorithms.Black, algorithms.Black, algorithms.Black},
		root:        38,
		msgError:    []string{"must is", "len is"},
	},
}

func TestAddNode(t *testing.T) {
	assert := assert.New(t)
	for _, d := range dataRedBlack {
		tree := algorithms.CreateRedBlackTree[int]()
		for i := range d.nums {
			tree.Add(d.nums[i])
		}

		assert.Equal(len(d.nums), tree.GetSize(), d.msgError[1])

		// test root
		assert.Equal(d.root, tree.GetRoot().GetKey(), fmt.Sprintf("root must is %v", d.root))

		arr := tree.InorderWalk()

		// testing inoreder walk
		for i := range d.expectKey {
			assert.Equal(d.expectKey[i], arr[i].GetKey(), fmt.Sprintf("expected is %v but actual is %v", d.expectKey[i], arr[i].GetKey()))
		}

		// testing color
		for i := range d.expectColor {
			label := "red"
			if d.expectColor[i] == algorithms.Black {
				label = "black"
			}
			assert.Equal(d.expectColor[i], arr[i].GetColor(), fmt.Sprintf("%v "+d.msgError[0]+" %s", arr[i].GetKey(), label))
		}

	}
}
