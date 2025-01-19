package repository

import (
	"context"
	"errors"
	"net/http"

	"ginexample.com/pkg/db/postgre"
	"ginexample.com/pkg/models"
	"github.com/gin-gonic/gin"
)

func HandleProductFieldsError(c *gin.Context, prod models.Product) bool {

	if prod.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Цена не может быть <=0"})
		return false
	}

	if prod.Quantity < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Количество не может быть отрицательным"})
		return false
	}

	return true
}

func GetProducts(limit int, page int, searchString string, sort string) ([]models.Product, error) {
	var products []models.Product

	offset := (page - 1) * limit
	query := postgre.DB.Limit(limit).Offset(offset)
	if searchString != "" {
		query = query.Where("name ILIKE ?", "%"+searchString+"%")
	}
	if sort != "" {
		query = query.Order(sort + " DESC")
	}

	err := query.Find(&products).Error
	return products, err
}

func GetProductsCtx(ctx *context.Context) ([]models.Product, error) {
	var products []models.Product
	err := postgre.DB.WithContext(*ctx).Find(&products).Error

	return products, err

}

func AddProducts(products []*models.Product) error {

	return postgre.DB.Create(products).Error
}

func GetProductById(id string) (models.Product, error) {
	var product models.Product
	err := postgre.DB.Where("id = ?", id).First(&product).Error
	return product, err
}

func DeleteProduct(prod models.Product) error {

	return postgre.DB.Delete(&prod).Error
}

func UpdateProduct(prod models.Product) error {
	if err := postgre.DB.Where("id = ?", prod.ID).First(models.Product{}).Error; err != nil {
		return errors.New("Такого id нет")
	}
	return postgre.DB.Save(prod).Error
}
