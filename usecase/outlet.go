package usecase

import (
	"context"

	"github.com/afandi-syaikhu/majoo/constant"
	"github.com/afandi-syaikhu/majoo/model"
	"github.com/afandi-syaikhu/majoo/repository"
	log "github.com/sirupsen/logrus"
)

type Outlet struct {
	OutletRepo repository.OutletRepository
}

func NewOutletUseCase(outletRepo repository.OutletRepository) OutletUseCase {
	return &Outlet{
		OutletRepo: outletRepo,
	}
}

func (_o *Outlet) GetReportByOutletID(ctx context.Context, id int64, pagination model.Pagination) (*model.Response, error) {
	response := &model.Response{}
	outlet, err := _o.OutletRepo.FindByID(ctx, id)
	if err != nil {
		log.Errorf("[%s] => %s", "OutletUC.GetReportByOutletID", err.Error())
		return nil, err
	}

	if outlet == nil {
		response.Message = constant.NotFound
		return response, nil
	}

	transactions, err := _o.OutletRepo.GetReportByOutletID(ctx, id, pagination)
	if err != nil {
		log.Errorf("[%s] => %s", "OutletUC.GetReportByOutletID", err.Error())
		return nil, err
	}

	response.Success = true
	response.Data = &model.OutletTransactionReport{
		OutletID:     outlet.ID,
		OutletName:   outlet.OutletName,
		MerchantID:   outlet.MerchantID,
		MerchantName: outlet.MerchantName,
		Transactions: transactions,
	}

	return response, nil
}

func (_o *Outlet) IsValidOutletForUser(ctx context.Context, id int64, username string) (bool, error) {
	merchant, err := _o.OutletRepo.FindByIDAndUsername(ctx, id, username)
	if err != nil {
		log.Errorf("[%s] => %s", "OutletUC.IsValidOutletForUser", err.Error())
		return false, err
	}

	if merchant == nil {
		return false, nil
	}

	return true, nil
}
