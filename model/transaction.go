package model

type TransactionReport struct {
	GrossRevenue float64 `json:"gross_revenue" db:"gross_revenue"`
	Date         string  `json:"date" db:"date"`
}
