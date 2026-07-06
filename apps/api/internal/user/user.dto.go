package user

type CreateUserDTO struct {
	Name     string `validate:"required,min=3,max=50" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8,strong_password" json:"password"`
}

func (d CreateUserDTO) ToModel() User {
	return User{
		Name:     d.Name,
		Email:    d.Email,
		Password: d.Password,
	}
}

type UpdateUserDTO struct {
	Name     string `validate:"omitempty,min=3,max=50" json:"name"`
	Email    string `validate:"omitempty,email" json:"email"`
	Password string `validate:"omitempty,min=8,strong_password" json:"password"`
}

func (d UpdateUserDTO) ToModel() User {
	return User{
		Name:     d.Name,
		Email:    d.Email,
		Password: d.Password,
	}
}
