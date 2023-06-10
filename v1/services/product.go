package services

import (
	"log"

	"github.com/zakirkun/store-poi/v1/models"
	"github.com/zakirkun/store-poi/v1/repositories"
	"github.com/zakirkun/store-poi/v1/types"
)

type ProductServices struct{}

func (ProductServices) GetAll() *types.ResponseProductList {
	repo := repositories.RepositoriesProduct{}

	err, data := repo.GetProduct()
	if err != nil {
		log.Fatalf("Error get product: %v", err)
		return nil
	}

	return &types.ResponseProductList{
		Product: *data,
	}
}

func (ProductServices) CreateProduct(request types.RequestCreateProduct) *types.ResponseCreateProduct {
	repo := repositories.RepositoriesProduct{}

	err := repo.CreateProduct(models.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Image:       request.Image,
		Category:    request.Category,
	})

	if err != nil {
		log.Fatalf("Failed create product : %v", err)
		return &types.ResponseCreateProduct{
			Status: false,
		}
	}

	return &types.ResponseCreateProduct{
		Status: true,
	}
}
