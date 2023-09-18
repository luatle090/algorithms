package test

import (
	"fmt"
	"testing"

	"github.com/algorithms"
	"github.com/stretchr/testify/require"
)

type cacheData struct {
	key      int
	value    int
	operator string
	expected int
}

func TestLRUCache(t *testing.T) {
	require := require.New(t)
	const put = "put"
	const get = "get"

	suit := []struct {
		capacity int
		suitTest []cacheData
	}{
		{
			capacity: 2,
			suitTest: []cacheData{
				{key: 1, value: 1, operator: put, expected: 0},
				{key: 2, value: 2, operator: put, expected: 0},
				{key: 1, value: 2, operator: get, expected: 1},
				{key: 3, value: 3, operator: put, expected: 0},
				{key: 2, value: 3, operator: get, expected: -1},
				{key: 4, value: 4, operator: put, expected: 0},
				{key: 1, value: 4, operator: get, expected: -1},
				{key: 3, value: 4, operator: get, expected: 3},
				{key: 4, value: 4, operator: get, expected: 4},
			},
		},
		{
			capacity: 4,
			suitTest: []cacheData{
				{key: 1, value: 1, operator: put, expected: 0},
				{key: 2, value: 2, operator: put, expected: 0},
				{key: 3, value: 3, operator: put, expected: 0},
				{key: 4, value: 4, operator: put, expected: 0},
				{key: 5, value: 5, operator: put, expected: 0},
				{key: 4, value: 4, operator: get, expected: 4},
				{key: 6, value: 6, operator: put, expected: 0},
				{key: 3, value: 3, operator: get, expected: 3},
				{key: 4, value: 4, operator: get, expected: 4},
				{key: 7, value: 7, operator: put, expected: 0},
			},
		},
		{
			capacity: 4,
			suitTest: []cacheData{
				{key: 1, value: 1, operator: put, expected: 0},
				{key: 2, value: 2, operator: put, expected: 0},
				{key: 3, value: 3, operator: put, expected: 0},
				{key: 4, value: 4, operator: put, expected: 0},
				{key: 2, value: 2, operator: get, expected: 2},
				{key: 5, value: 5, operator: put, expected: 0},
				{key: 3, value: 6, operator: put, expected: 0},
				{key: 7, value: 7, operator: put, expected: 0},
				{key: 8, value: 7, operator: put, expected: 0},
				{key: 3, value: 4, operator: get, expected: 6},
			},
		},
	}

	for i, s := range suit {
		lru := algorithms.Constructor(s.capacity)
		fmt.Printf("test %d\n", i)
		for _, d := range s.suitTest {
			actual := DoOperat(d.key, d.value, d.operator, &lru)
			require.Equal(d.expected, actual, fmt.Sprintf("%d %d %s", d.key, d.value, d.operator))
		}
	}
}

// func TestFindLowestRank(t *testing.T) {
// 	require := require.New(t)
// 	lru := algorithms.Constructor(4)

// 	lru.Put(1, 1)
// 	lru.Put(2, 2)
// 	lru.Put(3, 3)
// 	lru.Put(4, 4)
// 	lru.Get(2)
// 	lru.Get(3)
// 	lru.Put(5, 5)
// }

func DoOperat(key, value int, op string, lru *algorithms.LRUCache) int {
	switch op {
	case "put":
		lru.Put(key, value)
		return 0
	default:
		return lru.Get(key)
	}
}
