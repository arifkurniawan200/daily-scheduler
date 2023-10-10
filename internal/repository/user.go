package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"template/internal/model"
)

type UserHandler struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserHandler{db}
}

func (h UserHandler) FetchUserByFilter(param model.FetchUserParam) ([]model.User, error) {
	var (
		query string
		index int
		datas []model.User
	)

	var total = func(p model.FetchUserParam) int {
		temp := 0
		if len(p.IDs) != 0 {
			temp += len(p.IDs)
		}
		if len(p.Emails) != 0 {
			temp += len(p.Emails)
		}
		if p.BornDate != "" {
			temp += 2 // for date and month only
		}
		return temp
	}

	args := make([]interface{}, total(param))

	if len(param.IDs) != 0 {
		idQuery := "WHERE id IN (%s)"
		placeholders := make([]string, len(param.IDs))
		for i, id := range param.IDs {
			placeholders[i] = "?"
			args[index] = id
			index++
		}
		inClause := strings.Join(placeholders, ",")
		query += fmt.Sprintf(idQuery, inClause)
	}

	if len(param.Emails) != 0 {
		emailQuery := "WHERE email IN (%s)"
		if query != "" {
			emailQuery = "AND email IN (%s)"
		}
		placeholders := make([]string, len(param.Emails))
		for i, email := range param.Emails {
			placeholders[i] = "?"
			args[index] = email
			index++
		}
		inClause := strings.Join(placeholders, ",")
		query += fmt.Sprintf(emailQuery, inClause)
	}

	if param.BornDate != "" {
		queryBorn := `WHERE MONTH(born_date) = ?
			AND DAY(born_date) = ?`
		if query != "" {
			queryBorn = `AND MONTH(born_date) = ?
			AND DAY(born_date) = ?`
		}
		query += queryBorn
		dates := strings.Split(param.BornDate, "-")
		args[index] = dates[1]
		args[index+1] = dates[2]
	}
	finalQuery := fmt.Sprintf(baseGetUser, query)

	fmt.Println(finalQuery)
	fmt.Println(args)

	rows, err := h.db.Query(finalQuery, args...)
	if err != nil {
		return datas, err
	}
	defer rows.Close()

	for rows.Next() {
		var data model.User
		if err = rows.Scan(&data.ID, &data.NIK, &data.FullName, &data.BornPlace, &data.BornDate, &data.IsAdmin, &data.Email, &data.CreatedAt, &data.UpdatedAt, &data.DeletedAt); err != nil {
			return datas, err
		}
		datas = append(datas, data)
	}

	if err = rows.Err(); err != nil {
		return datas, err
	}
	return datas, err
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
