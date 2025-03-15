package product_handler

import (
	common_model "github.com/cogniia/core-api-template/src/common/model"
	"github.com/cogniia/core-api-template/src/repository"
	product_entity "github.com/mateushlsilva/core-api-template/src/product/entity"
	product_model "github.com/mateushlsilva/core-api-template/src/product/model"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Update product by ID
// @Description	Updates a product's details by its ID
// @Tags			Product
// @Accept			json
// @Produce		json
// @Param			body	body		product_model.UpdateWithId	true	"Product data to update"
// @Success		200		{object}	product_entity.Product		"Product updated successfully"
// @Router			/product [put]
// @Security		ApiKeyAuth
func UpdateProductByID(c *fiber.Ctx) error {
	var editProduct product_model.UpdateWithId
	if err := c.BodyParser(&editProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common_model.NewParseJsonError(err).Send())
	}

	// Buscar o produto antes de atualizar
	product, err := repository.First(
		product_entity.Product{
			Audit: common_model.Audit{
				Id: editProduct.Id,
			},
		},
		0, nil, nil, "", nil,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			common_model.NewApiError("unable to find product", err, "repository").Send(),
		)
	}

	data := product_entity.Product{
		Name:        editProduct.Name,
		Description: editProduct.Description,
		Price:       editProduct.Price,
	}

	// Atualizar produto no banco
	updatedProduct, err := repository.Updates(
		data,
		&product_entity.Product{
			Audit: common_model.Audit{
				Id: editProduct.Id,
			},
		},
		nil,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			common_model.NewApiError("unable to update product", err, "repository").Send(),
		)
	}

	return c.Status(fiber.StatusOK).JSON(updatedProduct)
}
