package types

import (
	"errors"
	"net/http"

	"github.com/zakirkun/store-poi/v1/models"
)

type ResponseProductList struct {
	Product []models.Product `json:"product"`
}

type RequestCreateProduct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int32  `json:"price"`
	Image       string `json:"image"`
	Category    string `json:"category"`
}

type ResponseCreateProduct struct {
	Status bool `json:"status"`
}

func (b *RequestCreateProduct) Bind(r *http.Request) error {

	if b.Name == "" {
		return errors.New("Field name required")
	}

	return nil
}
