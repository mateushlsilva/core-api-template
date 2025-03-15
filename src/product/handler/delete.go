package product_handler

import (
	common_model "github.com/cogniia/core-api-template/src/common/model"
	"github.com/cogniia/core-api-template/src/repository"
	product_entity "github.com/mateushlsilva/core-api-template/src/product/entity"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Delete product by ID
// @Description	Deletes a product by its ID (only accessible by admins)
// @Tags			Product
// @Accept			json
// @Produce		json
// @Param			body	body		common_model.RequiredId	true	"Product ID to delete"
// @Success		204 "Product deleted successfully"
// @Router			/product [delete]
// @Security		ApiKeyAuth
func DeleteProductByID(c *fiber.Ctx) error {
	var reqBody common_model.RequiredId
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			common_model.NewParseJsonError(err).Send(),
		)
	}

	product, err := repository.First(
		product_entity.Product{
			Audit: common_model.Audit{
				Id: reqBody.Id,
			},
		},
		0, nil, nil, "", nil,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			common_model.NewApiError("unable to find product", err, "repository").Send(),
		)
	}

	err = repository.DeleteById[product_entity.Product](reqBody.Id, nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			common_model.NewApiError("unable to delete product", err, "repository").Send(),
		)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
