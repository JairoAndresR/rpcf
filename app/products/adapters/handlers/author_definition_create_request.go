package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"rpcf/core/handlers"
	"rpcf/products"
)

type AuthorDefinitionCreateRequest struct {
	Definition string `json:"definition", validate:"required"`
}

func newAuthorDefinitionCreateRequest(c *gin.Context) (*AuthorDefinitionCreateRequest, error) {
	var req AuthorDefinitionCreateRequest
	err := c.BindJSON(&req)

	if err != nil {
		return nil, err
	}

	return &req, nil
}

func (r *AuthorDefinitionCreateRequest) IsValid() *handlers.Error {
	params := make([]handlers.InvalidParam, 0)

	validFields := r.HasValidFields()
	params = append(params, validFields...)

	if len(params) == 0 {
		return nil
	}

	error := handlers.NewErrorWithStatus(handlers.TitleInvalidRequestError, http.StatusBadRequest)
	return error
}

func (r *AuthorDefinitionCreateRequest) HasValidFields() []handlers.InvalidParam {
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

func (r *AuthorDefinitionCreateRequest) GetAuthorDefinition() *products.AuthorDefinition {
	return &products.AuthorDefinition{
		Definition: r.Definition,
	}
}
