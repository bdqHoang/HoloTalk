package db

import (
	"Auth-microservice/config"
	"Auth-microservice/models"
	"errors"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db_context *gorm.DB
var env = config.LoadEnv()

func Init() error {
	// get url database
	dbUrl := env.DB_URL
	var err error

	Db_context, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return err
	}

	// auto migration to create table
	if err := Db_context.AutoMigrate(
			&models.User{},
			&models.Role{},
		); err != nil {
		return err
	}

	// init data table role
	roles := []models.Role{{Name: "admin"}, {Name: "user"}}

	// check role admin and user exist
	for _, role := range roles {
		var existRole models.Role
		err := Db_context.Where("name = ?", role.Name).First(&existRole).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Không tìm thấy -> tạo mới
			if err := Db_context.Create(&role).Error; err != nil {
				return err
			}
		} else if err != nil {
			// Nếu là lỗi khác -> báo lỗi luôn
			return err
		}
	}

	log.Println("Database connected")
	return nil
}