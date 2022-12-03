package infra

import (
	"encoding/json"
	"fmt"
	"net/http"

	ab "github.com/devfullcycle/catalogo-fc/internal/application"
	application "github.com/devfullcycle/catalogo-fc/internal/application/category"
	"github.com/devfullcycle/catalogo-fc/pkg/response"
)

func (rs *CategoriesResource) Create(res http.ResponseWriter, req *http.Request) {
	var dto application.CategoryCreateInput
	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		response.Error(res, http.StatusBadRequest, err)
		return
	}

	var usecase ab.UseCase[application.CategoryCreateInput, application.CategoryOutput]
	usecase = application.NewCreateUseCase()
	result, err := usecase.Execute(&dto)
	if err != nil {
		response.Error(res, 500, err)
		return
	}

	fmt.Println("Passou ali no usecase e voltou => ", result)
	// usecase.Execute()
}
