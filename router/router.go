package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/zakirkun/store-poi/v1/controllers"
)

func InitRoutes() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Throttle(15))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))

	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, render.M{"message": "hello world!"})
	})

	// api versioning
	r.Route("/v1", func(r chi.Router) {
		// product endpoint
		r.Route("/product", func(r chi.Router) {
			productController := controllers.ProductController{}
			r.Get("/", productController.Index)
			r.Post("/store", productController.Create)
		})
	})

	// 👇 the walking function 🚶‍♂️
	chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
		return nil
	})

	return r
}
