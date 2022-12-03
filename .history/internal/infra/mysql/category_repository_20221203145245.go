package infra

import (
	domain "github.com/devfullcycle/catalogo-fc/internal/domain/entity"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	db.AutoMigrate(&Category{})
	return &repository{db}
}

func (repository *repository) Insert(category *domain.Category) (*domain.Category, error) {
	model := fromEntity(category)

	result := repository.db.Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	model.ID = category.ID

	return category, nil
}
