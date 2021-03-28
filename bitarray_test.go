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
package gbitarray

import "testing"

func TestEmptyBitArray(t *testing.T) {
	bits := NewBitArray(10)

	for i := 0; i < bits.Size; i++ {
		if bits.HasBit(i) {
			t.Error("Bit {} must be empty", i)
		}
	}
}

func TestSetGetBits(t *testing.T) {
	bits := NewBitArray(10)

	bits.SetBit(2)

	if !bits.HasBit(2) {
		t.Error("Bit 2 must be set")
	}
}

func TestDuplicateSetBit(t *testing.T) {
	bits := NewBitArray(10)

	bits.SetBit(4)

	if !bits.HasBit(4) {
		t.Error("Bit 4 must be set")
	}

	bits.SetBit(4)

	if !bits.HasBit(4) {
		t.Error("Bit 4 must be set")
	}
}

func BenchmarkSetBits(b *testing.B) {
	bits := NewBitArray(b.N)

	for i := 0; i < b.N; i++ {
		bits.SetBit(i)
	}
}

func BenchmarkGetBits(b *testing.B) {
	bits := NewBitArray(b.N)

	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			bits.SetBit(i)
		}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bits.HasBit(i)
	}
}
