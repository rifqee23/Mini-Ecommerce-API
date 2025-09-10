package dto

type UserResponse struct {
	UserId    int      `json:"user_id"`
	Email     string   `json:"email"`
	Role      RoleType `json:"role"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
}
