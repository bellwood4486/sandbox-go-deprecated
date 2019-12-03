package main

import (
	"fmt"
	"math/rand"

	"github.com/pkg/errors"
)

func random(src [][]string) func() ([]string, error) {
	return func() ([]string, error) {
		return src[rand.Intn(len(src))], nil
	}
}

func roundRobin(src [][]string) func() ([]string, error) {
	current := -1
	return func() ([]string, error) {
		if current == len(src)-1 {
			current = -1 // reset
		}
		current++
		return src[current], nil
	}
}

func once(src [][]string) func() ([]string, error) {
	current := -1
	return func() ([]string, error) {
		if current == len(src)-1 {
			return nil, errors.New("used up")
		}
		current++
		return src[current], nil
	}
}

func mapper(keys []string, selector func() ([]string, error)) func() (map[string]string, error) {
	return func() (map[string]string, error) {
		values, err := selector()
		if err != nil {
			return nil, err
		}
		m := make(map[string]string)
		for i, v := range keys {
			m[v] = values[i]
		}
		return m, nil
	}
}

func main() {
	single := [][]string{{"a"}, {"b"}}
	group := [][]string{{"1", "2"}, {"3", "4"}}

	mpr1 := mapper([]string{"k1"}, random(single))
	mpr2 := mapper([]string{"i1", "i2"}, once(group))

	for i := 0; i < 10; i++ {
		fmt.Println(mpr1())
	}
	for i := 0; i < 10; i++ {
		fmt.Println(mpr2())
	}
}
