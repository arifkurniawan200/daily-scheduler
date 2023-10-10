package usecase

import (
	"github.com/labstack/echo/v4"
	"template/internal/model"
)

type UserUcase interface {
	RegisterCustomer(ctx echo.Context, customer model.UserParam) error
	GetUserInfoByEmail(ctx echo.Context, email string) (model.User, error)
	CreateCampaignForBirthdayUser() error
	GetVoucerByUserID(userId int) ([]model.Campaign, error)
	GetListProduct() ([]model.Product, error)
	CreateCampaign(ctx echo.Context, param model.CampaignParam) error
	GetListUsers(ctx echo.Context, param model.FetchUserParam) ([]model.User, error)
}

type TransactionUcase interface {
	CreateTransaction(transaction model.TransactionParam) error
}
