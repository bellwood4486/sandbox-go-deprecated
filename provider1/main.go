package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"
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

type DataProviders []DataProvider

func (d DataProviders) TemplateParamMap() map[string]string {
	m := make(map[string]string)
	for _, p := range t1.Providers {
		for k, v := range p.Provide() {
			m[k] = v
		}
	}
	return m
}

type DataProvider interface {
	Provide() map[string]string
}

type SingleData struct {
	key        string
	dataSource []string
}

func (p SingleData) Provide() map[string]string {
	value := p.dataSource[rand.Intn(len(p.dataSource))]
	return map[string]string{p.key: value}
}

type GroupData struct {
	keys       []string
	dataSource [][]string
}

func (p GroupData) Provide() map[string]string {
	values := p.dataSource[rand.Intn(len(p.dataSource))]
	if len(p.keys) != len(values) {
		panic("length difference")
	}
	m := make(map[string]string)
	for i, key := range p.keys {
		m[key] = values[i]
	}
	return m
}

type Timestamp struct {
	key string
}

func (p Timestamp) Provide() map[string]string {
	value := strconv.FormatInt(time.Now().UnixNano(), 10)
	return map[string]string{p.key: value}
}

type RequestTemplate struct {
	Path          string
	Body          string
	Providers     DataProviders
	templateCache map[string]*template.Template
	once          sync.Once
}

func (r RequestTemplate) InjectToPath(data map[string]string) string {
	return r.inject("path", r.Path, data)
}

func (r RequestTemplate) InjectToBody(data map[string]string) string {
	return r.inject("body", r.Body, data)
}

func (r RequestTemplate) inject(cacheKey, text string, data map[string]string) string {
	templ := r.getTemplate(cacheKey, text)
	buf := new(bytes.Buffer)
	_ = templ.Execute(buf, data)
	return buf.String()
}

func (r *RequestTemplate) getTemplate(cacheKey, text string) *template.Template {
	r.once.Do(func() { r.templateCache = make(map[string]*template.Template) })

	templ, ok := r.templateCache[cacheKey]
	if !ok {
		templ = template.Must(template.New(cacheKey).Parse(text))
		r.templateCache[cacheKey] = templ
	}
	return templ
}

var t1 = RequestTemplate{
	Path: "/api/{{ .Name }}",
	Body: `id-{{ .Timestamp }} {{ .Name }} {{ .AAA }} {{ .BBB }} `,
	Providers: DataProviders{
		Timestamp{"Timestamp"},
		SingleData{"Name", []string{"a", "b", "c"}},
		GroupData{[]string{"AAA", "BBB"}, [][]string{{"1", "2"}, {"3", "4"}}},
	},
}

func main() {
	rand.Seed(time.Now().Unix())

	m := t1.Providers.TemplateParamMap()
	fmt.Println(m)

	p := t1.InjectToPath(m)
	fmt.Println(p)

	b := t1.InjectToBody(m)
	fmt.Println(b)

	//dataSource := []string{"アルファ", "ベータ", "ガンマ"}
	//f := provider1(dataSource)
	//fmt.Println(f())
	//
	//data2 := [][]string{
	//	{"a", "1", "あ"},
	//	{"b", "2", "い"},
	//}
	//f2 := provider2(data2)
	//p2 := strings.Split("k1,k2,k3", ",")
	//v2 := f2()
	//if len(p2) != len(v2) {
	//	panic("length difference")
	//}
	//m := make(map[string]string)
	//for i, v := range p2 {
	//	m[v] = v2[i]
	//}
	//fmt.Println(m)
	//
	//merge()
}
