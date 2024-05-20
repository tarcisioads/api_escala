package handler

import "github.com/tarcisioads/api_escala/schemas"



type CreateEscalaResponse struct {
	Message string                 `json:"message"`
	Data    schemas.EscalaResponse `json:"data"`
}

type DeleteEscalaResponse struct {
	Message string                 `json:"message"`
	Data    schemas.EscalaResponse `json:"data"`
}
type ShowEscalaResponse struct {
	Message string                 `json:"message"`
	Data    schemas.EscalaResponse `json:"data"`
}
type ListEscalaResponse struct {
	Message string                   `json:"message"`
	Data    []schemas.EscalaResponse `json:"data"`
}
type UpdateEscalaResponse struct {
	Message string                 `json:"message"`
	Data    schemas.EscalaResponse `json:"data"`
}
