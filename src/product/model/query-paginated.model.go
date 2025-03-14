package product_model

import (
	common_model "github.com/cogniia/core-api-template/src/common/model"
	database_model "github.com/cogniia/core-api-template/src/database/model"
)

type QueryPaginated struct {
	Name        string  `json:"name,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Ingredients string  `json:"ingredients,omitempty"`

	common_model.UnrequiredId
	database_model.Paginate
	database_model.DateOrder
	database_model.DateWhere
}
