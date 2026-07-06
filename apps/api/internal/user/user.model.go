package user

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name				string	`json:"user"`
	Email				string	`json:"email" gorm:"unique"`
	Password		string	`json:"-"`
	ValidUntil	*string	`json:"valid_until"`
}
