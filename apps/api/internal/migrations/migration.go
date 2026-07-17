package migrations

import (
	"log"

	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/database"
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/permission"
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/project"
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/role"
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/user"
)

func Migrate() {
	migErr := database.DB.AutoMigrate(
		&permission.Permission{},
		&role.Role{},
		&user.User{},
		&project.Project{},
	)

	if migErr != nil {
		log.Fatalf("migration error: %v", migErr)
	}

	log.Println("migration success")
}
