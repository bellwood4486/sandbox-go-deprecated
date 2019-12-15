package sampler

type Service1Sampler struct {
	data []map[string]string
}

func (s *Service1Sampler) Templates() ([]RequestTemplate, error) {
	if err := initDataSource(s); err != nil {
		return nil, err
	}

	return []RequestTemplate{
		{Weight: 1, Method: "GET", URL: "http://localhost:8080/a"},
		{Weight: 1, Method: "GET", URL: "http://localhost:8080/b"},
		{Weight: 1, Method: "GET", URL: "http://localhost:8080/c"},
		{Weight: 1, Method: "GET", URL: "http://localhost:8080/d"},
		{Weight: 1, Method: "GET", URL: "http://localhost:8080/e"},
	}, nil
}

func (s *Service1Sampler) HandleResponse(method, url string, code uint16, body []byte) error {
	// parse body
	// do something
	switch url {
	case "path is /a":
		s.data = append(s.data, map[string]string{"f": "5", "g": "6"})
	case "path is /b":
		s.data = append(s.data, map[string]string{"h": "7", "i": "8"})
	}

	return nil
}

func initDataSource(s *Service1Sampler) error {
	// fetch data from DB or local file...
	s.data = []map[string]string{
		{"a": "1", "b": "2"},
		{"c": "3", "d": "4"},
	}

	return nil
}
