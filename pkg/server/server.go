package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Server struct {
	Path      string
	Extension string
}

type Body struct {
	Target string `json:"target"`
}

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		return
	}

	var body Body
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&body)
	if err != nil {
		body.Target = ""
	}

	s.Handle(w, r, body.Target)
}

func (s *Server) Handle(w http.ResponseWriter, _ *http.Request, target string) {
	write := []byte("no target")

	if target != "" {
		write = []byte(s.read(target))
	}

	_, err := w.Write(write)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) read(d string) string {
	file := fmt.Sprintf("%s/%s%s", s.Path, d, s.Extension)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("no such file: %s", file)
		return "no such file"
	}

	return string(data)
}

