package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func provider(data []string) func() string {
	return func() string {
		return data[rand.Intn(len(data))]
	}
}

func provider2(data [][]string) func() []string {
	return func() []string {
		return data[rand.Intn(len(data))]
	}
}

func merge() {
	data := []string{"アルファ", "ベータ", "ガンマ"}
	data2 := [][]string{
		{"a", "1", "あ"},
		{"b", "2", "い"},
	}
	params := map[string]interface{}{
		"k1":       provider(data),
		"k2,K3,k4": provider2(data2),
	}

	m := make(map[string]string)
	for key, param := range params {
		ks := strings.Split(key, ",")
		switch ln := len(ks); {
		case ln == 1:
			m[ks[0]] = param.(func() string)()
		case ln >= 2:
			vs := param.(func() []string)()
			if len(vs) != ln {
				panic("length difference")
			}
			for i, v := range ks {
				m[v] = vs[i]
			}
		default:
			panic(fmt.Sprintf("unknown length: %d", ln))
		}
	}

	fmt.Println(m)
}

func main() {
	dataSource := []string{"アルファ", "ベータ", "ガンマ"}
	f := provider(dataSource)
	fmt.Println(f())

	data2 := [][]string{
		{"a", "1", "あ"},
		{"b", "2", "い"},
	}
	f2 := provider2(data2)
	p2 := strings.Split("k1,k2,k3", ",")
	v2 := f2()
	if len(p2) != len(v2) {
		panic("length difference")
	}
	m := make(map[string]string)
	for i, v := range p2 {
		m[v] = v2[i]
	}
	fmt.Println(m)

	merge()
}
