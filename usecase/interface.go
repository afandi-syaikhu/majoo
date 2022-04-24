package usecase

import (
	"context"

	"github.com/afandi-syaikhu/majoo/model"
)

//go:generate mockgen -destination=mock/auth_mock.go -package=mock github.com/afandi-syaikhu/majoo/usecase AuthUseCase
type AuthUseCase interface {
	Login(ctx context.Context, data model.Auth) (*model.Response, error)
	ValidateToken(ctx context.Context, headerAuth string) (*model.TokenExtraction, error)
}

//go:generate mockgen -destination=mock/merchant_mock.go -package=mock github.com/afandi-syaikhu/majoo/usecase MerchantUseCase
type MerchantUseCase interface {
	GetReportByMerchantID(ctx context.Context, id int64, pagination model.Pagination) (*model.Response, error)
	IsValidMerchantForUser(ctx context.Context, id int64, username string) (bool, error)
}

//go:generate mockgen -destination=mock/outlet_mock.go -package=mock github.com/afandi-syaikhu/majoo/usecase OutletUseCase
type OutletUseCase interface {
	GetReportByOutletID(ctx context.Context, id int64, pagination model.Pagination) (*model.Response, error)
	IsValidOutletForUser(ctx context.Context, id int64, username string) (bool, error)
}
