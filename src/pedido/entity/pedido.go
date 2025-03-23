package pedido_entity

import (
	common_model "github.com/cogniia/core-api-template/src/common/model"
	
)

type Pedido struct {
	UsuarioId string           `json:"usuarioId,omitempty" gorm:"not null"`
	Pedido    string           `json:"pedido,omitempty" gorm:"not null"`

	common_model.Audit
}

