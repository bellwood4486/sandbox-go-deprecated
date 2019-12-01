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
	templateCache map[string]*template.Template
	once          sync.Once
}

func (r *RequestTemplate) InjectToPath(data map[string]string) string {
	return r.inject("path", r.Path, data)
}

func (r *RequestTemplate) InjectToBody(data map[string]string) string {
	return r.inject("body", r.Body, data)
}

func (r *RequestTemplate) inject(cacheKey, text string, data map[string]string) string {
	templ := r.template(cacheKey, text)
	buf := new(bytes.Buffer)
	_ = templ.Execute(buf, data)
	return buf.String()
}

func (r *RequestTemplate) template(cacheKey, text string) *template.Template {
	r.once.Do(func() { r.templateCache = make(map[string]*template.Template) })

	templ, ok := r.templateCache[cacheKey]
	if !ok {
		templ = template.Must(template.New(cacheKey).Parse(text))
		r.templateCache[cacheKey] = templ
	}
	return templ
}
