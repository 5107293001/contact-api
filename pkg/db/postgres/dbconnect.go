package postgres

import (
	"github.com/5107293001/contact-api/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnetDatabase() *gorm.DB {

	dsn := "host=localhost user=postgres password=bungkot8 dbname=contactapi port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(
		&models.User{},
	)
	if err != nil {

		panic(err)
	}
	return db
}
