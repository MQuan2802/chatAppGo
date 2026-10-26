package repository

import "chatAppGo/internal/authentication/entity"

type (
	UserRepo interface {
		Store(user *entity.User) error
		GetById(userId string) (entity.User, error)
		GetByMail(mail string) (entity.User, error)
		GetByPhone(phone string) (entity.User, error)
	}
)
