package dto

import (
    "github.com/google/uuid"
)

type TokenDTO struct {
    UserID uuid.UUID `json:"userId"`
}

type AccessTokenDTO TokenDTO
type RefreshTokenDTO TokenDTO

type AccountAPIKeyDTO struct {
    IPAddresses []string `json:"ipAddresses"`
    Description string   `json:"description"`
    TokenDTO
}
