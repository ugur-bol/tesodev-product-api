package routes

import (
	"tesodev-product-api/controllers"

	"github.com/labstack/echo/v4"
)

// SetupRoutes - Define all routes
func SetupRoutes(e *echo.Echo) {
	e.GET("/products", controllers.ListProducts)
	e.POST("/products", controllers.CreateProduct)
	e.GET("/products/:id", controllers.GetProductByID)
	e.PUT("/products/:id", controllers.UpdateProduct)
	e.DELETE("/products/:id", controllers.DeleteProduct)
}
