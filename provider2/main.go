package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bellwood4486/sandbox-go/provider2/template"

	"github.com/bellwood4486/sandbox-go/provider2/datasource"
)

var t1 = template.RequestTemplate{
	Path: "/api/{{ .Name }}",
	Body: `id-{{ .Timestamp }} {{ .Name }} {{ .AAA }} {{ .BBB }} `,
	Providers: datasource.DataProviders{
		datasource.Timestamp{"Timestamp"},
		datasource.SingleData{"Name", []string{"a", "b", "c"}},
		datasource.GroupData{[]string{"AAA", "BBB"}, [][]string{{"1", "2"}, {"3", "4"}}},
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
}
