package controllers

import (
	rabbitmq "TweetDelivery/messagequeue/RabbitMQ"
	"TweetDelivery/users"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FollowController struct{}

func (f *FollowController) Post(c *gin.Context) {
	currentUser := c.Param("user")

	var followUser users.User

	if err := c.BindJSON(&followUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	//Code to check is the user sent in the body is present or not
	userMap := make(map[users.User]bool)
	for _, user := range users.Users {
		userMap[user] = true
	}

	// Check if the followUser exists in the userMap
	if !userMap[followUser] {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "User to follow does not exist"})
		return
	}
	err := rabbitmq.CreateBinding(currentUser, string(followUser))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%s followed user %s successfully", currentUser, followUser)})
}
