package test

import (
	"fmt"
	"testing"

	"github.com/algorithms"
	"github.com/stretchr/testify/assert"
)

var (
	Red   = algorithms.Red
	Black = algorithms.Black
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
		expectColor: []bool{Red, Black, Red, Black, Black, Black},
		root:        38,
		msgError:    []string{"must is", "len is"},
	},
	{
		nums:        []int{26, 41, 47, 28, 30, 3, 7, 12, 19, 20, 16, 17, 38, 39, 35},
		expectKey:   []int{3, 7, 12, 16, 17, 19, 20, 26, 28, 30, 35, 38, 39, 41, 47},
		expectColor: []bool{Black, Black, Red, Black, Red, Red, Red, Black, Black, Black, Red, Red, Black, Black, Black},
		root:        28,
		msgError:    []string{"must is", "len is"},
	},
}

var dataRedBlackDelete = []struct {
	nums        []int
	delete      []int
	expectKey   []int
	expectColor []bool
	root        int
	msgError    string
}{
	{
		nums:        []int{41, 38, 31, 12, 19, 8},
		delete:      []int{12, 41},
		expectKey:   []int{8, 19, 31, 38},
		expectColor: []bool{Black, Black, Red, Black},
		root:        19,
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
			if d.expectColor[i] == Black {
				label = "black"
			}
			assert.Equal(d.expectColor[i], arr[i].GetColor(), fmt.Sprintf("%v "+d.msgError[0]+" %s", arr[i].GetKey(), label))
		}

	}
}

func TestDeleteNode(t *testing.T) {
	assert := assert.New(t)

	for _, d := range dataRedBlackDelete {
		tree := algorithms.CreateRedBlackTree[int]()
		for i := range d.nums {
			tree.Add(d.nums[i])
		}

		assert.Equal(len(d.nums), tree.GetSize())

		for i := range d.delete {
			tree.Delete(d.delete[i])
		}

		assert.Equal(d.root, tree.GetRoot().GetKey())

		arr := tree.InorderWalk()

		for i := range d.expectColor {
			assert.Equal(d.expectColor[i], arr[i].GetColor())
		}
	}
}

// tìm element nhỏ thứ i trong mảng ko sắp xếp
func TestK_Element_Smallest(t *testing.T) {
	assert := assert.New(t)
	rankTest := 6
	for _, d := range dataRedBlack {
		tree := algorithms.CreateOrderStatistic[int]()
		for i := range d.nums {
			tree.Add(d.nums[i])
		}

		assert.Equal(len(d.nums), tree.GetSize(), d.msgError[1])

		node := algorithms.OsSelect[int](tree.GetRoot(), rankTest)
		rank := tree.OsRank(node)
		assert.Equal(rankTest, rank, fmt.Sprintf("%v is an element %vth smallest", node.GetKey(), rankTest))
	}
}

// tìm element lớn thứ i trong mảng ko sắp xếp
func TestK_Element_Largest(t *testing.T) {
	assert := assert.New(t)
	rankTest := 3
	for _, d := range dataRedBlack {
		rank := rankTest // reset rank for next loop
		tree := algorithms.CreateOrderStatistic[int]()
		for i := range d.nums {
			tree.Add(d.nums[i])
		}

		assert.Equal(len(d.nums), tree.GetSize(), d.msgError[1])

		// đổi sang largest
		rank = tree.GetSize() - rankTest + 1

		node := algorithms.OsSelect[int](tree.GetRoot(), rank)
		rankNode := tree.OsRank(node)
		fmt.Println("value node", node.GetKey(), ", rank actual in the tree", rankNode)
		assert.Equal(rank, rankNode, fmt.Sprintf("%v is an element %vth largest", node.GetKey(), rankTest))
	}
}
