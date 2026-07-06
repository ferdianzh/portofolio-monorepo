package project

import (
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/database"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Create(project *Project) error {
	return database.DB.Create(project).Error
}

func (r *Repository) FindAll() ([]Project, error) {
	var projects []Project

	err := database.DB.Find(&projects).Error

	return projects, err
}

func (r *Repository) FindOne(id string) (Project, error) {
	var project Project

	err := database.DB.First(&project, id).Error

	return project, err
}

func (r *Repository) Update(project *Project, id string) error {
	var found Project

	if err := database.DB.First(&found, id).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&found).Updates(project).Error; err != nil {
		return err
	}

	return database.DB.First(project, id).Error
}

func (r *Repository) Delete(id string) error {
	var found Project

	if err := database.DB.First(&found, id).Error; err != nil {
		return err
	}

	return database.DB.Delete(&found, id).Error
}
