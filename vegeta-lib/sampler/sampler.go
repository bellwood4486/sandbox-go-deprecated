package sampler

import "net/http"

type Sampler interface {
	Templates() ([]RequestTemplate, error)
	HandleResponse(method, url string, code uint16, body []byte) error
}

type RequestTemplate struct {
	Weight uint
	Method string
	URL    string
	Header http.Header
	Body   string
}
