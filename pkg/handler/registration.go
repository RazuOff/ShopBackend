package handler

import (
	"net/http"

	"ginexample.com/pkg/auth"
	"ginexample.com/pkg/repository"
	"github.com/gin-gonic/gin"
)

// Registrate godoc
// @Summary      Register a new user
// @Description  Creates a new user account if the username is available
// @Tags         authentication
// @Accept       json
// @Produce      json
// @Param        credentials  body      auth.Credentials  true  "User credentials"
// @Success      201  {object}  map[string]string  "Account created"
// @Failure      400  {object}  map[string]string  "User already exists or input error"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /registrate [post]
func Registrate(c *gin.Context) {
	var inputForm auth.Credentials

	if err := c.BindJSON(&inputForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "input error"})
		return
	}

	if _, err := repository.GetUserByLogin(inputForm.Username); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"massage": "user already exists"})
		return
	}

	if err := repository.AddUser(inputForm.Username, inputForm.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"massage": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "account created"})
}
