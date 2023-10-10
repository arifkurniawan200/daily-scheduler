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
	FetchUserByFilter(param model.FetchUserParam) ([]model.User, error)
}

type CampaignRepository interface {
	CreateCampaignTx(tx *sql.Tx, campaign model.Campaign) (int64, error)
	CampaignUsersTx(tx *sql.Tx, campaignID, userID int) error
	GetCampaignByCode(code string) (model.Campaign, error)
	CampaignUser(userId int) ([]model.Campaign, error)
	UpdateQuotaTx(tx *sql.Tx, total, userID int) error
}
type TransactionRepository interface {
	CreateTransactionTx(tx *sql.Tx, transaction model.Transaction) error
}

type ProductRepository interface {
	GetProductByID(productID int) (model.Product, error)
	GetProduct() ([]model.Product, error)
}
