package api

import (
	_ "ginexample.com/docs"
	"ginexample.com/pkg/auth"
	"ginexample.com/pkg/handler"
	"ginexample.com/pkg/handler/cart"
	"ginexample.com/pkg/handler/products"
	taskhandler "ginexample.com/pkg/handler/taskHandler"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartServer(port string) {
	router := gin.Default()
	fillEndpoints(router)
	router.Run(port)
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization

// @termsOfService  http://swagger.io/terms/

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func fillEndpoints(router *gin.Engine) {

	router.POST("/registrate", handler.Registrate)
	router.POST("/login", handler.Login)
	router.GET("/users", auth.AuthMiddleware(), handler.GetUsers)

	router.GET("/products", auth.AuthMiddleware(), products.GetProducts)
	router.GET("/products/:id", auth.AuthMiddleware(), products.GetProductsById)
	router.GET("/products/timeout", auth.AuthMiddleware(), products.GetProductsWithTimeout)
	router.POST("/products", auth.AuthMiddleware(), auth.AdminCheck(), products.CreateProduct)
	router.PUT("/products", auth.AuthMiddleware(), auth.AdminCheck(), products.UpdateProduct)
	router.DELETE("/products/:id", auth.AuthMiddleware(), auth.AdminCheck(), products.DeleteProduct)

	router.GET("/cart", auth.AuthMiddleware(), cart.GetAllProductsInCart)
	router.DELETE("/cart/products/:id", auth.AuthMiddleware(), cart.DeleteProductFormCart)
	router.PUT("/cart/products/:id", auth.AuthMiddleware(), cart.AddProductToCart)

	router.POST("/tasks", auth.AuthMiddleware(), taskhandler.CreateTask)
	router.GET("/tasks/:id", auth.AuthMiddleware(), taskhandler.GetTask)
	router.DELETE("/tasks/:id", auth.AuthMiddleware(), taskhandler.CancelTask)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
}
