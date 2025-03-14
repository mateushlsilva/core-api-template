package product_model

import common_model "github.com/cogniia/core-api-template/src/common/model"

type Update struct {
	Name        string  `json:"name,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Ingredients string  `json:"ingredients,omitempty"`
}

type UpdateWithId struct {
	common_model.RequiredId
	Update
}