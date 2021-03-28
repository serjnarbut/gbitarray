/*
Copyright 2021 Sergey narbut

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"),
to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

/*
Package bitarray represents a simple bit array.
*/
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
