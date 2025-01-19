package products

import (
	"net/http"

	"ginexample.com/pkg/models"
	"ginexample.com/pkg/repository"
	"github.com/gin-gonic/gin"
)

type RequestProduct struct {
	ManufacturerId    int     `json:"manufacturer_id"`
	ProductCategoryId int     `json:"productCategory_id"`
	Name              string  `json:"name"`
	Price             float64 `json:"price"`
	Quantity          int     `json:"quantity"`
	ArticleNumber     int     `json:"articleNumber"`
	Description       string  `json:"description"`
	ImageUrl          string  `json:"image_url"`
}

// CreateProduct godoc
// @Summary      Create a product
// @Description  Creating product
// @Tags         products
// @Accept       json
// @Produce      json
// @Security JWT
// @Param        product body   RequestProduct  true  "Add product"
// @Success      201  {object}  models.Product
// @Failure      400  {object}  map[string]string  "Invalid request or error message"
// @Failure      401  {object}  map[string]string  "Invalid request or error message"
// @Failure      403  {object}  map[string]string  "Invalid request or error message"
// @Router       /products [post]
func CreateProduct(c *gin.Context) {
	var newProduct models.Product

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	newProduct.ID = 0

	if !repository.HandleProductFieldsError(c, newProduct) {
		return
	}

	if err := repository.AddProducts([]*models.Product{&newProduct}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newProduct)
}
