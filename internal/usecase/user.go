package usecase

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/sync/errgroup"
	"template/internal/model"
	"template/internal/repository"
	"template/internal/utils"
	"time"
)

type UserHandler struct {
	u repository.UserRepository
	t repository.TransactionRepository
	c repository.CampaignRepository
	p repository.ProductRepository
}

func (u UserHandler) GetListProduct() ([]model.Product, error) {
	return u.p.GetProduct()
}

func (u UserHandler) GetVoucerByUserID(userId int) ([]model.Campaign, error) {
	return u.c.CampaignUser(userId)
}

const (
	secret = "abc&1*~#^2^#s0^=)^^7%b34"
)

func NewUserUsecase(u repository.UserRepository, t repository.TransactionRepository, c repository.CampaignRepository, p repository.ProductRepository) UserUcase {
	return &UserHandler{u, t, c, p}
}

func (u UserHandler) CreateCampaignForBirthdayUser() error {
	now := time.Now()
	nowStr := now.Format("2006-01-02")
	fmt.Println(nowStr)
	users, err := u.u.GetUserTodayBirthday(nowStr)
	if err != nil {
		return err
	}
	if len(users) == 0 {
		log.Infof("no users birthday in this day %s", nowStr)
		return nil
	}

	campaign := model.Campaign{
		Name:      fmt.Sprintf("Birthday voucher %s", nowStr),
		Amount:    100.000,
		StartDate: now,
		EndDate:   now.AddDate(0, 0, 7),
		Quota:     len(users),
	}
	campaign.GenerateCode()

	tx, err := u.u.BeginTx()
	if err != nil {
		return err
	}

	campaignID, err := u.c.CreateCampaignTx(tx, campaign)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, user := range users {
		err = u.c.CampaignUsersTx(tx, int(campaignID), user.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	log.Infof("success create campaign %s", nowStr)
	return nil
}

func (u UserHandler) GetUserInfoByEmail(ctx echo.Context, email string) (model.User, error) {
	var (
		err error
		g   errgroup.Group
	)
	userInfo, err := u.u.GetUserByEmail(email)
	if err != nil {
		return model.User{}, err
	}
	g.Go(func() error {
		userInfo.NIK, err = utils.Decrypt(userInfo.NIK, secret)
		if err != nil {
			log.Errorf("error when decrypt nik ")
			return err
		}
		return err
	})

	if err = g.Wait(); err != nil {
		return userInfo, err
	}
	return userInfo, err
}

func (u UserHandler) RegisterCustomer(ctx echo.Context, c model.UserParam) error {
	var (
		err error
		g   errgroup.Group
	)

	g.Go(func() error {
		// hash password
		c.Password, err = utils.HashPassword(c.Password)
		if err != nil {
			log.Errorf("error when hash password ")
			return err
		}
		return err
	})

	g.Go(func() error {
		//encrypt sensitive data
		c.NIK, err = utils.Encrypt(c.NIK, secret)
		if err != nil {
			log.Errorf("error when encrypt nik ")
			return err
		}
		return err
	})

	if err = g.Wait(); err != nil {
		return err
	}

	err = u.u.RegisterUser(c)
	if err != nil {
		log.Errorf("[usecase][RegisterCustomer] error when RegisterUser: %s", err.Error())
		return err
	}

	return nil
}
