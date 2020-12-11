package handlers

import "rpcf/products"

type AuthorResponse struct {
	ID    string `json:"id"`
	Names string `json:"names"`
}

type AuthorListResponse struct {
	Authors []*AuthorResponse `json:"authors"`
	Total   int               `json:"total"`
}

func NewAuthorsResponse(a *products.Author) *AuthorResponse {
	return &AuthorResponse{
		ID:    a.ID,
		Names: a.Names,
	}
}

func NewAuthorListResponse(list []*products.Author) *AuthorListResponse {

	authors := make([]*AuthorResponse, 0)

	for _, a := range list {
		authors = append(authors,NewAuthorsResponse(a) )
	}

	return &AuthorListResponse{
		Authors: authors,
		Total:   len(authors), // TODO: change this value from the query result
	}
}
