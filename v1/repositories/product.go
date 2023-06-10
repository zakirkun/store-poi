package repositories

import (
	"log"

	"github.com/zakirkun/infra-go/pkg/database"
	"github.com/zakirkun/store-poi/v1/models"
)

type RepositoriesProduct struct{}

func (RepositoriesProduct) GetProduct() (*error, *[]models.Product) {
	db, err := database.DB.OpenDB()
	if err != nil {
		log.Fatalf("Close connection database : %v", err)
	}

	var data []models.Product
	if err := db.Model(models.Product{}).Find(&data).Error; err != nil {
		return &err, nil
	}

	return nil, &data
}

func (RepositoriesProduct) CreateProduct(data models.Product) error {
	db, err := database.DB.OpenDB()
	if err != nil {
		log.Fatalf("Close connection database : %v", err)
	}

	return db.Create(&data).Error
}
