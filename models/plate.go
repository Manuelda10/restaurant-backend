package models

type Plate struct {
	ID uint `json:"id"`
	Description string `json:"description"`
	Price float64 `json:"price"`
}