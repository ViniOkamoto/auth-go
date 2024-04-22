package services

import (
	internal "github.com/viniokamoto/go-store/internal/exception"
	"github.com/viniokamoto/go-store/source/user/domain/exceptions"
	"github.com/viniokamoto/go-store/source/user/domain/models"
	"github.com/viniokamoto/go-store/source/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServices interface {
	GetUsers() (*[]models.UserResponse, *internal.Exception)
	GetUserByID(id string) (*models.UserResponse, *internal.Exception)
	CreateUser(request models.UserCreateRequest) *internal.Exception
	UpdateUser(id string, request models.UserUpdateRequest) (*models.UserResponse, *internal.Exception)
	DeleteUser(id string) *internal.Exception
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
		return nil, internal.DBException()
	}

	var usersResponse []models.UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, models.UserResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,

			Email: user.Email,
			Role: models.RoleResponse{
				ID:   user.Role.ID,
				Name: user.Role.Name,
			},
		})
	}

	return &usersResponse, nil
}

func (service *UserServicesImpl) GetUserByID(id string) (*models.UserResponse, *internal.Exception) {
	user, err := service.userRepository.FindById(id)
	if err != nil {
		return nil, exceptions.UserNotFoundException()
	}
	return &models.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,

		Email: user.Email,
		Role: models.RoleResponse{
			ID:   user.Role.ID,
			Name: user.Role.Name,
		},
	}, nil

}

func (service *UserServicesImpl) CreateUser(request models.UserCreateRequest) *internal.Exception {
	user, err := service.userRepository.FindByEmail(request.Email)

	if err != nil {

		return internal.DBException()
	}

	if user != nil {
		return exceptions.UserEmailAlreadyExistException()
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	request.Password = string(hash)

	_, err = service.userRepository.Create(request)

	if err != nil {
		if err.Error() == "invalidEmail" {
			return exceptions.UserEmailAlreadyExistException()
		}
		return internal.DBException()
	}

	return nil
}

func (service *UserServicesImpl) UpdateUser(id string, request models.UserUpdateRequest) (*models.UserResponse, *internal.Exception) {

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
		Role: models.RoleResponse{
			ID:   userUpdate.Role.ID,
			Name: userUpdate.Role.Name,
		},
	}, nil
}

func (service *UserServicesImpl) DeleteUser(id string) *internal.Exception {
	user, err := service.userRepository.FindById(id)

	if user == nil {
		return exceptions.UserNotFoundException()
	}

	if err != nil {
		return exceptions.UserNotFoundException()
	}

	err = service.userRepository.Delete(*user)

	if err != nil {
		return internal.DBException()
	}

	return nil
}
