// File: controllers/user_controller.go
package controller

import (
	"apis/model"
	"apis/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	Service service.Service
}

func NewController(service service.Service) Controller {
	return Controller{
		Service: service,
	}
}

func (m Controller) GetUsers(c *gin.Context) {
	users, err := m.Service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
	}
	c.JSON(http.StatusOK, users)
}

func (m Controller) GetUserDetails(c *gin.Context) {
	name := c.Param("name")
	user, err := m.Service.GetUserByname(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (m Controller) CreateUser(c *gin.Context) {
	var user model.MainDB
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	m.Service.CreateUser(user)
	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func (m Controller) DeleteUser(c *gin.Context) {
	name := c.Param("name")
	if err := m.Service.DeleteUser(name); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func (m Controller) DecHell(c *gin.Context) {
	name := c.Param("name")
	err := m.Service.DecHell(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Hell incremented"})
}

func (m Controller) DecHeaven(c *gin.Context) {
	name := c.Param("name")
	err := m.Service.DecHeaven(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Heaven incremented"})
}

func (m Controller) IncHell(c *gin.Context) {
	name := c.Param("name")
	err := m.Service.IncHell(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Hell incremented"})
}

func (m Controller) IncHeaven(c *gin.Context) {
	name := c.Param("name")
	err := m.Service.IncHeaven(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Heaven incremented"})
}
