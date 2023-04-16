package articles

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/articles", createArticleHandler)
	r.Get("/articles", listArticlesHandler)
	r.Put("/articles", updateArticlesHandler)
	// r.Delete("/articles/{id}", deleteArticlesHandler)
	return r
}
