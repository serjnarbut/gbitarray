package gbitarray

const bitsSize = 32

type BitArray struct {
	Size int
	bits []int32
}

func NewBitArray(size int) *BitArray {
	innerSize := ceil(size, bitsSize)
	bits := make([]int32, innerSize)
	return &BitArray{
		Size: size,
		bits: bits,
	}
}

func (b *BitArray) SetBit(index int) {
	innerIndex := index / bitsSize
	b.bits[innerIndex] |= 1 << (index % bitsSize)
}

func (b *BitArray) GetBit(index int) int32 {
	innerIndex := index / bitsSize
	return (b.bits[innerIndex] >> (index % bitsSize)) & 1
}

func (b *BitArray) HasBit(index int) bool {
	return b.GetBit(index) != 0
}

func ceil(x int, y int) int {
	return (x + y - 1) / y
}
