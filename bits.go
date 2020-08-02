package fid

import "fmt"

const BLOCK_SIZE = 64

type Bits struct {
	blocks []uint64
	size   uint64 // size represents the length of bits stored
}

// New returns Bits with length 1.
func New() *Bits {
	return &Bits{
		blocks: make([]uint64, 1),
		size:   0,
	}
}

// Size returns the length of bits.
func (b *Bits) Size() uint64 {
	return b.size
}

// GetBit returns the bit at position i
func (b *Bits) GetBit(i uint64) (bit bool, err error) {
	if i > b.size {
		return false, fmt.Errorf("i should be smaller than the size of bits: i %d, bsize %d", i, b.size)
	}

	blockIndex, bitIndex := getPositions(i)
	return ((b.blocks[blockIndex] >> (bitIndex - 1)) & 1) == 1, nil
}

// SetBit sets the bit b at position i.
// If i is larger than current size of bits, the size extends automatically.
func (b *Bits) SetBit(i uint64, bit bool) {
	blockIndex, bitIndex := getPositions(i)
	if i > b.size {
		b.resize(i)
	}

	mask := uint64(1) << (bitIndex - 1)
	if bit {
		b.blocks[blockIndex] |= mask
	} else {
		b.blocks[blockIndex] &= mask
	}
}

// GetSlice returns slice of sliceSize at position i.
// sliceSize must be smaller than BLOCK_SIZE and i + sliceSize must not be exeeded the size of bits.
func (b *Bits) GetSlice(i uint64, sliceSize uint8) (slice uint64, err error) {
	if sliceSize == 0 {
		return 0, nil
	}
	if sliceSize > BLOCK_SIZE {
		return 0, fmt.Errorf("sliceSize should be smaller than BLOCK_SIZE: sliceSize %d, BLOCK_SIZE %d", sliceSize, BLOCK_SIZE)
	}
	if i+uint64(sliceSize) > b.size {
		return 0, fmt.Errorf("sliceSize should be smaller than the size of bits: sliceSize %d, b.size %d", sliceSize, b.size)
	}

	blockIndex, bitIndex := getPositions(i)
	slice = (b.blocks[blockIndex] >> bitIndex)
	if bitIndex+uint64(sliceSize) > BLOCK_SIZE {
		slice |= (b.blocks[blockIndex+1] << (BLOCK_SIZE - bitIndex))
	}
	return slice & ((1 << sliceSize) - 1), nil
}

// SetSlice sets slice of length sliceSize at position i.
// If i + sliceSize is larger than current size of bits, the size extends automatically.
// sliceSize must be smaller than BLOCK_SIZE.
func (b *Bits) SetSlice(i uint64, sliceSize uint8, slice uint64) (err error) {
	if sliceSize == 0 {
		return nil
	}
	if sliceSize > BLOCK_SIZE {
		return fmt.Errorf("sliceSize should be smaller than BLOCK_SIZE: sliceSize %d, BLOCK_SIZE %d", sliceSize, BLOCK_SIZE)
	}
	if i+uint64(sliceSize) > b.size {
		b.resize(i + uint64(sliceSize))
	}

	blockIndex, bitIndex := getPositions(i)
	b.blocks[blockIndex] |= slice << bitIndex
	if bitIndex+uint64(sliceSize) > BLOCK_SIZE {
		b.blocks[blockIndex+1] |= (slice >> (BLOCK_SIZE - bitIndex))
	}
	return nil
}

// getPositions returns blockIndex and bitIndex
func getPositions(i uint64) (uint64, uint64) {
	return i / BLOCK_SIZE, i % BLOCK_SIZE
}

// resize extends the size of bits to newSize
// TODO: optimize memory allocation
func (b *Bits) resize(newSize uint64) {
	if b.size >= newSize {
		return
	}

	blockIndex, _ := getPositions(newSize)
	if blockIndex > uint64(len(b.blocks)-1) { // expand blocks
		diff := blockIndex - uint64(len(b.blocks)) + 1
		for i := uint64(0); i < diff; i++ {
			b.blocks = append(b.blocks, uint64(0))
		}
	}
	b.size = newSize
}

// for debug
func printBit(x uint64) {
	for i := 0; i < 64; i++ {
		fmt.Printf("%d", i%10)
	}
	fmt.Printf("\n")
	for i := 0; i < 64; i++ {
		bit := (x & (1 << (i - 1))) > 0
		if bit {
			fmt.Printf("1")
		} else {
			fmt.Printf("0")
		}
	}
	fmt.Printf("\n")
}
