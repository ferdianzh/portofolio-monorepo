package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID	`gorm:"primaryKey;uuid;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name				string	`json:"user"`
	Email				string	`gorm:"unique" json:"email"`
	Password		string	`json:"-"`
	ValidUntil	*string	`json:"valid_until"`
}
