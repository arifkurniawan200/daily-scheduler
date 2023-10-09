package usecase

import (
	"fmt"
	"template/internal/model"
	"template/internal/repository"
	"time"
)

type TransactionHandler struct {
	t repository.TransactionRepository
	u repository.UserRepository
	p repository.ProductRepository
	c repository.CampaignRepository
}

func NewTransactionsUsecase(t repository.TransactionRepository, u repository.UserRepository, p repository.ProductRepository, c repository.CampaignRepository) TransactionUcase {
	return &TransactionHandler{t, u, p, c}
}

func (t TransactionHandler) CreateTransaction(param model.TransactionParam) error {
	var (
		campaign model.Campaign
		err      error
	)
	if param.CampaignCode != "" {
		campaign, err = t.c.GetCampaignByCode(param.CampaignCode)
		if err != nil {
			return err
		}
		if campaign.ID == 0 {
			return fmt.Errorf("voucher code not found")
		}

		if campaign.Quota <= 0 {
			return fmt.Errorf("the voucher has expired, you cannot use it anymore")
		}
	}
	product, err := t.p.GetProductByID(param.ProductID)
	if err != nil {
		return err
	}
	if product.ID == 0 {
		return fmt.Errorf("product not found")
	}

	if product.ReleaseDate.After(time.Now()) {
		return fmt.Errorf("product has not released yet")
	}

	tx, err := t.u.BeginTx()
	if err != nil {
		return err
	}

	amountShouldBuy := (product.Price * float64(param.Total)) - campaign.Amount

	err = t.t.CreateTransactionTx(tx, model.Transaction{
		UserID:     param.UserID,
		ProductID:  product.ID,
		CampaignID: campaign.ID,
		Total:      param.Total,
		Amount:     amountShouldBuy,
		Status:     "Success",
	})
	if err != nil {
		tx.Rollback()
		return err
	}
	if campaign.ID != 0 {
		err = t.c.UpdateQuotaTx(tx, campaign.Quota-1, param.UserID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return err
}
