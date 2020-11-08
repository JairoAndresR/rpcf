package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"rpcf/core/handlers"
	"rpcf/products"
)

type ProductDefinitionCreateRequest struct {
	Name       string `json:"name" validate:"required"`
	Definition string `json:"definition", validate:"required"`
}

func newProductDefinitionCreateRequest(c *gin.Context) (*ProductDefinitionCreateRequest, error) {

	var req ProductDefinitionCreateRequest
	err := c.BindJSON(&req)

	if err != nil {
		return nil, err
	}

	return &req, nil
}

func (r *ProductDefinitionCreateRequest) IsValid() *handlers.Error {
	params := make([]handlers.InvalidParam, 0)

	validFields := r.HasValidFields()
	params = append(params, validFields...)

	if len(params) == 0 {
		return nil
	}

	error := handlers.NewErrorWithStatus(handlers.TitleInvalidRequestError, http.StatusBadRequest)
	return error
}

func (r *ProductDefinitionCreateRequest) HasValidFields() []handlers.InvalidParam {
	custom := handlers.NewValidator()
	err := custom.Validator.Struct(r)

	if err == nil {
		return nil
	}

	errors := err.(validator.ValidationErrors)
	if len(errors) == 0 {
		return nil
	}

	params := make([]handlers.InvalidParam, 0)

	for _, e := range errors {
		param := handlers.InvalidParam{
			Name:   handlers.ToSnakeCase(e.Field()),
			Reason: e.Translate(custom.Translator),
			Code:   handlers.RequiredFieldCode,
		}
		params = append(params, param)
	}
	return params
}

func (r *ProductDefinitionCreateRequest) GetProductDefinition() *products.ProductDefinition {
	return &products.ProductDefinition{
		Name:       r.Name,
		Definition: r.Definition,
	}
}
