package serialize

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deepCopyJSON(t *testing.T) {
	type args struct {
		src *SampleSmall
		dst *SampleSmall
	}
	s := "a"
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				src: &SampleSmall{[]*string{&s}},
				dst: &SampleSmall{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deepCopyJSON(tt.args.src, tt.args.dst)
			assert.Equal(t, tt.args.dst, tt.args.src)
			assert.NotEqual(t, memAddr(tt.args.dst.Values[0]), memAddr(tt.args.src.Values[0]))
		})
	}
}

func Test_deepCopyGob(t *testing.T) {
	s := "a"
	type args struct {
		src *SampleSmall
		dst *SampleSmall
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				src: &SampleSmall{[]*string{&s}},
				dst: &SampleSmall{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deepCopyGob(tt.args.src, tt.args.dst)
			assert.Equal(t, tt.args.dst, tt.args.src)
			assert.NotEqual(t, memAddr(tt.args.dst.Values[0]), memAddr(tt.args.src.Values[0]))
		})
	}
}

func Test_deepCopyGen(t *testing.T) {
	s := "a"
	type args struct {
		src *SampleSmall
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				src: &SampleSmall{[]*string{&s}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.src.DeepCopy()
			assert.Equal(t, got, tt.args.src)
		})
	}
}

func Benchmark_deepCopyJSON(b *testing.B) {
	var dst SampleLarge
	for i := 0; i < b.N; i++ {
		deepCopyJSON(largeObj, &dst)
	}
}

func Benchmark_deepCopyGob(b *testing.B) {
	var dst SampleLarge
	for i := 0; i < b.N; i++ {
		deepCopyGob(largeObj, &dst)
	}
}

func Benchmark_deepCopyGen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = largeObj.DeepCopy()
	}
}

func memAddr(v interface{}) string {
	return fmt.Sprintf("%p", v)
}
