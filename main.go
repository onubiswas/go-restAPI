package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"example.com/rest/articles"
)

func main() {
	fmt.Println("Lets build a RESTful API with Go")

	r := chi.NewRouter()
	fmt.Println("starting the server now")
	r.Mount("/v1/", articles.Routes())
	if err := http.ListenAndServe(":3030", r); err != nil {
		fmt.Println("server start failed", err)
	}

}
