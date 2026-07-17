package permission

import "github.com/ferdianzh/portofolio-monorepo/apps/api/internal/database"

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

type FindAllOptions struct {
	RoleID string
}

func (r *Repository) FindAll(opts ...FindAllOptions) ([]Permission, error) {
	var options FindAllOptions
	if len(opts) > 0 {
		options = opts[0]
	}

	var permissions []Permission
	db := database.DB

	if options.RoleID != "" {
		db = db.
			Joins("JOIN role_permissions rp ON rp.permission_id = permissions.id").
			Where("rp.role_id = ?", options.RoleID)
	}

	err := db.Order("alias ASC").Find(&permissions).Error

	return permissions, err
}
