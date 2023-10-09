package model

import "time"

type ProductType string
type Model string

const (
	ProductTypeSaham           ProductType = "Saham"
	ProductTypePasaruang       ProductType = "Pasar Uang"
	ProductTypePendapatanTetap ProductType = "Pendapatan Tetap"
)

const (
	ModelKonvensional Model = "Konvensional"
	ModelSyariah      Model = "Syariah"
)

type Product struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	Type           ProductType `json:"type"`
	ModelProduct   Model       `json:"model_product"`
	Price          float64     `json:"price"`
	Profit         float64     `json:"profit"`
	ProductManager string      `json:"product_manager"`
	ReleaseDate    time.Time   `json:"release_date"`
	CreatedAt      time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time  `json:"deleted_at,omitempty" db:"deleted_at"`
}
