package main

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestValidateCreditCard(t *testing.T) {
	app := newTestApp()
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name     string
		urlPath  string
		wantBody string
		wantCode int
	}{
		{
			name:     "Correct Validation",
			urlPath:  "/validate/4567890123456783",
			wantBody: "true",
			wantCode: http.StatusOK,
		},
		{
			name:     "Incorrect Validation",
			urlPath:  "/validate/4567890123456784",
			wantBody: "false",
			wantCode: http.StatusOK,
		},
		{
			name:     "Empty Number",
			urlPath:  "/validate/",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Not Numeric Number",
			urlPath:  "/validate/12345678a7654321",
			wantCode: http.StatusUnprocessableEntity,
		},
		{
			name:     "Invalid Length Number",
			urlPath:  "/validate/1234567654321",
			wantCode: http.StatusUnprocessableEntity,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)
			fmt.Println(code, body)
			if code != tt.wantCode {
				t.Error("den exei idio cpde")
			}
			if !strings.Contains(body, tt.wantBody) {
				t.Error("den exei idio body")
			}
		})
	}
}
