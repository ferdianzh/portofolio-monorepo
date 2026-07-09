package seeder

import (
	"os"

	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/user"
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func User(db *gorm.DB) error {
	admin := user.User{
		Name: os.Getenv("DEFAULT_ADMIN_NAME"),
		Email: os.Getenv("DEFAULT_ADMIN_EMAIL"),
		Password: utils.HashPassword(os.Getenv("DEFAULT_ADMIN_PASSWORD")),
	}

	return db.
		Clauses(clause.OnConflict{
			DoNothing: true,
		}).
		Create(&admin).Error
}