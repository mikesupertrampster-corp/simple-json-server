package main

import (
	"encoding/json"
	"fmt"
	"github.com/vrischmann/envconfig"
	"io/ioutil"
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

type Body struct {
	Target string `json:"target"`
}

func main() {
	config := new(cfg)
	if err := envconfig.Init(config); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", router)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.Port), nil); err != nil {
		log.Fatal(err)
	}
}

func router(w http.ResponseWriter, r *http.Request) {
	var body Body
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&body)
	if err != nil {
		body.Target = ""
	}

	handle(w, r, body.Target)
}

func handle(w http.ResponseWriter, _ *http.Request, target string) {
	write := []byte("no target")

	if target != "" {
		write = []byte(read(target))
	}

	_, err := w.Write(write)
	if err != nil {
		log.Fatal(err)
	}
}

func read(d string) string {
	config := new(cfg)
	if err := envconfig.Init(config); err != nil {
		log.Fatal(err)
	}

	file := fmt.Sprintf("%s/%s%s", config.Data.Path, d, config.Data.Extension)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("no such file: %s", file)
		return "no such file"
	}

	return string(data)
}
