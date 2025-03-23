package pedido_handler

import (
	"errors"

	common_model "github.com/cogniia/core-api-template/src/common/model"
	pedido_entity "github.com/cogniia/core-api-template/src/pedido/entity"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Gets current pedido
// @Description	Returns the currently authenticated pedido
// @Tags			Pedido
// @Accept			json
// @Produce		json
// @Success		200	{object}	pedido_entity.Pedido	"Current pedido"
// @Router			/pedido/me [get]
// @Security		ApiKeyAuth
func GetCurrentPedido(c *fiber.Ctx) error {
	// Retrieve the authenticated pedido from the context
	pedido, ok := c.Locals("pedido").(*pedido_entity.Pedido)
	if !ok || pedido == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			common_model.NewApiError("failed to retrieve pedido from context locals", errors.New("invalid convertion to type pedido_entity.Pedido"), "handler").Send(),
		)
	}

	return c.Status(fiber.StatusOK).JSON(pedido)
}
