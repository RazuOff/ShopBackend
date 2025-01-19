package products

import (
	"net/http"

	"ginexample.com/pkg/models"
	"ginexample.com/pkg/repository"
	"github.com/gin-gonic/gin"
)

// GetProductsById godoc
// @Summary      Get product by ID
// @Description  Retrieves a product from the repository by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Security JWT
// @Param        id   path      string  true  "Product ID"
// @Success      200  {object}  models.Product
// @Failure      404  {object}  map[string]string  "Product not found"
// @Failure      401  {object}  map[string]string  "Invalid request or error message"
// @Router       /products/{id} [get]
func GetProductsById(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	product, err := repository.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}
