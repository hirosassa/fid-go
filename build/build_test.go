package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombination(t *testing.T) {
	table := calcCombinationTable()
	expect := uint64(0)
	actual := table[0][4]
	assert.Equal(t, expect, actual, "0 choose 4 should be 0")

	expect = uint64(1)
	actual = table[3][3]
	assert.Equal(t, expect, actual, "3 choose 3 should be 1")

	expect = uint64(10)
	actual = table[5][2]
	assert.Equal(t, expect, actual, "5 choose 2 should be 10")
}
