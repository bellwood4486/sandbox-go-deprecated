package main

import (
	"fmt"
	"math/rand"
	"time"

	ds "github.com/bellwood4486/sandbox-go/provider2/datasource"
	"github.com/bellwood4486/sandbox-go/provider2/template"
)

var t1 = template.RequestTemplate{
	Path: "/api/{{ .Name }}",
	Body: `id-{{ .Timestamp }} {{ .Name }} {{ .AAA }} {{ .BBB }} `,
	Providers: ds.DataProviders{
		ds.Timestamp{Key: "Timestamp"},
		ds.SingleData{Key: "Name", Source: &ds.Names},
		ds.GroupData{Keys: []string{"AAA", "BBB"}, Source: &ds.Groups},
	},
}

func main() {
	rand.Seed(time.Now().Unix())

	m := t1.Providers.ParameterMap()
	fmt.Println(m)

	p := t1.InjectToPath(m)
	fmt.Println(p)

	b := t1.InjectToBody(m)
	fmt.Println(b)
}
