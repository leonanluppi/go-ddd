package domain

// import "github.com/devfullcycle/catalogo-fc/internal/domain/category"
// type Repository[T comparable] interface {
// 	Insert(*T) (*T, error)
// 	Update(*T) (*T, error)
// 	Delete(entity.ID) error
// 	Paginate(SearchParams) (*SearchResult[*T], error)
// 	FindByID(entity.ID) (*T, error)
// }

type Repository interface {
	Insert(name string) (string, error)
}
