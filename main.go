package main

import (
	"github.com/mikesupertrampster/simple-json-server/pkg/server"
	"github.com/vrischmann/envconfig"
	"log"
	"net/http"
)

type cfg struct {
	Port string `envconfig:"default=8000"`

	Data struct {
		Path      string `envconfig:"default=./data"`
		Extension string `envconfig:"default=.json"`
	}
}

func main() {
	config := new(cfg)
	if err := envconfig.Init(config); err != nil {
		log.Fatal(err)
	}

	s := server.Server{
		Path:      config.Data.Path,
		Extension: config.Data.Extension,
	}

	http.HandleFunc("/", s.Handler)
	server := ":" + config.Port
	if err := http.ListenAndServe(server, nil); err != nil {
		log.Fatal(err)
	}
}

