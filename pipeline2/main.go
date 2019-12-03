package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/pkg/errors"
)

type SelectorFunc func() (map[string]string, error)

type MapperFunc func() (map[string]string, error)

func random(src []map[string]string) SelectorFunc {
	return func() (map[string]string, error) {
		return src[rand.Intn(len(src))], nil
	}
}

func roundRobin(src []map[string]string) SelectorFunc {
	current := -1
	return func() (map[string]string, error) {
		if current == len(src)-1 {
			current = -1 // reset
		}
		current++
		return src[current], nil
	}
}

func roundRobinOnce(src []map[string]string) SelectorFunc {
	current := -1
	return func() (map[string]string, error) {
		if current == len(src)-1 {
			return nil, errors.New("used up")
		}
		current++
		return src[current], nil
	}
}

func timestamp(key string) MapperFunc {
	return func() (map[string]string, error) {
		return map[string]string{"key": "11111111"}, nil
	}
}

func mapper(keys []string, selector SelectorFunc) MapperFunc {
	return func() (map[string]string, error) {
		values, err := selector()
		if err != nil {
			return nil, err
		}
		m := make(map[string]string)
		for _, v := range keys {
			if _, ok := values[v]; !ok {
				return nil, errors.Errorf("unknown key %q in %v", v, values)
			}
			m[v] = values[v]
		}
		return m, nil
	}
}

func main() {
	single := []map[string]string{{"k": "va"}, {"k": "vb"}}
	group := []map[string]string{{"k1": "1", "k2": "2"}, {"k1": "3", "k2": "4"}}

	mpr1 := mapper([]string{"k"}, random(single))
	//mpr1 := mapper([]string{"a"}, random(single))
	mpr2 := mapper([]string{"k1", "k2"}, roundRobin(group))
	//mpr2 := mapper([]string{"a1"}, roundRobin(group))

	var mprs []MapperFunc
	mprs = append(mprs, mpr1)
	mprs = append(mprs, mpr2)

	mm := make(map[string]string)
	for _, mpr := range mprs {
		m, err := mpr()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		for k, v := range m {
			mm[k] = v
		}
	}
	fmt.Println(mm)

	//for i := 0; i < 3; i++ {
	//	fmt.Println(mpr1())
	//}
	//for i := 0; i < 3; i++ {
	//	fmt.Println(mpr2())
	//}
}
