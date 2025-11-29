package controllers

import (
	"net/http"
	"strings"
	"taskManager/domain"
	"taskManager/usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	userUseCases usecases.UserUseCases
}

func NewUserController(userUseCases usecases.UserUseCases) *UserController {
	return &UserController{userUseCases: userUseCases}
}

func (userController *UserController) SignUp(c *gin.Context) {
	var newUser domain.UserInput

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "request payload invalid",
			"message": err.Error(),
		})
		return
	}

	newUser.Username = strings.TrimSpace(newUser.Username)
	newUser.Password = strings.TrimSpace(newUser.Password)

	if newUser.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
		return
	} else if newUser.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is required"})
		return
	}

	user, err := userController.userUseCases.Register(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func (userController *UserController) Login(c *gin.Context) {
	var userInput domain.UserInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "request payload invalid",
			"message": err.Error(),
		})
		return
	}

	userInput.Username = strings.TrimSpace(userInput.Username)
	userInput.Password = strings.TrimSpace(userInput.Password)

	if userInput.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
		return
	} else if userInput.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is required"})
		return
	}

	token, err := userController.userUseCases.LoginUser(userInput)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"token": token}})
}

func (userController *UserController) PromoteUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = userController.userUseCases.PromoteUser(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated to admin successfully"})

}
