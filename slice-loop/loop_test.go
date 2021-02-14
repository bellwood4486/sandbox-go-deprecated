package slice_loop

import (
	"testing"
)

type BigStruct100 struct {
	Data [100]byte
}

type BigStruct1000 struct {
	Data [1000]byte
}

func BenchmarkLoopWithIndex_100(b *testing.B) {
	s := make([]BigStruct100, b.N)
	b.ResetTimer()
	for i := range s {
		s[i].Data[0] = 'a'
	}
}

func BenchmarkLoopWithValue_100(b *testing.B) {
	s := make([]BigStruct100, b.N)
	b.ResetTimer()
	for _, v := range s {
		v.Data[0] = 'a'
	}
}

func BenchmarkLoopWithIndex_1000(b *testing.B) {
	s := make([]BigStruct1000, b.N)
	b.ResetTimer()
	for i := range s {
		s[i].Data[0] = 'a'
	}
}

func BenchmarkLoopWithValue_1000(b *testing.B) {
	s := make([]BigStruct1000, b.N)
	b.ResetTimer()
	for _, v := range s {
		v.Data[0] = 'a'
	}
}
