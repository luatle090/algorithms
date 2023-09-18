package test

import (
	"testing"

	"github.com/algorithms"
	"github.com/stretchr/testify/require"
)

func TestLinkedList(t *testing.T) {
	require := require.New(t)

	l := algorithms.LinkedList[int]{}

	require.Equal(true, l.IsEmpty())
	require.NotEqual(true, l.DeleteNode(100))

	l.AddLastNode(100)
	l.AddLastNode(200)
	l.AddLastNode(-100)
	l.AddLastNode(500)

	require.Equal(4, l.Length)

	require.Equal(true, l.Contains(-100))
	require.NotEqual(true, l.Contains(600))

	l.DeleteNode(100)
	require.Equal(false, l.Contains(100))
	require.Equal(3, l.Length)
	require.Equal(true, l.Contains(-100))

	l.DeleteNode(-100)
	require.Equal(false, l.Contains(-100))

	// delete tail
	l.DeleteNode(500)
	require.Equal(false, l.Contains(500))
}
