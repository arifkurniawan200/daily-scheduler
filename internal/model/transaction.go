package model

import "time"

type StatusTransaction string

const (
	StatusTransactionSuccess StatusTransaction = "Success"
	StatusTransactionFailed  StatusTransaction = "Failed"
)

type Transaction struct {
	ID         int        `json:"id"`
	UserID     int        `json:"user_id"`
	ProductID  int        `json:"product_id"`
	CampaignID int        `json:"campaign_id"`
	Total      int        `json:"total"`
	Amount     float64    `json:"amount"`
	Status     float64    `json:"status"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
