package infra

import (
	domain "github.com/devfullcycle/catalogo-fc/internal/domain/entity"
	repo "github.com/devfullcycle/catalogo-fc/internal/domain/repository"
	"github.com/devfullcycle/catalogo-fc/pkg/entity"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *repository {
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

func (repository *repository) Update(category *domain.Category) (*domain.Category, error) {
	model := fromEntity(category)

	result := repository.db.Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	model.ID = category.ID

	return category, nil
}

func (repository *repository) Delete(entity.ID) error {
	return nil
}

func (repository *repository) Paginate(repo.SearchParams) (*repo.SearchResult[*domain.Category], error) {
	return nil, nil
}

func (repository *repository) FindByID(entity.ID) (domain.Category, error) {
	return , nil
}
