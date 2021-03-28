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

func BenchmarkSetBits(b *testing.B) {
	bits := NewBitArray(b.N)

	for i := 0; i < b.N; i++ {
		bits.SetBit(i)
	}
}

func BenchmarkGetBits(b * testing.B) {
	bits := NewBitArray(b.N)

	for i := 0; i < b.N; i++ {
		if i % 2 == 0  {
			bits.SetBit(i)
		}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bits.HasBit(i)
	}
}