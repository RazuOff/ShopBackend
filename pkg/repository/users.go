package repository

import (
	"log"
	"strconv"

	"ginexample.com/pkg/db/postgre"
	"ginexample.com/pkg/models"
)

func GetUserByLogin(login string) (models.User, error) {
	var user models.User
	err := postgre.DB.Where("username = ?", login).First(&user).Error
	return user, err
}
func GetUsers() ([]models.User, error) {
	var users []models.User
	err := postgre.DB.Find(&users).Error
	return users, err
}

func AddUser(login string, password string) error {
	user := models.User{Username: login, Password: password}

	if err := postgre.DB.Create(&user).Error; err != nil {
		return err
	}

	if _, err := CreateCart(user.ID); err != nil {
		log.Println("TRY TO CREATE CART FAIL, USER_ID=" + strconv.Itoa(user.ID))
		return err
	}

	return nil
}
