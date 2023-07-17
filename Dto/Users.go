package dto

type CreateUserRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Role     string `json:"role" form:"email" `
}

type UpdateUserRequest struct {
	Password string `json:"password" form:"password"`
	Username string `json:"username" form:"username"`
	Role     string `json:"role" form:"role"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
