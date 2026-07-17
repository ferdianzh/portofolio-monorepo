package role

import (
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/database"
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/permission"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Create(role *Role) error {
	return database.DB.Create(role).Error
}

func (r *Repository) FindAll() ([]Role, error) {
	var roles []Role

	err := database.DB.Find(&roles).Error

	return roles, err
}

func (r *Repository) FindOne(id string) (*Role, error) {
	var role Role

	err := database.DB.Where("id = ?", id).First(&role).Error

	return &role, err
}

func (r *Repository) Update(role *Role, id string) error {
	var found Role

	if err := database.DB.First(&found, "id = ?", id).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&found).Updates(role).Error; err != nil {
		return err
	}

	return database.DB.First(role, "id = ?", id).Error
}

func (r *Repository) Delete(id string) error {
	var found Role

	if err := database.DB.First(&found, "id = ?", id).Error; err != nil {
		return err
	}

	return database.DB.Delete(&found, "id = ?", id).Error
}

func (r *Repository) ReplacePermissions(permissionIds []string, id string) error {
	var role Role

	if err := database.DB.First(&role, "id = ?", id).Error; err != nil {
		return err
	}

	var permissions []permission.Permission

	if len(permissionIds) > 0 {
		if err := database.DB.
			Where("id IN ?", permissionIds).
			Find(&permissions).Error; err != nil {
			return err
		}
	}

	return database.DB.Model(&role).
		Association("Permissions").
		Replace(permissions)
}
