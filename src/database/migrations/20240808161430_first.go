package migrations

import (
	"context"
	"database/sql"

	"github.com/cogniia/core-api-template/src/database"
	entity "github.com/cogniia/core-api-template/src/user/entity"
	product_entity "github.com/mateushlsilva/core-api-template/src/product/entity"
	"github.com/pressly/goose/v3"
	"github.com/pterm/pterm"
)

func init() {
	goose.AddMigrationContext(upFirst, downFirst)
}

func upFirst(ctx context.Context, tx *sql.Tx) error {
	db := database.Connection()
	db.AutoMigrate(&product_entity.Product{})

	var existingUser entity.User
	var existingProduct product_entity.Product
	errProduct := db.First(&existingProduct).Error
	errUser := db.First(&existingUser).Error

	if errUser != nil {
		pterm.DefaultLogger.Info("No user in database.")
	} else {
		pterm.DefaultLogger.Info("User found in database.")
	}
	if errProduct != nil {
		pterm.DefaultLogger.Info("No Product in database.")
	} else {
		pterm.DefaultLogger.Info("Product found in database.")
	}

	return nil
}

func downFirst(ctx context.Context, tx *sql.Tx) error {
	var existingUser entity.User
	var existingProduct product_entity.Product
	errProduct := database.Connection().First(&existingProduct).Error
	errUser := database.Connection().First(&existingUser).Error

	if errUser != nil {
		pterm.DefaultLogger.Info("No user in database.")
	} else {
		pterm.DefaultLogger.Info("User found in database.")
	}
	if errProduct != nil {
		pterm.DefaultLogger.Info("No Product in database.")
	} else {
		pterm.DefaultLogger.Info("Product found in database.")
	}

	return nil
}