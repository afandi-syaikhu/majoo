package model

import "time"

type Outlet struct {
	ID           int64     `json:"id" db:"id"`
	OutletName   string    `json:"outlet_name" db:"outlet_name"`
	UserID       int64     `json:"user_id" db:"user_id"`
	MerchantID   int64     `json:"merchant_id" db:"merchant_id"`
	MerchantName string    `json:"merchant_name" db:"merchant_name"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	CreatedBy    int64     `json:"created_by" db:"created_by"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy    int64     `json:"updated_by" db:"updated_by"`
}

type OutletTransactionReport struct {
	OutletID     int64                `json:"outlet_id"`
	OutletName   string               `json:"outlet_name"`
	MerchantID   int64                `json:"merchant_id"`
	MerchantName string               `json:"merchant_name"`
	Transactions *[]TransactionReport `json:"transactions"`
}
