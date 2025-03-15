package product_handler

import (
	common_model "github.com/cogniia/core-api-template/src/common/model"
	"github.com/cogniia/core-api-template/src/repository"
	product_entity "github.com/mateushlsilva/core-api-template/src/product/entity"
	product_model "github.com/mateushlsilva/core-api-template/src/product/model"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Get products paginated
// @Description	Returns a paginated list of products
// @Tags			Product
// @Accept			json
// @Produce		json
// @Param			product	query		product_model.QueryPaginated	true	"Pagination and query parameters"
// @Success		200			{array}		product_entity.Product		"List of products"
// @Router			/product [get]
// @Security		ApiKeyAuth
func GetProducts(c *fiber.Ctx) error {
	query := new(product_model.QueryPaginated)
	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common_model.NewParseJsonError(err).Send())
	}

	products, err := repository.GetPaginated(
		product_entity.Product{
			Name:  query.Name,
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

	return c.Status(fiber.StatusOK).JSON(products)
}
