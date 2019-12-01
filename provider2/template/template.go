package template

import (
	"bytes"
	"html/template"
	"sync"

	"github.com/bellwood4486/sandbox-go/provider2/datasource"
)

// RequestTemplate is ...
type RequestTemplate struct {
	Path          string
	Body          string
	Providers     datasource.DataProviders
	initCacheOnce sync.Once
	templateCache map[string]*template.Template
}

func (r *RequestTemplate) InjectToPath(data map[string]string) string {
	return r.inject(r.Path, data)
}

func (r *RequestTemplate) InjectToBody(data map[string]string) string {
	return r.inject(r.Body, data)
}

func (r *RequestTemplate) inject(text string, data map[string]string) string {
	templ := r.template(text)
	buf := new(bytes.Buffer)
	_ = templ.Execute(buf, data)
	return buf.String()
}

func (r *RequestTemplate) template(text string) *template.Template {
	r.initCacheOnce.Do(func() { r.templateCache = make(map[string]*template.Template) })

	templ, ok := r.templateCache[text]
	if !ok {
		templ = template.Must(template.New("template for " + text).Parse(text))
		r.templateCache[text] = templ
	}
	return templ
}
