package usecase

import (
	"context"

	"github.com/afandi-syaikhu/majoo/model"
)

//go:generate mockgen -destination=mock/auth_mock.go -package=mock github.com/afandi-syaikhu/majoo/usecase AuthUseCase
type AuthUseCase interface {
	Login(ctx context.Context, data model.Auth) (*model.Response, error)
}
