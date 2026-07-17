package user

import (
	"time"

	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/role"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID	`gorm:"primaryKey;uuid;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name				string	`json:"name"`
	Email				string	`gorm:"unique" json:"email"`
	Password		string	`json:"-"`
	ValidUntil	*string	`json:"valid_until"`

	RoleID			*uuid.UUID	`json:"role_id"`
	Role				*role.Role	`gorm:"foreignKey:RoleID" json:"role,omitempty"`
}
