package handler

import (
	"net/http"

	"ginexample.com/pkg/models"
	"ginexample.com/pkg/repository"
	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary      Get all users
// @Description  Retrieves a list of all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Security JWT
// @Success      200  {array}   models.User  "List of users"
// @Failure      401  {object}  map[string]string  "Invalid request or error message"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /users [get]
func GetUsers(c *gin.Context) {

	var users []models.User

	users, err := repository.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
