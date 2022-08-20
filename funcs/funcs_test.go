package funcs_test

import (
	"testing"

	"github.com/LukasDeco/slicefuncs/funcs"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	slice := []int{1, 2, 3}
	mapFunc := func(v int) int {
		return v * 2
	}
	newSlice := funcs.Map(slice, mapFunc)

	expectedSlice := []int{2, 4, 6}
	assert.Equal(t, expectedSlice, newSlice)
}

func TestFilter(t *testing.T) {
	slice := []int{1, 2, 3}
	filterFunc := func(v int) bool {
		return v != 2
	}
	newSlice := funcs.Filter(slice, filterFunc)

	expectedSlice := []int{1, 3}
	assert.Equal(t, expectedSlice, newSlice)
}

func TestReduce(t *testing.T) {
	slice := []int{1, 2, 3}
	reduceFunc := func(prev int, curr int) int {
		return prev + curr
	}
	result := funcs.Reduce(slice, reduceFunc, 0)

	expectedResult := 6
	assert.Equal(t, expectedResult, result)
}