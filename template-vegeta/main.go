package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"os"
	"text/template"
)

type RequestTemplate struct {
	Method string
	URL    string
	Header map[string][]string
	Body   string
}

type VegetaRequest struct {
	Method     string              `json:"method"`
	URL        string              `json:"url"`
	Header     map[string][]string `json:"header"`
	BodyBase64 string              `json:"body"`
}

var orderTempl = RequestTemplate{
	Method: "POST",
	URL:    "/api/order/{{.ID}}",
	Header: map[string][]string{"Content-type": {"application/json"}},
	Body: `{
  "name": "{{.Name}}"
}`,
}

type Param struct {
	ID   int
	Name string
}

func main() {
	p := Param{1, "mike"}
	templ := orderTempl

	urlTempl := template.Must(template.New("url").Parse(templ.URL))
	urlBuf := new(bytes.Buffer)
	_ = urlTempl.Execute(urlBuf, p)

	bodyTempl := template.Must(template.New("body").Parse(templ.Body))
	bodyBuf := new(bytes.Buffer)
	_ = bodyTempl.Execute(bodyBuf, p)

	req := VegetaRequest{
		Method:     templ.Method,
		URL:        urlBuf.String(),
		Header:     templ.Header,
		BodyBase64: base64.StdEncoding.EncodeToString(bodyBuf.Bytes()),
	}

	_ = json.NewEncoder(os.Stdout).Encode(req)
}
