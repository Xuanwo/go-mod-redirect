package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Xuanwo/go-mod-redirect/config"
	"github.com/Xuanwo/go-mod-redirect/handler"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Config file must be provided.")
	}

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	service, err := config.Parse(content)
	if err != nil {
		log.Fatal(err)
	}

	h, err := handler.New(service)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", h)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
