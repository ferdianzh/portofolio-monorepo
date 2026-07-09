package seeder

import (
	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
    if err := Permission(db); err != nil {
        return err
    }

    // if err := User(db); err != nil {
    //     return err
    // }

    return nil
}