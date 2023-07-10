package endpoint

import (
	"github.com/Darkren/getmark-home/pkg/data"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/Darkren/getmark-home/pkg/service/auth"
	"github.com/Darkren/getmark-home/pkg/service/pricetag"
)

// CreateRouter creates router with all the API endpoints.
func CreateRouter(log *logrus.Logger, authService auth.Service, usersRepo data.UserRepository,
	productsRepo data.ProductRepository, priceTagService pricetag.Service) http.Handler {
	router := gin.Default()

	router.POST("/users/", AddUser(log, usersRepo))

	router.POST("/auth", Auth(log, authService, usersRepo))

	products := router.Group("/products")
	products.POST("/", AddProduct(log, authService, usersRepo, productsRepo))
	products.DELETE("/:barcode", DeleteProduct(log, authService, usersRepo, productsRepo))
	products.GET("/", ListProducts(log, authService, usersRepo, productsRepo))
	products.GET("/:barcode", GetProduct(log, authService, usersRepo, productsRepo))
	products.POST("/:barcode/tag",
		GeneratePriceTag(log, authService, usersRepo, productsRepo, priceTagService))

	router.GET("/tags/:filepath", GetPriceTag(log, authService, usersRepo, productsRepo))

	return router
}
