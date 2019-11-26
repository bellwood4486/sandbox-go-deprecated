package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/pkg/errors"
)

type Shuffler struct {
	data            map[int]string
	totalPercentage int
}

func NewShuffler() *Shuffler {
	return &Shuffler{
		data:            make(map[int]string, 100),
		totalPercentage: 0,
	}
}

func (s *Shuffler) Add(percentage int, value string) {
	if percentage < 0 || 100 < percentage {
		panic(fmt.Sprintf("invalid percentage: %v must be 0 <= percentage <=100", percentage))
	}

	for i := s.totalPercentage; i < s.totalPercentage+percentage; i++ {
		s.data[i] = value
	}
	s.totalPercentage += percentage
}

func (s *Shuffler) Get() (string, error) {
	if s.totalPercentage != 100 {
		return "", errors.Errorf("invalid total percentage: %v must be 100", s.totalPercentage)
	}

	return s.data[rand.Intn(100)], nil
}

func main() {
	s := NewShuffler()
	s.Add(40, "a")
	s.Add(30, "b")
	s.Add(20, "c")
	s.Add(5, "d")
	s.Add(5, "e")

	for i := 0; i < 100; i++ {
		v, err := s.Get()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(v)
	}
}
