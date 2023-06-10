package controllers

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/zakirkun/store-poi/v1/services"
	"github.com/zakirkun/store-poi/v1/types"
)

type ProductController struct{}

func (ProductController) Index(w http.ResponseWriter, r *http.Request) {
	services := services.ProductServices{}

	productList := services.GetAll()
	response := types.Reponse{
		Code:    0,
		Message: "Success fetch",
		Data:    productList,
	}

	render.JSON(w, r, response)
}

func (ProductController) Create(w http.ResponseWriter, r *http.Request) {
	request := &types.RequestCreateProduct{}
	if err := render.Bind(r, request); err != nil {
		render.JSON(w, r, render.M{"error": err})
	}

	services := services.ProductServices{}
	createProduct := services.CreateProduct(*request)
	reponse := types.Reponse{
		Code:    0,
		Message: "Create Product",
		Data:    createProduct,
	}

	render.JSON(w, r, reponse)

}
