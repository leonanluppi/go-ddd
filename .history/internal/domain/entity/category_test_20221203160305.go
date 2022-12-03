package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyName_WhenCreateANewCategory_ThenShould_ReceiveAnError(t *testing.T) {
	category, err := NewCategory("", "", true)

	assert.Nil(t, category)
	assert.NotNil(t, err)
	assert.EqualError(t, ErrNameIsRequired, err.Error())
}

func TestGivenAValidParams_WhenCreateANewCategory_ThenShould_ReceiveCreatedCategoryWithAllParams(t *testing.T) {
	category, err := NewCategory("Category 1", "Description", true)

	assert.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, "Category 1", category.Name)
	assert.Equal(t, "Description", category.Description)
}

func TestGivenAValidInactiveCategory_WhenCallActive_ThenShould_ReturnCategoryActivated(t *testing.T) {
	category, err := NewCategory("Category 1", "Description", false)

	assert.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, "Category 1", category.Name)
	assert.Equal(t, "Description", category.Description)
	assert.False(t, category.IsActive)

	category.Activate()

	assert.NoError(t, err)
	assert.NotNil(t, category)
	assert.True(t, category.IsActive)
}

func TestGivenAValidActiveCategory_WhenCallDeActive_ThenShould_ReturnCategoryDeActivated(t *testing.T) {
	category, err := NewCategory("Category 1", "Description", true)

	assert.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, "Category 1", category.Name)
	assert.Equal(t, "Description", category.Description)
	assert.True(t, category.IsActive)

	category.Deactivate()

	assert.NoError(t, err)
	assert.NotNil(t, category)
	assert.False(t, category.IsActive)
}

func TestGivenAValidCategory_WhenCallUpdate_ThenShould_ReturnCategoryUpdated(t *testing.T) {
	category, err := NewCategory("Category 1", "Description", true)

	var (
		expectedUpdatedCategoryName     = "Category Updated"
		expectedDescriptionCategoryName = "New Description"
	)

	assert.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, "Category 1", category.Name)
	assert.Equal(t, "Description", category.Description)
	assert.True(t, category.IsActive)

	uc, errUc := category.Update(expectedUpdatedCategoryName, expectedDescriptionCategoryName)

	assert.NoError(t, errUc)
	assert.NotNil(t, uc)
	assert.Equal(t, expectedUpdatedCategoryName, uc.Name)
	assert.Equal(t, expectedDescriptionCategoryName, uc.Description)
	assert.True(t, uc.IsActive)
}

func TestGivenAValidCategory_WhenCallUpdate_ThenShould_ReceiveAnError(t *testing.T) {
	category, err := NewCategory("Category 1", "Description", true)

	var (
		expectedInvalidUpdatedCategoryName = ""
		expectedDescriptionCategoryName    = "New Description"
	)

	assert.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, "Category 1", category.Name)
	assert.Equal(t, "Description", category.Description)
	assert.True(t, category.IsActive)

	uc, errUc := category.Update(expectedInvalidUpdatedCategoryName, expectedDescriptionCategoryName)

	assert.Nil(t, uc)
	assert.NotNil(t, errUc)
	assert.EqualError(t, ErrNameIsRequired, errUc.Error())
}
