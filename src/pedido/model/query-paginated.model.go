package pedido_model

import (
	common_model "github.com/cogniia/core-api-template/src/common/model"
	database_model "github.com/cogniia/core-api-template/src/database/model"
)

type QueryPaginated struct {
	UsuarioId string `json:"usuarioId,omitempty"`
	Pedido string `json:"pedido,omitempty"`

	common_model.UnrequiredId
	database_model.Paginate
	database_model.DateOrder
	database_model.DateWhere
}
