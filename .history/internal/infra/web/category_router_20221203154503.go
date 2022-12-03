package infra

import (
	"context"
	"net/http"

	application "github.com/devfullcycle/catalogo-fc/internal/application/usecase"
	domain "github.com/devfullcycle/catalogo-fc/internal/domain/repository"
	"github.com/go-chi/chi/v5"
)

type (
	CategoriesResource struct {
		repo                  domain.CategoryRepository
		createCategoryUseCase application.UseCase[
			application.CategoryCreateInput,
			application.CategoryOutput,
		]
	}
)

func (rs CategoriesResource) NewRoutes(repo domain.CategoryRepository) chi.Router {
	rs = CategoriesResource{
		repo:                  repo,
		createCategoryUseCase: application.NewCreateUseCase(repo),
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
