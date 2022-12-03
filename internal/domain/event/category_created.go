package domain

import "github.com/devfullcycle/catalogo-fc/pkg/entity"

type CategoryCreated struct {
	ID          entity.ID
	Name        string
	Description string
	IsActive    bool
}
