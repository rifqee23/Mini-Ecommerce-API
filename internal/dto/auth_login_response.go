package dto

type AuthLoginResponse struct {
	UserId string   `json:"userId"`
	Email  string   `json:"email"`
	Role   RoleType `json:"role"`
	Token  string   `json:"token"`
}
