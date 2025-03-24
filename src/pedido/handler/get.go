package pedido_handler

import (
	common_model "github.com/cogniia/core-api-template/src/common/model"
	"github.com/cogniia/core-api-template/src/repository"
	pedido_entity "github.com/cogniia/core-api-template/src/pedido/entity"
	pedido_model "github.com/cogniia/core-api-template/src/pedido/model"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Get pedidos paginated
// @Description	Returns a paginated list of pedidos
// @Tags			Pedido
// @Accept			json
// @Produce		json
// @Param			pedido	query		pedido_model.QueryPaginated	true	"Pagination and query parameters"
// @Success		200			{array}		pedido_entity.Pedido			"List of pedidos"
// @Router			/pedido [get]
// @Security		ApiKeyAuth
func Get(c *fiber.Ctx) error {
	query := new(pedido_model.QueryPaginated)
	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common_model.NewParseJsonError(err).Send())
	}

	pedidos, err := repository.GetPaginated(
		pedido_entity.Pedido{
			UsuarioId:  query.UsuarioId,
			Pedido: query.Pedido,
			Audit: common_model.Audit{Id: query.Id},
		},
		&query.Paginate,
		&query.DateOrder,
		&query.DateWhere,
		"",
		nil,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common_model.NewApiError("unable to get paginated", err, "repository").Send())
	}

	return c.Status(fiber.StatusOK).JSON(pedidos)
}
