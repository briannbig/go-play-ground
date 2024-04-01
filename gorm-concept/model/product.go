package model

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

func (product Product) String() string {
	out, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		return "marshaling error"
	}
	return string(out)
}
