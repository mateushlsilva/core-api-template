package product_handler

import (
	common_model "github.com/cogniia/core-api-template/src/common/model"
	"github.com/cogniia/core-api-template/src/database"
	product_entity "github.com/mateushlsilva/core-api-template/src/product/entity"
	product_model "github.com/mateushlsilva/core-api-template/src/product/model"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Creates a new product
// @Description	Creates a new product item
// @Tags			Product
// @Accept			json
// @Produce		json
// @Param			product	body		product_model.Create	true	"Product data"
// @Success		201		{object}	product_entity.Product	"Created product"
// @Router			/product [post]
func CreateProduct(c *fiber.Ctx) error {
	// Parse the request body
	var newProduct product_model.Create
	if err := c.BodyParser(&newProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common_model.NewParseJsonError(err).Send())
	}

	// Create the new product entity
	newEntity := product_entity.Product{
		Name:        newProduct.Name,
		Price:       newProduct.Price,
		Ingredients: newProduct.Ingredients,
	}

	// Save the new product to the database
	if err := database.Connection().Create(&newEntity).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			common_model.NewApiError("unable to create product", err, "gorm.io/gorm").Send(),
		)
	}

	// Return the created product
	return c.Status(fiber.StatusCreated).JSON(newEntity)
}
