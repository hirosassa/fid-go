package fid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSize(t *testing.T) {
	bits := New()

	bits.SetSlice(0, 7, uint64(0))
	expected := uint64(7)
	actual := bits.Size()
	assert.Equal(t, expected, actual, "bits size should be 7")

	bits.SetSlice(64, 23, uint64(0))
	expected = uint64(87)
	actual = bits.Size()
	assert.Equal(t, expected, actual, "bits size should be 87")
}

func TestGetBit(t *testing.T) {
	bits := New()
	bits.SetSlice(0, 6, uint64(1<<5))
	expected := true
	actual, _ := bits.GetBit(6)
	assert.Equal(t, expected, actual, "bit at 6 should be true")

	expected = false
	actual, _ = bits.GetBit(2)
	assert.Equal(t, expected, actual, "bit at 2 should be true")

	bits.SetSlice(64, 6, uint64(1<<5))
	expected = true
	actual, _ = bits.GetBit(70)
	assert.Equal(t, expected, actual, "bit at 70 should be true")

	expected = false
	actual, err := bits.GetBit(100)
	assert.Equal(t, expected, actual)
	assert.Error(t, err, "when the i is larger than the size of bits, should returns error")
}

func TestGetSlice(t *testing.T) {
	bits := New()
	bits.SetSlice(0, 7, uint64(1<<6))
	expected := uint64(64)
	actual, _ := bits.GetSlice(0, 7)
	assert.Equal(t, expected, actual, "slice at range(0, 6) should be %d", expected)
	

	expected = uint64(0)
	actual, err := bits.GetSlice(0, 10)
	assert.Equal(t, expected, actual)
	assert.Error(t, err, "when the sliceSize larger than the size of bits, should returns error")

	expected = uint64(0)
	actual, err = bits.GetSlice(0, 123)
	assert.Equal(t, expected, actual)
	assert.Error(t, err, "when the sliceSize larger than the BLOCK_SIZE, should return error")

	expected = uint64(0)
	actual, err = bits.GetSlice(12, 0)
	assert.Equal(t, expected, actual)
	assert.Equal(t, nil, err, "when the sliceSize is 0, should return 0, nil")

	bits.SetSlice(128, 7, uint64(1<<6))
	expected = uint64(1 << 8)
	actual, _ = bits.GetSlice(126, 9)
	assert.Equal(t, expected, actual, "slice at range(126, 134) should be %d", expected)
}

func TestSetBit(t *testing.T) {
	bits := New()
	bits.SetBit(4, false)
	expected := false
	actual, _ := bits.GetBit(4)
	assert.Equal(t, expected, actual, "bit at 4 should be false")

	bits.SetBit(6, true)
	expected = true
	actual, _ = bits.GetBit(6)
	assert.Equal(t, expected, actual, "bit at 6 should be true")
}

func TestSetSlice(t *testing.T) {
	bits := New()
	bits.SetSlice(0, 8, uint64((1<<7)-1))
	expected := uint64((1 << 7) - 1)
	actual, _ := bits.GetSlice(0, 8)
	assert.Equal(t, expected, actual, "slice at range(0, 8) should be %d", expected)

	expected = uint64((1 << 3) - 1)
	actual, _ = bits.GetSlice(4, 4)
	assert.Equal(t, expected, actual, "slice at range(4, 8) should be %d", expected)

	bits.SetSlice(63, 8, uint64((1<<7)-1))
	expected = uint64((1 << 7) - 1)
	actual, _ = bits.GetSlice(63, 8)
	assert.Equal(t, expected, actual, "slice at range(63, 71) should be %d", expected)

	err := bits.SetSlice(0, 0, 5)
	assert.Nil(t, err, "if the sliceSize is 0 then return nil")

	err = bits.SetSlice(0, BLOCK_SIZE+5, 10)
	assert.Error(t, err, "if the sliceSize is larger than BLOCK_SIZE then return error")
}

func TestResize(t *testing.T) {
	bits := New()
	bits.resize(4)
	expected := uint64(4)
	actual := bits.Size()
	assert.Equal(t, expected, actual, "the size of bits should be %d", expected)

	bits.resize(456)
	expectedSize := uint64(456)
	expectedLen := expectedSize/BLOCK_SIZE + 1
	actualSize := bits.Size()
	actualLen := uint64(len(bits.blocks))
	assert.Equal(t, expectedSize, actualSize, "the size of bits should be %d", expectedSize)
	assert.Equal(t, expectedLen, actualLen, "the length of blocks should be %d", expectedLen)
}
