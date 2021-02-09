package slice_allocate

import (
	"testing"
)

const (
	_       = iota
	KiB int = 1 << (10 * iota)
	MiB
)

const capacity = 10000

type BigStruct struct {
	Data [100 * KiB]byte
}

func BenchmarkAllocateOnce(b *testing.B) {
	_ = make([]BigStruct, 0, capacity)
}

func BenchmarkAllocateEverytime(b *testing.B) {
	s := make([]*BigStruct, 0, capacity)
	for i := 0; i < capacity; i++ {
		s = append(s, &BigStruct{})
	}
}

