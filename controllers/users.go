package controllers

import (
	"net/http"

	"github.com/dennisferdian9/golang-sqlite/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := models.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
func GetOneUsers(c *gin.Context) {
	username := c.Param("username")
	users, err := models.GetOneUsers(username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
func PostUser(c *gin.Context) {
	var newUser models.Users
	if err := c.Bind(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "errorMessage": err})
		return
	}
	if _, err := models.PostUser(newUser.Username, newUser.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user", "errorLog": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Post Success",
		"data":    newUser,
	})

	// message, err := models.PostUser()

}
