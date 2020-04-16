package database

import (
	"log"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func Migrate(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// Initial migration
		{
			ID: "1",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.AutoMigrate(&User{}).Error; err != nil {
					return err
				}
				if err := tx.AutoMigrate(&GarmentModel{}).Error; err != nil {
					return err
				}
				if err := tx.AutoMigrate(&Garment{}).Error; err != nil {
					return err
				}
				if err := tx.Model(&Garment{}).AddForeignKey(
					"garment_model_id",
					"garment_models(id)",
					"CASCADE",
					"RESTRICT",
				).Error; err != nil {
					return err
				}
				if err := tx.AutoMigrate(&OwnedGarment{}).Error; err != nil {
					return err
				}
				if err := tx.Model(&OwnedGarment{}).AddForeignKey(
					"garment_id",
					"garments(id)",
					"CASCADE",
					"RESTRICT",
				).Error; err != nil {
					return err
				}
				if err := tx.Model(&OwnedGarment{}).AddForeignKey(
					"user_id",
					"users(id)",
					"CASCADE",
					"RESTRICT",
				).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.DropTable(&User{}).Error; err != nil {
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
