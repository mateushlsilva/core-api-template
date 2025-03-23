package pedido_model

import common_model "github.com/cogniia/core-api-template/src/common/model"

type Update struct {
	UsuarioId string `json:"usuarioId,omitempty"`
	Pedido string `json:"pedido,omitempty"`
}

type UpdateWithId struct {

	common_model.RequiredId
	Update
}
