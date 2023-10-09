package app

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"template/internal/model"
	"template/internal/utils"
	"time"
)

func (u handler) RegisterUser(c echo.Context) error {
	customer := new(model.UserParam)
	if err := c.Bind(customer); err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to register user",
			Error:    err.Error(),
		})
	}

	validator := validator.New()

	// Validasi struktur data customer
	if err := validator.Struct(customer); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseFailed{
			Messages: "invalid payload",
			Error:    err.Error()})
	}

	err := u.User.RegisterCustomer(c, *customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to register user",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, ResponseSuccess{
		Messages: "success register user",
	})
}

func (u handler) LoginUser(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	userInfo, err := u.User.GetUserInfoByEmail(c, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to login",
			Error:    err.Error(),
		})
	}

	if !utils.VerifyPassword(password, userInfo.Password) {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "invalid username/password",
			Error:    "username or password is mismatch",
		})
	}
	claims := &jwtCustomClaims{
		userInfo.Email,
		userInfo.IsAdmin,
		int64(userInfo.ID),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	accessToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "error when generate token",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"token":   accessToken,
	})
}

func (u handler) BuyProduct(c echo.Context) error {
	transaction := new(model.TransactionParam)
	if err := c.Bind(transaction); err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to register user",
			Error:    err.Error(),
		})
	}

	validator := validator.New()

	// Validasi struktur data customer
	if err := validator.Struct(transaction); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseFailed{
			Messages: "invalid payload",
			Error:    err.Error()})
	}

	claims, ok := c.Get("claims").(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid or missing claims",
		})
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to get user id",
		})
	}
	transaction.UserID = int(userID)
	err := u.Transaction.CreateTransaction(*transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to buy product",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, ResponseSuccess{
		Messages: "success buy product",
	})
}

func (u handler) MyVoucher(c echo.Context) error {
	fmt.Println("masuk")
	claims, ok := c.Get("claims").(jwt.MapClaims)
	if !ok {
		fmt.Println("masuk")
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid or missing claims",
		})
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to get user id",
		})
	}

	data, err := u.User.GetVoucerByUserID(int(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to get voucher",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success fetch voucher",
		Data:     data,
	})
}

func (u handler) ListProduct(c echo.Context) error {
	data, err := u.User.GetListProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to get product",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success fetch product",
		Data:     data,
	})
}
