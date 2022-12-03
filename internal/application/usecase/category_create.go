package application

import (
	"fmt"

	domain "github.com/devfullcycle/catalogo-fc/internal/domain/entity"
	repo "github.com/devfullcycle/catalogo-fc/internal/domain/repository"
	"github.com/devfullcycle/catalogo-fc/pkg/entity"
)

type (
	CategoryCreateInput struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		IsActive    bool   `json:"is_active"`
	}
	CategoryUpdateInput struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		IsActive    bool   `json:"is_active"`
	}
	CategoryOutput struct {
		ID          entity.ID
		Name        string
		Description string
		IsActive    bool
	}
	createCategoryUseCase struct {
		repo repo.CategoryRepository
	}
)

func NewCreateUseCase(repo repo.CategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{repo}
}

// func NewCreateUseCase() *createCategoryUseCase {
// 	return &createCategoryUseCase{}
// }

func (usecase *createCategoryUseCase) Execute(input *CategoryCreateInput) (*CategoryOutput, error) {
	entity, errEntity := domain.NewCategory(
		input.Name,
		input.Description,
		input.IsActive,
	)

	if errEntity != nil {
		return nil, errEntity
	}

	result, err := usecase.repo.Insert(entity)
	if err != nil {
		return nil, err
	}
	fmt.Println("Dentro do usecase => ", entity)
	fmt.Println("Dentro do usecase resultado DB => ", result)

	return &CategoryOutput{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		IsActive:    entity.IsActive,
	}, nil

	// return &CategoryOutput{
	// 	ID:          result.ID,
	// 	Name:        result.Name,
	// 	Description: result.Description,
	// 	IsActive:    result.IsActive,
	// }, nil
}
