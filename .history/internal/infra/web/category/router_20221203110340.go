package infra

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

type (
	categoriesResource struct {
		// repo domain.CategoryRepository
	}
)

func (rs categoriesResource) Routes() chi.Router {
	router = categoriesResource{
		// repo: repo
	}

	r := chi.NewRouter()
	r.Post(Create)
}

func CategoryCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := context.WithValue(req.Context(), "id", chi.URLParam(req, "id"))
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}
