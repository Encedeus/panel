package dto

type CreateRoleDTO struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}
type UpdateRoleDTO struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
	Id          int      `json:"id"`
}

type DeleteRoleDTO struct {
	Id int `json:"id"`
}
type GetRoleDTO struct {
	Id int `json:"id"`
}
