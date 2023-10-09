package repository

import (
	"context"
	"database/sql"
	"strings"
	"template/internal/model"
)

type UserHandler struct {
	db *sql.DB
}

func (h UserHandler) GetUserTodayBirthday(date string) ([]model.User, error) {
	var (
		datas []model.User
		err   error
	)
	args := strings.Split(date, "-")
	rows, err := h.db.Query(getUserBirthdayByDate, args[1], args[2])
	if err != nil {
		return datas, err
	}
	defer rows.Close()

	for rows.Next() {
		var data model.User
		if err = rows.Scan(&data.ID, &data.NIK, &data.FullName, &data.BornPlace, &data.BornDate, &data.IsAdmin, &data.Email, &data.Password, &data.CreatedAt, &data.UpdatedAt, &data.DeletedAt); err != nil {
			return datas, err
		}
		datas = append(datas, data)
	}

	if err = rows.Err(); err != nil {
		return datas, err
	}
	return datas, err
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserHandler{db}
}

func (h UserHandler) BeginTx() (*sql.Tx, error) {
	return h.db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
}

func (h UserHandler) RegisterUser(c model.UserParam) error {
	_, err := h.db.Exec(insertNewCostumer, c.NIK, c.FullName, c.BornPlace, c.BornDate, false, c.Email, c.Password)
	if err != nil {
		return err
	}
	return err
}

func (h UserHandler) GetUserByEmail(email string) (model.User, error) {
	var (
		data model.User
		err  error
	)
	rows, err := h.db.Query(getCostumerByEmail, email)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&data.ID, &data.NIK, &data.FullName, &data.BornPlace, &data.BornDate, &data.IsAdmin,
			&data.Email, &data.Password,
			&data.CreatedAt, &data.UpdatedAt, &data.DeletedAt,
		); err != nil {
			return data, err
		}
	}

	if err = rows.Err(); err != nil {
		return data, err
	}
	return data, err
}
