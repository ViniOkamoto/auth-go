package repository

import (
	"fmt"

	"github.com/viniokamoto/go-store/source/user/domain/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	Create(request models.UserCreateRequest) (models.User, error)
	FindById(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user models.User, request models.UserUpdateRequest) (models.User, error)
	Delete(user models.User) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func UserRepositoryFactory(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *UserRepositoryImpl) FindById(id string) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Role").First(&user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user *models.User
	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return user, err

}

func (r *UserRepositoryImpl) Create(request models.UserCreateRequest) (models.User, error) {
	var existingUser models.User
	if err := r.db.Where("email = ?", request.Email).First(&existingUser).Error; err != gorm.ErrRecordNotFound {
		if err != nil {
			return models.User{}, err
		}
		return models.User{}, fmt.Errorf("invalidEmail")
	}

	user := models.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
	}
	err := r.db.Create(&user).Error

	return user, err
}

func (r *UserRepositoryImpl) Update(user models.User, request models.UserUpdateRequest) (models.User, error) {

	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email

	err := r.db.Save(&user).Error
	return user, err
}

func (r *UserRepositoryImpl) Delete(user models.User) error {
	err := r.db.Delete(&user).Error
	return err
}
