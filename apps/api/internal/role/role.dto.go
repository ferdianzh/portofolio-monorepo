package role

type CreateRoleDTO struct {
	Name  string `validate:"required,min=3,max=50" json:"name"`
	Alias string `validate:"required,min=3,max=50" json:"alias"`
}

func (d CreateRoleDTO) ToModel() Role {
	return Role{
		Name:  d.Name,
		Alias: d.Alias,
	}
}

type UpdateRoleDTO struct {
	Name  string `validate:"omitempty,min=3,max=50" json:"name"`
	Alias string `validate:"omitempty,min=3,max=50" json:"alias"`
}

func (d UpdateRoleDTO) ToModel() Role {
	return Role{
		Name:  d.Name,
		Alias: d.Alias,
	}
}

type ReplacePermissionsDTO struct {
	PermissionIds []string `validate:"required" json:"permissionIds"`
}
