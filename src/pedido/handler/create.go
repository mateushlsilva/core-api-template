package pedido_handler

import (
	common_model "github.com/cogniia/core-api-template/src/common/model"
	"github.com/cogniia/core-api-template/src/database"
	pedido_entity "github.com/cogniia/core-api-template/src/pedido/entity"
	pedido_model "github.com/cogniia/core-api-template/src/pedido/model"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Creates a new pedido
// @Description	Creates a new pedido 
// @Tags			Pedido
// @Accept			json
// @Produce		json
// @Param			pedido	body		pedido_model.Create	true	"Pedido data"
// @Success		201		{object}	pedido_entity.Pedido		"Created Pedido"
// @Router			/pedido [post]
func CreatePedido(c *fiber.Ctx) error {
	// Parse the request body
	var newPedido pedido_model.Create
	if err := c.BodyParser(&newPedido); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common_model.NewParseJsonError(err).Send())
	}

	// Create the new pedido
	newEntity := pedido_entity.Pedido{
		UsuarioId:     newPedido.UsuarioId,
		Pedido:    newPedido.Pedido,
	}

	// Save the new pedido to the database
	if err := database.Connection().Create(&newEntity).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			common_model.NewApiError("unable to create pedido", err, "gorm.io/gorm").Send(),
		)
	}

	// Return the created pedido (or just a success message)
	return c.Status(fiber.StatusCreated).JSON(newEntity)
}
