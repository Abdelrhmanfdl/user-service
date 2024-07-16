package user

import "github.com/Abdelrhmanfdl/user-service/internal/models"

type UserRepository interface {
	GetUserById(id string) (user *models.User, err error)
	GetUsersByIds(userIds []string) (user []models.User, err error)
	GetUserByEmail(email string) (user *models.User, err error)
	CreateUser(user *models.User) (id string, err error)
}
