package repository

import (
	"context"
	"database/sql"

	"github.com/afandi-syaikhu/majoo/constant"
	"github.com/afandi-syaikhu/majoo/model"
)

type Merchant struct {
	DB *sql.DB
}

func NewMerchantRepository(db *sql.DB) MerchantRepository {
	return &Merchant{
		DB: db,
	}
}

func (_m *Merchant) GetReportByMerchantID(ctx context.Context, id int64, pagination model.Pagination) (*[]model.TransactionReport, error) {
	query := `
		select
			temp_date.date::varchar,
			COALESCE(temp_report.gross_revenue, 0) gross_revenue
		from (
			select i::date as date
			from
				generate_series('2021-11-01', '2021-11-30', '1 day'::interval) i
		) temp_date
		left join (
			select
				t.merchant_id, m.merchant_name, sum(t.bill_total) as gross_revenue, t.created_at::date as transaction_date
			from
				transactions t 
			inner join
				merchants m on t.merchant_id = m.id 
			where
				t.merchant_id = $1
			group by 
				t.merchant_id, m.merchant_name, t.created_at::date
			order by
				t.created_at::date
		) temp_report on temp_date.date = temp_report.transaction_date
		order by
			temp_date.date
		limit $2
		offset $3
	`

	limit := pagination.Limit
	if limit <= 0 {
		limit = constant.ParamLimit
	}

	page := pagination.Page
	if page <= 0 {
		page = constant.ParamPage
	}

	offset := (page - 1) * limit
	rows, err := _m.DB.QueryContext(ctx, query, id, limit, offset)
	if err != nil {
		return nil, err
	}

	transactions := []model.TransactionReport{}
	for rows.Next() {
		t := model.TransactionReport{}
		err = rows.Scan(&t.Date, &t.GrossRevenue)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, t)
	}

	return &transactions, nil
}

func (_m *Merchant) FindByIDAndUsername(ctx context.Context, id int64, username string) (*model.Merchant, error) {
	query := `
		select m.id, m.user_id, m.merchant_name, m.created_at, m.created_by, m.updated_at, m.updated_by
		from
			merchants m
		inner join 
			users u on m.user_id = u.id
		where 
			m.id = $1
			and u.user_name = $2
	`
	rows, err := _m.DB.QueryContext(ctx, query, id, username)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var merchant *model.Merchant
	for rows.Next() {
		merchant = &model.Merchant{}
		err = rows.Scan(&merchant.ID, &merchant.UserID, &merchant.Name, &merchant.CreatedAt, &merchant.CreatedBy, &merchant.UpdatedAt, &merchant.UpdatedBy)
		if err != nil {
			return nil, err
		}
	}

	return merchant, nil
}

func (_m *Merchant) FindByID(ctx context.Context, id int64) (*model.Merchant, error) {
	query := `
		select m.id, m.user_id, m.merchant_name, m.created_at, m.created_by, m.updated_at, m.updated_by
		from
			merchants m
		where 
			m.id = $1
	`
	rows, err := _m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var merchant *model.Merchant
	for rows.Next() {
		merchant = &model.Merchant{}
		err = rows.Scan(&merchant.ID, &merchant.UserID, &merchant.Name, &merchant.CreatedAt, &merchant.CreatedBy, &merchant.UpdatedAt, &merchant.UpdatedBy)
		if err != nil {
			return nil, err
		}
	}

	return merchant, nil
}
