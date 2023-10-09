package repository

import (
	"database/sql"
	"template/internal/model"
)

type ProductHandler struct {
	db *sql.DB
}

func (p ProductHandler) GetProductByID(productID int) (model.Product, error) {
	var (
		data model.Product
		err  error
	)
	rows, err := p.db.Query(getProductById, productID)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&data.ID, &data.Name, &data.Type, &data.ModelProduct, &data.Price, &data.Profit,
			&data.ProductManager, &data.ReleaseDate, &data.CreatedAt, &data.UpdatedAt, &data.DeletedAt,
		); err != nil {
			return data, err
		}
	}

	if err = rows.Err(); err != nil {
		return data, err
	}
	return data, err
}

func (p ProductHandler) GetProduct() ([]model.Product, error) {
	var (
		datas []model.Product
		err   error
	)
	rows, err := p.db.Query(getAllProduct)
	if err != nil {
		return datas, err
	}
	defer rows.Close()

	for rows.Next() {
		var data model.Product
		if err = rows.Scan(&data.ID, &data.Name, &data.Type, &data.ModelProduct, &data.Price, &data.Profit,
			&data.ProductManager, &data.ReleaseDate, &data.CreatedAt, &data.UpdatedAt, &data.DeletedAt,
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

func NewProductRepository(db *sql.DB) ProductRepository {
	return &ProductHandler{db}
}
