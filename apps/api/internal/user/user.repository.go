package user

import "github.com/ferdianzh/portofolio-monorepo/apps/api/internal/database"

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Create(user *User) error {
	return database.DB.Create(user).Error
}

func (r *Repository) FindAll() ([]User, error) {
	var users []User

	err := database.DB.Preload("Role").Find(&users).Error

	return users, err
}

func (r *Repository) FindOne(id string) (*User, error) {
	var user User

	err := database.DB.Preload("Role").Where("id = ?", id).First(&user).Error

	return &user, err
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	var user User

	err := database.DB.Where("email = ?", email).First(&user).Error

	return &user, err
}

func (r *Repository) Update(user *User, id string) error {
	var found User

	if err := database.DB.First(&found, id).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&found).Updates(user).Error; err != nil {
		return err
	}

	return database.DB.First(user, id).Error
}

func (r *Repository) Delete(id string) error {
	var found User

	if err := database.DB.First(&found, id).Error; err != nil {
		return err
	}

	return database.DB.Delete(&found, id).Error
}
