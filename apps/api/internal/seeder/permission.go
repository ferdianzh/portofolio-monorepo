package seeder

import (
	"strings"

	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/permission"
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Permission(db *gorm.DB) error {
	var permissions []permission.Permission

	r := strings.NewReplacer(
		".", " ",
		"-", " ",
	)

	for _, p := range permission.All {
		name := utils.CapitalizeEachWord(r.Replace(string(p)))
		permissions = append(permissions, permission.Permission{
			Name: name,
			Alias: string(p),
		})
	}

	return db.
		Clauses(clause.OnConflict{
			DoNothing: true,
		}).
		Create(&permissions).Error
}