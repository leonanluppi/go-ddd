package domain

import (
	"time"
	// categories_events "fc/go-fc3/internal/domain/categories/events"
	// shared_domain "fc/go-fc3/internal/domain/shared"
	// domain_errs "fc/go-fc3/internal/domain/shared/errs"
	// base "fc/go-fc3/pkg/entity"
)

type Category struct {
	shared_domain.AggregateRoot

	ID          base.ID
	Name        string
	Description string
	IsActive    bool
	CreatedAt   time.Time
}

func (category *Category) IsValid() error {
	if category.Name == "" {
		return domain_errs.IsRequired("name")
	}

	return nil
}

func NewEntity(name string, description string, isActive bool) (*Category, error) {
	category := &Category{
		ID:          base.NewID(),
		Name:        name,
		Description: description,
		IsActive:    isActive,
		CreatedAt:   time.Now(),
	}

	err := category.IsValid()
	if err != nil {
		return nil, err
	}

	category.AddEvent(&categories_events.CategoryCreated{
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

	err := category.IsValid()
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
