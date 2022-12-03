package infra

import (
	"context"
	"net/http"

	domain "github.com/devfullcycle/catalogo-fc/internal/domain/category"
	"github.com/go-chi/chi"
)

type (
	CategoriesResource struct {
		repo domain.CategoryRepository
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
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := context.WithValue(req.Context(), "id", chi.URLParam(req, "id"))
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}
