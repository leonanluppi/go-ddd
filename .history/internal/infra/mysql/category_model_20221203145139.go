package infra

import (
	"time"

	domain "github.com/devfullcycle/catalogo-fc/internal/domain/entity"
	"github.com/devfullcycle/catalogo-fc/pkg/entity"
)

type Category struct {
	ID          entity.ID `gorm:"primaryKey"`
	Name        string
	Description string
	IsActive    bool
	CreatedAt   time.Time
}

func (c *Category) toEntity() *domain.Category {
	return &domain.Category{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		IsActive:    c.IsActive,
		CreatedAt:   c.CreatedAt,
	}
}

func fromEntity(d *domain.Category) Category {
	return Category{
		ID:          d.ID,
		Name:        d.Name,
		Description: d.Description,
		IsActive:    d.IsActive,
		CreatedAt:   d.CreatedAt,
	}
}
