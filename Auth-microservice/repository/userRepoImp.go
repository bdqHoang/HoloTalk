package repository

import (
	"errors"

	"gorm.io/gorm"

	"Auth-microservice/Interface/IRepository"
	"Auth-microservice/models"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IRepository.IUserRepo {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Role").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Role").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryImpl) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepositoryImpl) Delete(id uint) error {
	result := r.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

