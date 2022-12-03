package domain

import domain "github.com/devfullcycle/catalogo-fc/internal/domain/entity"

type CategoryRepository interface {
	Repository[*domain.Category]
}
