package model

import "time"

type Merchant struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Name      string    `json:"name" db:"merchant_name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	CreatedBy int64     `json:"created_by" db:"created_by"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy int64     `json:"updated_by" db:"updated_by"`
}

type MerchantTransactionReport struct {
	ID           int64                `json:"merchant_id"`
	Name         string               `json:"merchant_name"`
	Transactions *[]TransactionReport `json:"transactions"`
}
