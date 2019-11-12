package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type args struct {
	w *httptest.ResponseRecorder
	r *http.Request
}

func newArgs(method, target string, body io.Reader) args {
	return args{
		httptest.NewRecorder(),
		httptest.NewRequest(method, target, body),
	}
}

func Test_getOneEvent(t *testing.T) {
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			"found",
			newArgs(http.MethodGet, "/events/1", nil),
			http.StatusOK,
		},
		{
			"not found",
			newArgs(http.MethodGet, "/events/100", nil),
			http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newRouter().ServeHTTP(tt.args.w, tt.args.r)
			if tt.args.w.Code != tt.wantCode {
				t.Errorf("getOneEvent() = %v, want %v", tt.args.w.Code, tt.wantCode)
			}
		})
	}
}
