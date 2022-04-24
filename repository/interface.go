package repository

import (
	"context"

	"github.com/afandi-syaikhu/majoo/model"
)

//go:generate mockgen -destination=mock/user_mock.go -package=mock github.com/afandi-syaikhu/majoo/repository UserRepository
type UserRepository interface {
	FindByUsernameAndPassword(ctx context.Context, data model.Auth) (*model.User, error)
}
