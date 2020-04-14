package database

import (
	"log"

	"github.com/Coffie/fortress/models"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func Migrate(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// Initial migration
		{
			ID: "1",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.AutoMigrate(&models.Flag{}).Error; err != nil {
					return err
				}
				if err := tx.AutoMigrate(&models.TshirtGroup{}).Error; err != nil {
					return err
				}
				if err := tx.Model(&models.TshirtGroup{}).AddForeignKey(
					"flag_id",
					"flags(id)",
					"CASCADE",
					"RESTRICT",
				).Error; err != nil {
					return err
				}
				if err := tx.AutoMigrate(&models.Tshirt{}).Error; err != nil {
					return err
				}
				if err := tx.Model(&models.Tshirt{}).AddForeignKey(
					"tshirt_group_id",
					"tshirt_groups(id)",
					"CASCADE",
					"RESTRICT",
				).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.DropTable(&models.Tshirt{}).Error; err != nil {
					return err
				}
				if err := tx.DropTable(&models.TshirtGroup{}).Error; err != nil {
					return err
				}
				if err := tx.DropTable(&models.Flag{}).Error; err != nil {
					return err
				}
				return nil
			},
		},
	})
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate DB: %v", err)
	}
	log.Printf("DB Migration ran successfully")
}
