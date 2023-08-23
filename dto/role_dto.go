package dto

import "github.com/google/uuid"

type CreateRoleDTO struct {
    Name        string   `json:"name"`
    Permissions []string `json:"permissions"`
}
type UpdateRoleDTO struct {
    Name        string    `json:"name"`
    Permissions []string  `json:"permissions"`
    ID          uuid.UUID `json:"id"`
}

type DeleteRoleDTO struct {
    ID uuid.UUID `json:"id"`
}
type FindRoleDTO struct {
    ID uuid.UUID `json:"id"`
}
