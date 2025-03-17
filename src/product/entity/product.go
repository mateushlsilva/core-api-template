package product_entity

import (
	common_model "github.com/cogniia/core-api-template/src/common/model"
)

type Product struct {
    Name        string  `gorm:"not null" json:"name,omitempty"`
    Price       float64 `gorm:"not null" json:"price,omitempty"`
    Ingredients string  `gorm:"not null" json:"ingredients,omitempty"`
	common_model.Audit
}

