package service

import (
	"log"

	"github.com/Abdelrhmanfdl/user-service/internal/errs"
	"github.com/Abdelrhmanfdl/user-service/internal/models"
	"github.com/Abdelrhmanfdl/user-service/internal/repository/user"
	"github.com/Abdelrhmanfdl/user-service/internal/utils"
	"github.com/Abdelrhmanfdl/user-service/internal/utils/jwt"
)

type UserService struct {
	userRepository user.UserRepository
}

func NewUserService() *UserService {
	userService := &UserService{}
	userService.InitService()
	return userService
}

func (userService *UserService) connectToUserRepository() {
	userService.userRepository = user.NewScyllaUserRepository("127.0.0.1")
}

func (userService *UserService) InitService() {
	userService.connectToUserRepository()
}

func (userService *UserService) SignupUser(user models.DtoSignupRequest) (token string, err error) {
	if _, err := userService.userRepository.GetUserByEmail(user.Email); err != nil {
		if _, ok := err.(*errs.NotFoundError); !ok {
			return "", err
		}
	} else {
		return "", &errs.UserExisting{Message: "user already existing"}
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println("failed to hash password")
		return "", &errs.HashingError{Message: "failed to hash password"}
	}

	id, err := userService.userRepository.CreateUser(&models.User{Username: user.Username, Email: user.Email, Password: hashedPassword})
	if err != nil {
		log.Println("failed to create account")
		return "", err
	}

	return jwt.GenerateJWT(id)
}

func (userService *UserService) LoginUser(userDto models.DtoLoginRequest) (token string, err error) {
	fetchedUser, err := userService.userRepository.GetUserByEmail(userDto.Email)
	if err != nil {
		if _, ok := err.(*errs.NotFoundError); ok {
			return "", &errs.NotFoundError{Message: "wrong email or password"}
		} else {
			return "", err
		}
	}

	if isCorrect := utils.CheckPasswordHash(userDto.Password, fetchedUser.Password); isCorrect {
		return jwt.GenerateJWT(fetchedUser.ID)
	} else {
		return "", &errs.WrongEmailOrPassword{Message: "wrong email or password"}

	}
}

func (userService *UserService) GetUserData(userId string) (user *models.User, err error) {
	user, err = userService.userRepository.GetUserById(userId)
	if err != nil {
		if _, ok := err.(*errs.NotFoundError); ok {
			return nil, &errs.NotFoundError{Message: "user not found"}
		} else {
			return nil, err
		}
	}

	return user, err
}

func (userService *UserService) GetUsersData(userIds []string) (users []models.User, err error) {
	users, err = userService.userRepository.GetUsersByIds(userIds)
	if err != nil {
		return nil, err
	}
	return users, err
}
