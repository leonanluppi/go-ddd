package shared_domain

type SearchParams struct {
	Filter  string
	Order   string
	Page    int
	PerPage int
}

func (sp SearchParams) Offset() int {
	return (sp.Page - 1) * sp.PerPage
}

type SearchResult[T any] struct {
	Items       []T
	Total       int64
	CurrentPage int
	LastPage    int
	PerPage     int
}
