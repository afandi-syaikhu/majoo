package usecase

import (
	"context"

	"github.com/afandi-syaikhu/majoo/constant"
	"github.com/afandi-syaikhu/majoo/model"
	"github.com/afandi-syaikhu/majoo/repository"
	log "github.com/sirupsen/logrus"
)

type Merchant struct {
	MerchantRepo repository.MerchantRepository
}

func NewMerchantUseCase(merchantRepo repository.MerchantRepository) MerchantUseCase {
	return &Merchant{
		MerchantRepo: merchantRepo,
	}
}

func (_m *Merchant) GetReportByMerchantID(ctx context.Context, id int64, pagination model.Pagination) (*model.Response, error) {
	response := &model.Response{}
	merchant, err := _m.MerchantRepo.FindByID(ctx, id)
	if err != nil {
		log.Errorf("[%s] => %s", "MerchantUC.GetReportByMerchantID", err.Error())
		return nil, err
	}

	if merchant == nil {
		response.Message = constant.NotFound
		return response, nil
	}

	transactions, err := _m.MerchantRepo.GetReportByMerchantID(ctx, id, pagination)
	if err != nil {
		log.Errorf("[%s] => %s", "MerchantUC.GetReportByMerchantID", err.Error())
		return nil, err
	}

	response.Success = true
	response.Data = &model.MerchantTransactionReport{
		MerchantID:   merchant.ID,
		MerchantName: merchant.MerchantName,
		Transactions: transactions,
	}

	return response, nil
}

func (_m *Merchant) IsValidMerchantForUser(ctx context.Context, id int64, username string) (bool, error) {
	merchant, err := _m.MerchantRepo.FindByIDAndUsername(ctx, id, username)
	if err != nil {
		log.Errorf("[%s] => %s", "MerchantUC.IsValidMerchantForUser", err.Error())
		return false, err
	}

	if merchant == nil {
		return false, nil
	}

	return true, nil
}
