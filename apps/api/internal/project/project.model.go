package project

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID        uuid.UUID      `gorm:"primaryKey;uuid;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`

	Title					string	`json:"title"`
	Slug					string	`json:"slug" gorm:"unique"`
	IsFeatured		bool		`json:"is_featured" gorm:"default:true"`
	LiveUrl				*string	`json:"live_url"`
	Content				*string	`json:"content"`
}