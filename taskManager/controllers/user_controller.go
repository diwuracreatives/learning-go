package controllers

import (
	"net/http"
	"strings"
	"taskManager/data"
	"taskManager/models"
	"taskManager/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	userService *data.UserService
}

func NewUserController(userService *data.UserService) *UserController {
	return &UserController{userService: userService}
}

func (userController *UserController) SignUp(c *gin.Context) {
	var newUser models.UserInput

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

	user, err := userController.userService.CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func (userController *UserController) Login(c *gin.Context) {
	var userInput models.UserInput

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

	user, err := userController.userService.LoginUser(userInput)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	roleString := string(user.Role)

	token, err := utils.GenerateToken(user.ID.Hex(), user.Username, roleString)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"token": token, "user": user}})
}

func (userController *UserController) PromoteUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = userController.userService.PromoteUser(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated to admin successfully"})

}
