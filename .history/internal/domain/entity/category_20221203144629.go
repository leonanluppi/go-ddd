package domain

import (
	"time"

	err "github.com/devfullcycle/catalogo-fc/internal/domain/error"
	domain "github.com/devfullcycle/catalogo-fc/internal/domain/event"
	"github.com/devfullcycle/catalogo-fc/pkg/entity"
)

var (
	ErrIDIsRequired   = err.IsRequired("id")
	ErrInvalidID      = err.IsInvalid("id")
	ErrNameIsRequired = err.IsRequired("name")
)

type Category struct {
	AggregateRoot
	ID          entity.ID
	Name        string
	Description string
	IsActive    bool
	CreatedAt   time.Time
}

func (category *Category) Validate() error {
	if category.ID.String() == "" {
		return ErrIDIsRequired
	}
	if category.Name == "" {
		return ErrNameIsRequired
	}
	if _, err := entity.ParseID(category.ID.String()); err != nil {
		return ErrInvalidID
	}
	return nil
}

func NewCategory(name, description string, isActive bool) (*Category, error) {
	category := &Category{
		ID:          entity.NewID(),
		Name:        name,
		Description: description,
		IsActive:    isActive,
		CreatedAt:   time.Now(),
	}

	err := category.Validate()
	if err != nil {
		return nil, err
	}

	category.AddEvent(&domain.CategoryCreated{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive,
	})

	return category, nil
}

func (category *Category) Update(name, description string) (*Category, error) {
	category.Name = name
	category.Description = description

	err := category.Validate()
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (category *Category) Activate() {
	category.IsActive = true
}

func (category *Category) Deactivate() {
	category.IsActive = false
}
