package dto

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleId   int    `json:"role_id"`
	RoleName string `json:"role_name"`
}

type UpdateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleId   int    `json:"role_id"`
	RoleName string `json:"role_name"`
}
