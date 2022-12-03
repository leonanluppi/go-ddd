package domain

import domain "github.com/devfullcycle/catalogo-fc/internal/domain/shared"

type CategoryRepository interface {
	domain.Repository[*Category]
}
