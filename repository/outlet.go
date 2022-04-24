package repository

import (
	"context"
	"database/sql"

	"github.com/afandi-syaikhu/majoo/constant"
	"github.com/afandi-syaikhu/majoo/model"
)

type Outlet struct {
	DB *sql.DB
}

func NewOutletRepository(db *sql.DB) OutletRepository {
	return &Outlet{
		DB: db,
	}
}

func (_o *Outlet) GetReportByOutletID(ctx context.Context, id int64, pagination model.Pagination) (*[]model.TransactionReport, error) {
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
				t.outlet_id, sum(t.bill_total) as gross_revenue, t.created_at::date as transaction_date
			from
				transactions t 
			inner join
				outlets o on t.outlet_id = o.id 
			where
				t.outlet_id = $1
			group by 
				t.outlet_id, t.created_at::date
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
	rows, err := _o.DB.QueryContext(ctx, query, id, limit, offset)
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

func (_o *Outlet) FindByIDAndUsername(ctx context.Context, id int64, username string) (*model.Outlet, error) {
	query := `
		select
			o.id,
			o.outlet_name,
			m.user_id,
			o.merchant_id,
			m.merchant_name,
			o.created_at,
			o.created_by,
			o.updated_at,
			o.updated_by
		from
			outlets o
			inner join
				merchants m on o.merchant_id = m.id 
			inner join 
				users u on m.user_id = u.id 
		where 
			o.id = $1
			and u.user_name = $2
	`
	rows, err := _o.DB.QueryContext(ctx, query, id, username)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var outlet *model.Outlet
	for rows.Next() {
		outlet = &model.Outlet{}
		err = rows.Scan(&outlet.ID, &outlet.OutletName, &outlet.UserID, &outlet.MerchantID, &outlet.MerchantName, &outlet.CreatedAt, &outlet.CreatedBy, &outlet.UpdatedAt, &outlet.UpdatedBy)
		if err != nil {
			return nil, err
		}
	}

	return outlet, nil
}

func (_o *Outlet) FindByID(ctx context.Context, id int64) (*model.Outlet, error) {
	query := `
		select
			o.id,
			o.outlet_name,
			m.user_id,
			o.merchant_id,
			m.merchant_name,
			o.created_at,
			o.created_by,
			o.updated_at,
			o.updated_by
		from
			outlets o
			inner join
				merchants m on o.merchant_id = m.id
		where 
			o.id = $1
	`
	rows, err := _o.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var outlet *model.Outlet
	for rows.Next() {
		outlet = &model.Outlet{}
		err = rows.Scan(&outlet.ID, &outlet.OutletName, &outlet.UserID, &outlet.MerchantID, &outlet.MerchantName, &outlet.CreatedAt, &outlet.CreatedBy, &outlet.UpdatedAt, &outlet.UpdatedBy)
		if err != nil {
			return nil, err
		}
	}

	return outlet, nil
}
