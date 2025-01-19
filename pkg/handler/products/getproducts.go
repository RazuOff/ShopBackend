package products

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"ginexample.com/pkg/models"
	"ginexample.com/pkg/repository"
	"github.com/gin-gonic/gin"
)

// GetProducts godoc
// @Summary      Get products with filters
// @Description  Retrieves a list of products with optional filters and pagination
// @Tags         products
// @Accept       json
// @Produce      json
// @Security JWT
// @Param        page   query     int     false  "Page number (default: 1)"
// @Param        limit  query     int     false  "Number of products per page (default: 10)"
// @Param        name   query     string  false  "Filter by product name"
// @Param        sort   query     string  false  "Sort column"
// @Success      200    {array}   models.Product
// @Failure      400    {object}  map[string]string  "Invalid request"
// @Failure      401  {object}  map[string]string  "Invalid request or error message"
// @Failure      500    {object}  map[string]string  "Internal server error"
// @Router       /products [get]
func GetProducts(c *gin.Context) {
	var prods []models.Product
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	name := c.Query("name")
	sort := c.Query("sort")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	prods, err = repository.GetProducts(limitInt, pageInt, name, sort)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, prods)
}

// GetProductsWithTimeout godoc
// @Summary      Get products with a timeout
// @Description  Retrieves a list of products, but cancels the request if it takes longer than 2 seconds
// @Tags         products
// @Accept       json
// @Produce      json
// @Security JWT
// @Success      200  {array}   models.Product
// @Failure      401  {object}  map[string]string  "Invalid request or error message"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /products/timeout [get]
func GetProductsWithTimeout(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	prods, err := repository.GetProductsCtx(&ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prods)
}
