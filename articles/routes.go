package articles

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/articles", createArticleHandler)
	// r.Get("/articles", listArticlesHandler)
	// r.Delete("/articles/{id}", deleteArticlesHandler)
	// r.Put("/articles/{id}", updateArticlesHandler)

	return r
}
