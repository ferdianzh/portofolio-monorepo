package project

import (
	"time"
)

type Project struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`

	Title					string	`json:"title"`
	Slug					string	`json:"slug" gorm:"unique"`
	IsFeatured		bool		`json:"is_featured" gorm:"default:true"`
	LiveUrl				*string	`json:"live_url"`
	Content				*string	`json:"content"`
}