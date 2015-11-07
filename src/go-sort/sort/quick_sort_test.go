package sort

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSortIntArray(t *testing.T) {
	intArray := []int{22,11,6,336,34,56,23,56,4,88,93,1}

	intArray = qsortInts(intArray)

	assert.Equal(t, []int{1,4,6,11,22,23,34,56,56,88,93,336}, intArray)
}