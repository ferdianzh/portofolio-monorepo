package user

import "github.com/google/uuid"

type CreateUserDTO struct {
	Name     string `validate:"required,min=3,max=50" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8,strong_password" json:"password"`
	RoleID   string `validate:"required" json:"role_id"`
}

func (d CreateUserDTO) ToModel() User {
	rid, _ := uuid.Parse(d.RoleID)
	return User{
		Name:     d.Name,
		Email:    d.Email,
		Password: d.Password,
		RoleID:   &rid,
	}
}

type UpdateUserDTO struct {
	Name     string `validate:"omitempty,min=3,max=50" json:"name"`
	Email    string `validate:"omitempty,email" json:"email"`
	Password string `validate:"omitempty,min=8,strong_password" json:"password"`
	RoleID   string `validate:"omitempty" json:"role_id"`
}

func (d UpdateUserDTO) ToModel() User {
	var rid *uuid.UUID
	if d.RoleID != "" {
		tmp, _ := uuid.Parse(d.RoleID)
		rid = &tmp
	}
	return User{
		Name:     d.Name,
		Email:    d.Email,
		Password: d.Password,
		RoleID:   rid,
	}
}
