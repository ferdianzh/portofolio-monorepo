package role

import (
	"time"

	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/permission"
	"github.com/google/uuid"
)

type Role struct {
	ID				uuid.UUID `gorm:"primaryKey;uuid;default:gen_random_uuid()" json:"id"`
	CreatedAt	time.Time `json:"created_at"`
	UpdatedAt	time.Time `json:"updated_at"`

	Name	string	`json:"name"`
	Alias	string	`json:"alias"`

	Permissions	[]permission.Permission	`gorm:"many2many:role_permissions;constraint:OnDelete:CASCADE" json:"-"`
}
