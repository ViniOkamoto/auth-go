package services

import (
	database "github.com/viniokamoto/go-store/internal/environment/database/exceptions"
	internal "github.com/viniokamoto/go-store/internal/exception"
	"github.com/viniokamoto/go-store/source/user/domain/exceptions"
	"github.com/viniokamoto/go-store/source/user/domain/models"
	repository "github.com/viniokamoto/go-store/source/user/respository"
	"golang.org/x/crypto/bcrypt"
)

type UserServices interface {
	GetUsers() (*[]models.UserResponse, *internal.Exception)
	GetUserByID(id uint) (*models.UserResponse, *internal.Exception)
	CreateUser(request models.UserCreateRequest) *internal.Exception
	UpdateUser(id uint, request models.UserUpdateRequest) (*models.UserResponse, *internal.Exception)
	DeleteUser(id uint) *internal.Exception
}

type UserServicesImpl struct {
	userRepository repository.UserRepository
}

func UserServicesFactory(userRepository repository.UserRepository) UserServices {
	return &UserServicesImpl{userRepository: userRepository}
}

func (service *UserServicesImpl) GetUsers() (*[]models.UserResponse, *internal.Exception) {
	users, err := service.userRepository.FindAll()
	if err != nil {
		return nil, database.DBException()
	}

	var usersResponse []models.UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, models.UserResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,

			Email: user.Email,
			Role:  user.Role,
		})
	}

	return &usersResponse, nil
}

func (service *UserServicesImpl) GetUserByID(id uint) (*models.UserResponse, *internal.Exception) {
	user, err := service.userRepository.FindById(id)
	if err != nil {
		return nil, exceptions.UserNotFoundException()
	}
	return &models.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,

		Email: user.Email,
		Role:  user.Role,
	}, nil

}

func (service *UserServicesImpl) CreateUser(request models.UserCreateRequest) *internal.Exception {
	user, err := service.userRepository.FindByEmail(request.Email)

	if err != nil {
		return database.DBException()
	}

	if user != nil {
		return exceptions.UserEmailAlreadyExistException()
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	request.Password = string(hash)

	_, err = service.userRepository.Create(request)

	if err != nil {
		return database.DBException()
	}

	return nil
}

func (service *UserServicesImpl) UpdateUser(id uint, request models.UserUpdateRequest) (*models.UserResponse, *internal.Exception) {

	user, err := service.userRepository.FindById(id)

	if err != nil {
		return nil, exceptions.UserNotFoundException()
	}

	userUpdate, err := service.userRepository.Update(*user, request)

	if err != nil {
		return nil, exceptions.UserNotFoundException()
	}

	return &models.UserResponse{
		ID:        userUpdate.ID,
		FirstName: userUpdate.FirstName,
		LastName:  userUpdate.LastName,
		Email:     userUpdate.Email,
		Role:      userUpdate.Role,
	}, nil
}

func (service *UserServicesImpl) DeleteUser(id uint) *internal.Exception {
	user, err := service.userRepository.FindById(id)

	if user == nil {
		return exceptions.UserNotFoundException()
	}

	if err != nil {
		return exceptions.UserNotFoundException()
	}

	err = service.userRepository.Delete(*user)

	if err != nil {
		return database.DBException()
	}

	return nil
}
