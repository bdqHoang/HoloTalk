package db

import (
	"Auth-microservice/cache"
	"Auth-microservice/config"
	"Auth-microservice/models"
	"errors"
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (Db_context *gorm.DB
	env = config.LoadEnv()
	Enforcer *casbin.Enforcer)

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
			&gormadapter.CasbinRule{},
		); err != nil {
		return err
	}

	// init casbin dapper
	adapter, err := gormadapter.NewAdapterByDB(Db_context)
	if err != nil {
		return err
	}

	// init casbin enforcer
	Enforcer, err = casbin.NewEnforcer("auth_model.conf", adapter)
	if err != nil {
		return err
	}

	// init data table role
	roles := []models.Role{{Name: "admin"}, {Name: "user"}}

	// check role admin and user exist
	for _, role := range roles {
		var existRole models.Role
		err := Db_context.Where("name = ?", role.Name).First(&existRole).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// if role not exist -> create role
			if err := Db_context.Create(&role).Error; err != nil {
				return err
			}
		} else if err != nil {
			// if error -> return
			return err
		}
	}

	// add policy to allow admin to access all resources
	Enforcer.AddPolicy("admin", "*", "*")

	// add policy to allow user to access all resources
	Enforcer.AddPolicy("user", "*", "*")

	// init redis cache
	if err := cache.InitRedis(); err != nil {
		return err
	}

	log.Println("Database connected")
	return nil
}