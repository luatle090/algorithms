package test

import (
	"testing"

	"github.com/algorithms"
	"github.com/stretchr/testify/assert"
)

// test for run-length decoding
func TestRLEDecoding(t *testing.T) {
	assert := assert.New(t)
	result := algorithms.Decoding("3e2f5u")
	assert.Equal("eeeffuuuuu", result, "wrong: "+result)
}

func TestRLEEncoding(t *testing.T) {
	assert := assert.New(t)
	result := algorithms.Encoding("eeeffuuuuu")
	assert.Equal("3e2f5u", result, "wrong: "+result)
}
