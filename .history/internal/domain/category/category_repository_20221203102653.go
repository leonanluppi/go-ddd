package categories

// import (
// 	shared_domain "fc/go-fc3/internal/domain/shared"
// 	base "fc/go-fc3/pkg/entity"
// )

type Repository interface {
	Insert(*Category) (*Category, error)
	Update(*Category) (*Category, error)
	// Delete(base.ID) error
	// Paginate(shared_domain.SearchParams) (*shared_domain.SearchResult[*Category], error)
	// FindByID(base.ID) (*Category, error)
}
