package handler

import "github.com/tarcisioads/api_escala/schemas"

type CreateMinistryResponse struct {
	Message string                 `json:"message"`
	Data    schemas.MinistryResponse `json:"data"`
}

type DeleteMinistryResponse struct {
	Message string                 `json:"message"`
	Data    schemas.MinistryResponse `json:"data"`
}
type ShowMinistryResponse struct {
	Message string                 `json:"message"`
	Data    schemas.MinistryResponse `json:"data"`
}
type ListMinistryResponse struct {
	Message string                   `json:"message"`
	Data    []schemas.MinistryResponse `json:"data"`
}
type UpdateMinistryResponse struct {
	Message string                 `json:"message"`
	Data    schemas.MinistryResponse `json:"data"`
}
