package repository

import (
	"database/sql"
	"template/internal/model"
)

type UserRepository interface {
	RegisterUser(user model.UserParam) error
	GetUserByEmail(email string) (model.User, error)
	GetUserTodayBirthday(date string) ([]model.User, error)
	BeginTx() (*sql.Tx, error)
}

type CampaignRepository interface {
	CreateCampaignTx(tx *sql.Tx, campaign model.Campaign) (int64, error)
	CampaignUsersTx(tx *sql.Tx, campaignID, userID int) error
}
type TransactionRepository interface {
}
