package infra

import (
	"encoding/json"
	"net/http"

	application "github.com/devfullcycle/catalogo-fc/internal/application/category"
)

func (rs *CategoriesResource) Create(res http.ResponseWriter, req *http.Request) {
	var dto application.CategoryCreateInput
	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		return
	}
}
