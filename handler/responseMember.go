package handler

import "github.com/tarcisioads/api_escala/schemas"

type CreateMemberResponse struct {
	Message string                 `json:"message"`
	Data    schemas.MemberResponse `json:"data"`
}

type DeleteMemberResponse struct {
	Message string                 `json:"message"`
	Data    schemas.MemberResponse `json:"data"`
}
type ShowMemberResponse struct {
	Message string                 `json:"message"`
	Data    schemas.MemberResponse `json:"data"`
}
type ListMemberResponse struct {
	Message string                   `json:"message"`
	Data    []schemas.MemberResponse `json:"data"`
}
type UpdateMemberResponse struct {
	Message string                 `json:"message"`
	Data    schemas.MemberResponse `json:"data"`
}
