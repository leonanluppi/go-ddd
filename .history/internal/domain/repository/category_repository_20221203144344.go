package domain

type CategoryRepository interface {
	domain.Repository[*Category]
}
