package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"template/internal/model"
)

func (u handler) CreateCampaign(c echo.Context) error {
	campaign := new(model.CampaignParam)
	if err := c.Bind(campaign); err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to register user",
			Error:    err.Error(),
		})
	}

	validator := validator.New()

	if err := validator.Struct(campaign); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseFailed{
			Messages: "invalid payload",
			Error:    err.Error()})
	}

	err := u.User.CreateCampaign(c, *campaign)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to create campaign",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, ResponseSuccess{
		Messages: "success create campaign",
	})
}

func (u handler) FetchUser(c echo.Context) error {
	user := new(model.FetchUserParam)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to register user",
			Error:    err.Error(),
		})
	}

	validator := validator.New()

	if err := validator.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseFailed{
			Messages: "invalid payload",
			Error:    err.Error()})
	}

	datas, err := u.User.GetListUsers(c, *user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to create campaign",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success fetch users",
		Data:     datas,
	})
}

func (u handler) TriggerCron(c echo.Context) error {

	err := u.User.CreateCampaignForBirthdayUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to trigger cron",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success trigger cron",
	})
}
