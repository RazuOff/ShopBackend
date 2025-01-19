package repository

import (
	"ginexample.com/pkg/db/postgre"
	"ginexample.com/pkg/models"
)

func CreateCart(userID int) (models.Cart, error) {

	cart := models.Cart{UserID: userID}
	if err := postgre.DB.Create(&cart).Error; err != nil {
		return models.Cart{}, err
	}

	return cart, nil
}

func GetUserCart(userID int) (int, error) {
	var result struct {
		ID int
	}
	if err := postgre.DB.
		Model(&models.Cart{}).
		Where("user_id = ?", userID).
		Select("id").
		First(&result).Error; err != nil {
		return -1, err
	}
	return result.ID, nil
}

func AddProductToCart(cartID int, productID int) (models.CartProduct, error) {
	cartProduct := models.CartProduct{CartID: cartID, ProductID: productID}
	tx := postgre.DB.Create(&cartProduct)

	return cartProduct, tx.Error
}

func RemoveProductFromCart(cartID int, productID int) error {
	if err := postgre.DB.
		Where("cart_id = ?", cartID).
		Where("product_id = ?", productID).
		Delete(&models.CartProduct{}).Error; err != nil {
		return err
	}
	return nil
}

func GetProductsInCart(cartID int) ([]models.Product, error) {
	var productIDs []int
	var products []models.Product

	// Получение списка product_id для заданного cartID
	if err := postgre.DB.
		Model(&models.CartProduct{}).
		Where("cart_id = ?", cartID).
		Pluck("product_id", &productIDs).
		Error; err != nil {
		return nil, err
	}

	// Получение данных о продуктах по списку product_id
	if err := postgre.DB.
		Where("id IN ?", productIDs).
		Find(&products).
		Error; err != nil {
		return nil, err
	}

	return products, nil
}
