package dto

import "github.com/google/uuid"

type CreateUserDTO struct {
    Name     string    `json:"name"`
    Email    string    `json:"email"`
    Password string    `json:"password"`
    RoleId   uuid.UUID `json:"role_id"`
    RoleName string    `json:"role_name"`
}

type UpdateUserDTO struct {
    UserId   uuid.UUID `json:"id"`
    Name     string    `json:"name"`
    Email    string    `json:"email"`
    Password string    `json:"password"`
    RoleId   uuid.UUID `json:"role_id"`
    RoleName string    `json:"role_name"`
}

/*type DeleteUserDTO struct {
	UserId uuid.UUID `json:"id"`
}*/

/*type GetUserDTO struct {
	UserId uuid.UUID `json:"id"`
}*/
