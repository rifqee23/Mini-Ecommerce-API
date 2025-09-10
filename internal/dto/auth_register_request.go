package dto

type RoleType string

const (
	RoleUser   RoleType = "user"
	RoleAdmin  RoleType = "admin"
	RoleSeller RoleType = "seller"
)

type AuthRegisterRequest struct {
	Email    string
	Password string
	Role     RoleType
}
