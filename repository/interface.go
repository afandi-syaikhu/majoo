package repository

import (
	"context"

	"github.com/afandi-syaikhu/majoo/model"
)

//go:generate mockgen -destination=mock/user_mock.go -package=mock github.com/afandi-syaikhu/majoo/repository UserRepository
type UserRepository interface {
	FindByUsernameAndPassword(ctx context.Context, data model.Auth) (*model.User, error)
}

//go:generate mockgen -destination=mock/merchant_mock.go -package=mock github.com/afandi-syaikhu/majoo/repository MerchantRepository
type MerchantRepository interface {
	GetReportByMerchantID(ctx context.Context, id int64, pagination model.Pagination) (*[]model.TransactionReport, error)
	FindByIDAndUsername(ctx context.Context, id int64, username string) (*model.Merchant, error)
	FindByID(ctx context.Context, id int64) (*model.Merchant, error)
}
