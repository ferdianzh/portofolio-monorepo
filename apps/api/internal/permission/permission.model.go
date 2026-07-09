package permission

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID				uuid.UUID `gorm:"primaryKey;uuid;default:gen_random_uuid()" json:"id"`
	CreatedAt	time.Time `json:"created_at"`
	UpdatedAt	time.Time `json:"updated_at"`

	Name	string	`json:"name"`
	Alias	string	`gorm:"unique" json:"alias"`
}
