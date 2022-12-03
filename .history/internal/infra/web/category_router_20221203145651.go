package infra

import (
	"context"
	"net/http"

	domain "github.com/devfullcycle/catalogo-fc/internal/domain/repository"
	"github.com/go-chi/chi/v5"
)

type (
	CategoriesResource struct {
		repo domain.CategoryRepository
	}
)

func (rs CategoriesResource) NewRoutes(repo domain.CategoryRepository) chi.Router {
	rs = CategoriesResource{
		repo: repo
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
