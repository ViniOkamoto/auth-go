package services

import (
	internal "github.com/viniokamoto/go-store/internal/exception"
	"github.com/viniokamoto/go-store/source/authentication/domain/exceptions"
	"github.com/viniokamoto/go-store/source/authentication/domain/models"
	"github.com/viniokamoto/go-store/source/authentication/jwt"
	userModels "github.com/viniokamoto/go-store/source/user/domain/models"
	userRepository "github.com/viniokamoto/go-store/source/user/repository"
	"github.com/viniokamoto/go-store/source/user/services"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Login(request models.AuthenticationRequest) (*models.AuthenticationResponse, *internal.Exception)
	Register(request models.RegisterRequest) *internal.Exception
}

type AuthService struct {
	userRepository userRepository.UserRepository
	userService    services.UserServices
	jwt            jwt.JWTInterface
}

func AuthServiceFactory(userRepository userRepository.UserRepository, userService services.UserServices) AuthServiceInterface {
	jwt := jwt.Instance
	return &AuthService{userRepository: userRepository, userService: userService, jwt: jwt}
}

func (service *AuthService) Login(request models.AuthenticationRequest) (*models.AuthenticationResponse, *internal.Exception) {
	email := request.Email
	password := request.Password

	user, err := service.userRepository.FindByEmail(email)

	if err != nil {
		return nil, internal.DBException()
	}

	if user == nil {
		return nil, exceptions.UserNotFoundException()
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, exceptions.InvalidPasswordException()
	}

	token, expiresIn, err := service.jwt.GenerateAccessToken(user.ID, user.RoleID)

	if err != nil {
		return nil, internal.ServerException()
	}

	return &models.AuthenticationResponse{
			AccessToken: token,
			ExpiresIn:   expiresIn,
		},
		nil

}

func (service *AuthService) Register(request models.RegisterRequest) *internal.Exception {
	exceptions := service.userService.CreateUser(userModels.UserCreateRequest{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
		RoleID:    request.RoleID,
	})

	if exceptions != nil {
		return exceptions
	}

	return nil
}
