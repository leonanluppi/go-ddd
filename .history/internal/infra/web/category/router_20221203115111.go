package infra

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type (
	CategoriesResource struct {
		// repo domain.CategoryRepository
	}
)

func (rs CategoriesResource) NewRoutes() chi.Router {
	rs = CategoriesResource{
		// repo: repo
	}

	router := chi.NewRouter()
	router.Post("/", rs.Create)

	return router
}

func CategoryCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := context.WithValue(req.Context(), "id", chi.URLParam(req, "id"))
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}
