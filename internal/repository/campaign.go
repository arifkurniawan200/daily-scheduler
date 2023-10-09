package repository

import (
	"database/sql"
	"template/internal/model"
)

type CampaignHandler struct {
	db *sql.DB
}

func (c CampaignHandler) CreateCampaignTx(tx *sql.Tx, campaign model.Campaign) (int64, error) {
	rows, err := tx.Exec(createCampaign, campaign.Code, campaign.Name, campaign.Amount, campaign.StartDate, campaign.EndDate, campaign.Quota)
	if err != nil {
		return 0, err
	}
	return rows.LastInsertId()
}

func (c CampaignHandler) CampaignUsersTx(tx *sql.Tx, campaignID, userID int) error {
	_, err := tx.Exec(createUserCampaign, userID, campaignID)
	if err != nil {
		return err
	}
	return err
}

func NewCampaignRepository(db *sql.DB) CampaignRepository {
	return &CampaignHandler{db}
}
