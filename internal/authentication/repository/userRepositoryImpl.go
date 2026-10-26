package repository

import (
	"chatAppGo/internal/authentication/bootstrap"
	"chatAppGo/internal/authentication/entity"
)

type UserRepositoryImpl struct {
	*bootstrap.Postgres
}

func (r *UserRepositoryImpl) Store(user *entity.User) {
	sql, args, err := r.Builder.
		Insert("users").
		Columns("id, username, ").
		Values(user.Id, user.Username, user.PasswordHash, user.Phone, user.Mail, user.CreatedAt, user.LastUpdateAt).
		ToSql()

}
