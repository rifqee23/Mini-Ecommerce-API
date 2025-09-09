package domain

import "time"

type RoleType string

const (
	RoleUser   RoleType = "user"
	RoleAdmin  RoleType = "admin"
	RoleSeller RoleType = "seller"
)

type Users struct {
	UserId    int
	Email     string
	Password  string
	Role      RoleType
	CreatedAt time.Time
	UpdatedAt time.Time
}
