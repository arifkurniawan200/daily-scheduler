package repository

import (
	"database/sql"
	"template/internal/model"
)

type CampaignHandler struct {
	db *sql.DB
}

func (c CampaignHandler) UpdateQuotaTx(tx *sql.Tx, total, campaignID int) error {
	_, err := tx.Exec(updateQuotaVoucher, total, campaignID)
	if err != nil {
		return err
	}
	return err
}

func (c CampaignHandler) CampaignUser(userId int) ([]model.Campaign, error) {
	var (
		datas []model.Campaign
		err   error
	)
	rows, err := c.db.Query(getCampaignByUserID, userId)
	if err != nil {
		return datas, err
	}
	defer rows.Close()

	for rows.Next() {
		var data model.Campaign
		if err = rows.Scan(&data.ID, &data.Code, &data.Name, &data.Amount, &data.StartDate, &data.EndDate,
			&data.Quota, &data.CreatedAt, &data.UpdatedAt, &data.DeletedAt,
		); err != nil {
			return datas, err
		}
		datas = append(datas, data)
	}

	if err = rows.Err(); err != nil {
		return datas, err
	}
	return datas, err
}

func (c CampaignHandler) GetCampaignByCode(code string) (model.Campaign, error) {
	var (
		data model.Campaign
		err  error
	)
	rows, err := c.db.Query(getCampaignByCode, code)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&data.ID, &data.Code, &data.Name, &data.Amount, &data.StartDate, &data.EndDate,
			&data.Quota, &data.CreatedAt, &data.UpdatedAt, &data.DeletedAt,
		); err != nil {
			return data, err
		}
	}

	if err = rows.Err(); err != nil {
		return data, err
	}
	return data, err
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
