package cart

import (
	"net/http"
	"strconv"

	"ginexample.com/pkg/repository"
	"github.com/gin-gonic/gin"
)

// GetAllProductsInCart godoc
// @Summary      Get products from cart of current User
// @Description  Retrieves a list of products
// @Tags         cart
// @Produce      json
// @Security JWT
// @Success      200  {array}   []models.Product "All products in cart"
// @Failure      401  {object}  map[string]string  "Invalid request or error message"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /cart [get]
func GetAllProductsInCart(c *gin.Context) {
	userID, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userName not found in Claims. Need to authorize "})
		return
	}

	cartID, err := repository.GetUserCart(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	products, err := repository.GetProductsInCart(cartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// DeleteProductFromCart godoc
// @Summary      Delete a product from the cart
// @Description  Removes a product from the user's cart
// @Tags         cart
// @Produce      json
// @Security     JWT
// @Param        id      path      int                 true  "Product ID to remove"
// @Success      200     {array}   []models.Product      "Updated list of products in the cart"
// @Failure      400     {object}  map[string]string   "Invalid request or failed to remove product"
// @Failure      401     {object}  map[string]string   "Unauthorized, user ID not found"
// @Failure      500     {object}  map[string]string   "Internal server error"
// @Router       /cart/products/{id} [delete]
func DeleteProductFormCart(c *gin.Context) {
	productIDstr := c.Param("id")

	productID, err := strconv.Atoi(productIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "product id invalid"})
		return
	}

	userID, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userName not found in Claims. Need to authorize "})
		return
	}

	cartID, err := repository.GetUserCart(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err = repository.RemoveProductFromCart(cartID, productID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	existingProds, err := repository.GetProductsInCart(cartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingProds)
}

// AddProductToCart godoc
// @Summary      Add a product into the cart
// @Description  Add  a product in the user's cart
// @Tags         cart
// @Produce      json
// @Security     JWT
// @Param        id      path      int                 true  "Product ID to add"
// @Success      200     {array}   []models.Product      "Updated list of products in the cart"
// @Failure      400     {object}  map[string]string   "Invalid request or failed to remove product"
// @Failure      401     {object}  map[string]string   "Unauthorized, user ID not found"
// @Failure      500     {object}  map[string]string   "Internal server error"
// @Router       /cart/products/{id} [put]
func AddProductToCart(c *gin.Context) {
	productIDstr := c.Param("id")

	productID, err := strconv.Atoi(productIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "product id invalid"})
		return
	}

	userID, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userName not found in Claims. Need to authorize "})
		return
	}

	cartID, err := repository.GetUserCart(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if _, err := repository.AddProductToCart(cartID, productID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	existingProds, err := repository.GetProductsInCart(cartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingProds)
}
