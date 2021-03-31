package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestResponse(t *testing.T) {
	s := Server{
		Path:      "./data",
		Extension: ".json",
	}

	tt := []struct {
		test     string
		target   string
	}{
		{
			"happy path",
			"foo",
		},
		{
			"invalid target",
			"bar",
		},
		{
			"empty body",
			"SKIP",
		},
	}

	for _, tc := range tt {
		t.Run(tc.test, func(t *testing.T) {
			req, err := json.Marshal(map[string]string{
				"target": tc.target,
			})
			if err != nil {
				t.Fatal(err)
			}

			var body io.Reader
			if body = bytes.NewBuffer(req); tc.target == "SKIP" {
				body = nil
			}

			r, err := http.NewRequest("POST", "/", body)
			if err != nil {
				t.Fatal(err)
			}

			w := httptest.NewRecorder()
			s.Handler(w, r)

			if status := w.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
		})
	}
}
