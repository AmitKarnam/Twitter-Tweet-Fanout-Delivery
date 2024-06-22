package controllers

import (
	rabbitmq "TweetDelivery/messagequeue/RabbitMQ"
	"TweetDelivery/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u *UserController) Get(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users.Users)
}

func (u *UserController) Post(c *gin.Context) {
	var newUser users.User

	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Check if the user already exists
	for _, user := range users.Users {
		if user == newUser {
			c.IndentedJSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		}
	}

	// Add the new user to the slice
	users.Users = append(users.Users, newUser)

	// Creates a queue for each user
	rabbitmq.CreateQueue(string(newUser))

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User added successfully"})

}
