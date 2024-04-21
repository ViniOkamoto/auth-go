package repository

import (
	"github.com/viniokamoto/go-store/source/user/domain/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	Create(request models.UserRequest) (models.User, error)
	FindById(id uint) (models.User, error)
	Update(id uint, request models.UserRequest) (models.User, error)
	Delete(id uint) error
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

func (r *UserRepositoryImpl) Create(request models.UserRequest) (models.User, error) {
	user := models.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Username:  request.Username,
		Email:     request.Email,
		Password:  request.Password,
	}
	err := r.db.Create(&user).Error
	return user, err
}

func (r *UserRepositoryImpl) FindById(id uint) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *UserRepositoryImpl) Update(id uint, request models.UserRequest) (models.User, error) {
	user, err := r.FindById(id)
	if err != nil {
		return user, err
	}

	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Username = request.Username
	user.Email = request.Email
	user.Password = request.Password

	err = r.db.Save(&user).Error
	return user, err
}

func (r *UserRepositoryImpl) Delete(id uint) error {
	user, err := r.FindById(id)
	if err != nil {
		return err
	}

	err = r.db.Delete(&user).Error
	return err
}
