package datasource

import (
	"math/rand"
	"strconv"
	"time"
)

// DataProviders contains ...
type DataProviders []DataProvider

func (d DataProviders) ParameterMap() map[string]string {
	m := make(map[string]string)
	for _, p := range d {
		for k, v := range p.Provide() {
			m[k] = v
		}
	}
	return m
}

// DataProvider represents ...
type DataProvider interface {
	Provide() map[string]string
}

// SingleData provides ...
type SingleData struct {
	Key    string
	Source *[]string
}

func (p SingleData) Provide() map[string]string {
	value := (*p.Source)[rand.Intn(len(*p.Source))]
	return map[string]string{p.Key: value}
}

// GroupData provides ...
type GroupData struct {
	Keys   []string
	Source *[][]string
}

func (p GroupData) Provide() map[string]string {
	values := (*p.Source)[rand.Intn(len(*p.Source))]
	if len(p.Keys) != len(values) {
		panic("length difference")
	}
	m := make(map[string]string)
	for i, key := range p.Keys {
		m[key] = values[i]
	}
	return m
}

// Timestamp provides ...
type Timestamp struct {
	Key string
}

func (p Timestamp) Provide() map[string]string {
	value := strconv.FormatInt(time.Now().UnixNano(), 10)
	return map[string]string{p.Key: value}
}
