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

	r.Mount("/v1/", articles.Routes())
	http.ListenAndServe(":3000", r)

}
